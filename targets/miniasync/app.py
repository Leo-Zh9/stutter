"""miniasync: async FastAPI routes over PostgreSQL, the shape stutter's held
commit (concurrent(#n@commitK)) has to handle.

  POST /pay                  Idempotency-Key; {"account", "amount_cents"}
  POST /threads/{id}/resume  answers the step pending on the thread
  POST /notes                Idempotency-Key; {"text"}
  GET  /healthz

/pay makes one sync psycopg call before it awaits the async pool, as apps do
that call a sync pool from an async route: while that call waits, the event
loop is frozen. /notes serializes on an in-process lock and leaves one write
to a background task. BUG=dup-charge checks for the charge, then inserts it
with no unique key; with BUG unset the insert is ON CONFLICT DO NOTHING.

stutter.json names an empty model script only to carry interrupt_sql.
"""

import argparse
import asyncio
import os
from contextlib import asynccontextmanager

import psycopg
from fastapi import FastAPI, Header
from psycopg_pool import AsyncConnectionPool, ConnectionPool
from pydantic import BaseModel

BUG = os.environ.get("BUG", "")
if BUG not in ("", "dup-charge"):
    raise SystemExit(f"unknown BUG={BUG!r}")
DB = os.environ["DATABASE_URL"]
SCHEMA = os.path.join(os.path.dirname(os.path.abspath(__file__)), "schema.sql")

# Two pools, not shared: sync calls block the loop, async ones yield it.
pool = ConnectionPool(DB, kwargs={"autocommit": True}, min_size=1, max_size=4, open=False)
apool = AsyncConnectionPool(DB, kwargs={"autocommit": True}, min_size=2, max_size=4, open=False)
lock = asyncio.Lock()
background = set()


@asynccontextmanager
async def lifespan(app):
    pool.open()
    await apool.open()
    yield
    await apool.close()
    pool.close()


app = FastAPI(lifespan=lifespan)


@app.get("/healthz")
async def healthz():
    return {"ok": True}


class Pay(BaseModel):
    account: str
    amount_cents: int


@app.post("/pay")
async def pay(req: Pay, idempotency_key: str = Header()):
    key = idempotency_key
    with pool.connection() as c:  # commit 1, sync: the loop waits here
        c.execute("UPDATE accounts SET last_key = %s WHERE id = %s", (key, req.account))
    async with apool.connection() as c:
        cur = await c.execute("SELECT 1 FROM charges WHERE key = %s", (key,))
        seen = await cur.fetchone()
        await c.execute("UPDATE accounts SET checked_key = %s WHERE id = %s", (key, req.account))  # commit 2
        if not seen:  # commit 3
            sql = "INSERT INTO charges (key, account, amount_cents) VALUES (%s, %s, %s)"
            if BUG != "dup-charge":
                sql += " ON CONFLICT (key) DO NOTHING"
            await c.execute(sql, (key, req.account, req.amount_cents))
    return {"key": key, "account": req.account, "amount_cents": req.amount_cents}


@app.post("/threads/{thread}/resume")
async def resume(thread: str):
    async with apool.connection() as c:
        # commit 1 answers the pending step once; stutter reads it as the interrupt
        await c.execute("UPDATE steps SET n = n + 1 WHERE thread = %s AND n = 0", (thread,))
        await c.execute("UPDATE steps SET resumed = true WHERE thread = %s", (thread,))  # commit 2
    return {"thread": thread, "status": "done"}


async def audit(key):
    async with apool.connection() as c:
        await c.execute("INSERT INTO audit (key) VALUES (%s) ON CONFLICT DO NOTHING", (key,))


class Note(BaseModel):
    text: str


@app.post("/notes")
async def notes(req: Note, idempotency_key: str = Header()):
    async with lock:  # a duplicate waits here until the first attempt answers
        async with apool.connection() as c:  # commit 1
            await c.execute("INSERT INTO notes (key, text) VALUES (%s, %s) ON CONFLICT DO NOTHING", (idempotency_key, req.text))
        t = asyncio.create_task(audit(idempotency_key))  # commit 2, off the response path
        background.add(t)
        t.add_done_callback(background.discard)
        await asyncio.sleep(0.1)  # the rest of the handler's work
    return {"key": idempotency_key}


def migrate():
    with psycopg.connect(DB, autocommit=True) as c:
        with open(SCHEMA) as f:
            c.execute(f.read())
        if BUG != "dup-charge":
            c.execute("CREATE UNIQUE INDEX charges_key ON charges (key)")


if __name__ == "__main__":
    ap = argparse.ArgumentParser()
    ap.add_argument("--migrate", action="store_true")
    ap.add_argument("--port", type=int, default=8000)
    a = ap.parse_args()
    if a.migrate:
        migrate()
    else:
        import uvicorn

        uvicorn.run(app, host="127.0.0.1", port=a.port, log_level="warning")
