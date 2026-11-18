"""Forge workload: two conversations through the project's run API, each paused by the
workflow's approval gate; one approved (the agent saves a memory), one rejected.
Every request carries its own Idempotency-Key, as the run API documents.
usage: python3 workload.py BASE_URL"""
import json
import sys
import urllib.request

base = sys.argv[1].rstrip("/")
PROJECT = "74ed7105-e6a4-47b7-a571-172b4eb7af47"


def run(key, body):
    req = urllib.request.Request(f"{base}/v1/projects/{PROJECT}/run", data=json.dumps(body).encode(),
                                 headers={"Content-Type": "application/json", "Idempotency-Key": key,
                                          "Authorization": "Bearer stutter-service-token"})
    with urllib.request.urlopen(req) as r:
        out = json.loads(r.read())
    print(key, json.dumps(out)[:300])
    return out


def say(key, text):
    r = run(key, {"input": {"messages": [{"role": "user", "content": text}]}, "stream": False})
    assert r["interrupted"], r
    return r["thread_id"]


t = say("acme-1", "Remember that customer ACME prefers email.")
r = run("acme-2", {"thread_id": t, "resume": {"value": "approve"}, "stream": False})
assert r["status"] == "done", r
t = say("globex-1", "Remember that customer Globex wants invoices in EUR.")
r = run("globex-2", {"thread_id": t, "resume": {"value": "reject"}, "stream": False})
assert r["status"] == "done", r
