"""miniagent: a small support agent built the way real LangGraph apps are.

FastAPI serves a LangGraph graph (an agent node calling an OpenAI-compatible
model, and the prebuilt ToolNode) that checkpoints to PostgreSQL with
PostgresSaver at LangGraph's default durability ("async"). One tool issues a
refund, but only after a human approves it through interrupt().

  POST /chat                     {"thread_id", "message_id", "text"}
  POST /threads/{id}/approve     {"decision": "approve" | "reject"}
  GET  /healthz

The BUG environment variable selects one seeded defect per variant, each a
mistake real agent code makes. With BUG unset the app is correct under
retries, double submits and crashes. See README.md.
"""

import argparse
import hashlib
import json
import os
from typing import Annotated

import psycopg
from fastapi import FastAPI, HTTPException
from fastapi.responses import JSONResponse
from langchain_core.messages import HumanMessage, ToolMessage
from langchain_core.runnables import RunnableConfig
from langchain_core.tools import InjectedToolCallId, tool
from langchain_openai import ChatOpenAI
from langgraph.checkpoint.postgres import PostgresSaver
from langgraph.graph import START, MessagesState, StateGraph
from langgraph.prebuilt import ToolNode, tools_condition
from langgraph.types import Command, interrupt
from psycopg.rows import dict_row
from psycopg_pool import ConnectionPool
from pydantic import BaseModel

BUG = os.environ.get("BUG", "")
VARIANTS = {
    "",
    "approval-no-claim",  # approval checked with SELECT, then acted on: double-click refunds twice
    "chat-no-dedupe",  # /chat ignores message_id: a retried message starts the run again
    "claim-before-refund",  # approval claimed in one commit, refund in the next: a crash loses it
    "note-keyed-on-tool-call",  # note deduped on tool_call_id, which a re-run model call changes
    "write-before-interrupt",  # control: a row written before interrupt() is written again on resume
}
if BUG not in VARIANTS:
    raise SystemExit(f"unknown BUG={BUG!r}")

DB = os.environ["DATABASE_URL"]
SCHEMA = os.path.join(os.path.dirname(os.path.abspath(__file__)), "schema.sql")

pool = ConnectionPool(DB, kwargs={"autocommit": True, "row_factory": dict_row}, min_size=1, max_size=20, open=False)


def one(conn, sql, *args):
    cur = conn.execute(sql, args)
    return cur.fetchone() if cur.description else None


# ---------------------------------------------------------------- tools


@tool
def lookup_order(order_id: int) -> str:
    """Look up an order by its number."""
    with pool.connection() as c:
        row = one(c, "SELECT number, customer, total_cents, status FROM orders WHERE number = %s", order_id)
    return json.dumps(row or {"error": "no such order"})


@tool
def add_order_note(order_id: int, note: str, config: RunnableConfig, tool_call_id: Annotated[str, InjectedToolCallId]) -> str:
    """Attach a note to an order."""
    thread = config["configurable"]["thread_id"]
    with pool.connection() as c:
        if BUG == "note-keyed-on-tool-call":
            key = tool_call_id  # changes whenever the model is called again
        else:
            key = hashlib.sha256(f"{thread}|{order_id}|{note}".encode()).hexdigest()
        c.execute(
            "INSERT INTO order_notes (order_number, note, dedupe_key) VALUES (%s, %s, %s) ON CONFLICT (dedupe_key) DO NOTHING",
            (order_id, note, key),
        )
    return json.dumps({"ok": True})


@tool
def issue_refund(order_id: int, amount_cents: int, config: RunnableConfig) -> str:
    """Refund part or all of an order. A human must approve it first."""
    thread = config["configurable"]["thread_id"]
    with pool.connection() as c:
        if BUG == "write-before-interrupt":
            # LangGraph re-runs a node from its start on resume, so this
            # insert runs once before the interrupt and again after it.
            c.execute(
                "INSERT INTO approvals (thread_id, order_number, amount_cents) VALUES (%s, %s, %s)",
                (thread, order_id, amount_cents),
            )
        else:
            # Keyed by the business action (this order in this thread), not
            # by the model's tool_call_id: if the agent node is re-run, a
            # real model mints a new id. On resume the row already exists,
            # so this is a read and the node writes nothing before the
            # approval decision.
            if not one(c, "SELECT 1 FROM approvals WHERE thread_id = %s AND order_number = %s", thread, order_id):
                c.execute(
                    "INSERT INTO approvals (thread_id, order_number, amount_cents) VALUES (%s, %s, %s)",
                    (thread, order_id, amount_cents),
                )
    decision = interrupt({"action": "issue_refund", "order_id": order_id, "amount_cents": amount_cents})
    with pool.connection() as c:
        if decision != "approve":
            c.execute(
                "UPDATE approvals SET status = 'rejected' WHERE thread_id = %s AND order_number = %s AND status = 'pending'",
                (thread, order_id),
            )
            return json.dumps({"refunded": False, "reason": "rejected"})
        if BUG == "approval-no-claim":
            # Check, then act: two approvals that read 'pending' at the same
            # time both refund.
            row = one(c, "SELECT id, status FROM approvals WHERE thread_id = %s AND order_number = %s", thread, order_id)
            if row["status"] != "pending":
                return json.dumps({"refunded": False, "reason": "already " + row["status"]})
            with c.transaction():
                c.execute("INSERT INTO refunds (approval_id, order_number, amount_cents) VALUES (%s, %s, %s)", (row["id"], order_id, amount_cents))
                c.execute("INSERT INTO outbox (kind, order_number) VALUES ('refund_email', %s)", (order_id,))
                c.execute("UPDATE approvals SET status = 'approved' WHERE id = %s", (row["id"],))
            return json.dumps({"refunded": True})
        if BUG == "claim-before-refund":
            # The claim commits on its own; a crash before the refund commits
            # leaves an approval that is 'approved' with nothing refunded,
            # and the retry finds nothing left to claim.
            claimed = one(
                c,
                "UPDATE approvals SET status = 'approved' WHERE thread_id = %s AND order_number = %s AND status = 'pending' RETURNING id",
                thread, order_id,
            )
            if not claimed:
                return json.dumps({"refunded": False, "reason": "already handled"})
            with c.transaction():
                c.execute("INSERT INTO refunds (approval_id, order_number, amount_cents) VALUES (%s, %s, %s)", (claimed["id"], order_id, amount_cents))
                c.execute("INSERT INTO outbox (kind, order_number) VALUES ('refund_email', %s)", (order_id,))
            return json.dumps({"refunded": True})
        # Correct: claim and act in one transaction. Only the attempt whose
        # UPDATE moves the row out of 'pending' refunds, and a crash either
        # loses both or keeps both.
        with c.transaction():
            claimed = one(
                c,
                "UPDATE approvals SET status = 'approved' WHERE thread_id = %s AND order_number = %s AND status = 'pending' RETURNING id",
                thread, order_id,
            )
            if not claimed:
                return json.dumps({"refunded": False, "reason": "already handled"})
            c.execute("INSERT INTO refunds (approval_id, order_number, amount_cents) VALUES (%s, %s, %s)", (claimed["id"], order_id, amount_cents))
            c.execute("INSERT INTO outbox (kind, order_number) VALUES ('refund_email', %s)", (order_id,))
        return json.dumps({"refunded": True})


TOOLS = [lookup_order, add_order_note, issue_refund]
model = ChatOpenAI(
    model="gpt-4o-mini",
    base_url=os.environ.get("OPENAI_BASE_URL"),
    api_key=os.environ.get("OPENAI_API_KEY", "sk-unused"),
    temperature=0,
    max_retries=0,
).bind_tools(TOOLS)


def agent(state: MessagesState):
    return {"messages": [model.invoke([("system", "You are a support agent for a small shop.")] + state["messages"])]}


builder = StateGraph(MessagesState)
builder.add_node("agent", agent)
builder.add_node("tools", ToolNode(TOOLS))
builder.add_edge(START, "agent")
builder.add_conditional_edges("agent", tools_condition)
builder.add_edge("tools", "agent")

saver_pool = ConnectionPool(DB, kwargs={"autocommit": True, "prepare_threshold": 0, "row_factory": dict_row}, min_size=1, max_size=20, open=False)
graph = None


# ---------------------------------------------------------------- HTTP

app = FastAPI()


@app.on_event("startup")
def startup():
    global graph
    pool.open()
    saver_pool.open()
    graph = builder.compile(checkpointer=PostgresSaver(saver_pool))


class Chat(BaseModel):
    thread_id: str
    message_id: str
    text: str


class Decision(BaseModel):
    decision: str


def view(thread_id):
    st = graph.get_state({"configurable": {"thread_id": thread_id}})
    msgs = st.values.get("messages", []) if st.values else []
    last = msgs[-1].content if msgs else ""
    out = {"thread_id": thread_id, "reply": last if isinstance(last, str) else json.dumps(last)}
    if st.interrupts:
        out["status"] = "awaiting_approval"
        out["approval"] = st.interrupts[0].value
    else:
        out["status"] = "done"
    return out


def lock_key(s):
    return int.from_bytes(hashlib.sha256(s.encode()).digest()[:8], "big", signed=True)


@app.get("/healthz")
def healthz():
    return {"ok": True}


@app.post("/chat")
def chat(req: Chat):
    config = {"configurable": {"thread_id": req.thread_id}}
    if BUG == "chat-no-dedupe":
        graph.invoke({"messages": [HumanMessage(req.text)]}, config)
        return view(req.thread_id)
    with pool.connection() as c:
        # Claim the message id first. A duplicate that arrives while the
        # first attempt runs is told to retry (409); one that arrives after
        # it finished gets the stored reply; one that arrives after the
        # first attempt's process died recovers the run.
        claimed = one(
            c,
            "INSERT INTO chat_requests (message_id, thread_id) VALUES (%s, %s) ON CONFLICT (message_id) DO NOTHING RETURNING message_id",
            req.message_id, req.thread_id,
        )
        if not one(c, "SELECT pg_try_advisory_lock(%s) AS ok", lock_key(req.message_id))["ok"]:
            return JSONResponse({"error": "in progress"}, status_code=409)
        try:
            if not claimed:
                row = one(c, "SELECT status, response FROM chat_requests WHERE message_id = %s", req.message_id)
                if row["status"] == "done":
                    return row["response"]
            st = graph.get_state(config)
            ids = [m.id for m in (st.values or {}).get("messages", [])]
            if req.message_id not in ids:
                graph.invoke({"messages": [HumanMessage(req.text, id=req.message_id)]}, config)
            elif st.tasks and not st.interrupts:
                # Finish a run a crash cut short. Check tasks, not next: a
                # checkpoint whose pending tasks all saved their writes
                # reports next == () although the run never moved past it.
                graph.invoke(None, config)
            out = view(req.thread_id)
            c.execute("UPDATE chat_requests SET status = 'done', response = %s WHERE message_id = %s", (json.dumps(out), req.message_id))
            return out
        finally:
            c.execute("SELECT pg_advisory_unlock(%s)", (lock_key(req.message_id),))


@app.post("/threads/{thread_id}/approve")
def approve(thread_id: str, req: Decision):
    config = {"configurable": {"thread_id": thread_id}}
    st = graph.get_state(config)
    if not st.interrupts:
        if st.tasks:
            graph.invoke(None, config)  # a crash cut the resumed run short (see /chat)
        return view(thread_id)
    graph.invoke(Command(resume=req.decision), config)
    return view(thread_id)


def migrate():
    with psycopg.connect(DB, autocommit=True) as c:
        with open(SCHEMA) as f:
            c.execute(f.read())
        PostgresSaver(c).setup()
        c.execute(
            "INSERT INTO orders (number, customer, total_cents, status) VALUES "
            "(1001, 'ada@example.test', 4200, 'delivered'), (1002, 'grace@example.test', 1999, 'shipped'), "
            "(1003, 'linus@example.test', 8800, 'delivered') ON CONFLICT DO NOTHING"
        )


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
