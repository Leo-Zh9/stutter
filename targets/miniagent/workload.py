"""miniagent workload: four conversations, two of them through an approval.

usage: python3 workload.py BASE_URL
"""
import json
import sys
import urllib.request

base = sys.argv[1].rstrip("/")


def post(path, body):
    req = urllib.request.Request(base + path, data=json.dumps(body).encode(), headers={"Content-Type": "application/json"})
    with urllib.request.urlopen(req) as r:
        out = json.loads(r.read())
    print(path, out.get("status"), "|", out.get("reply", "")[:60])
    return out


assert post("/chat", {"thread_id": "t-ada", "message_id": "m-1", "text": "Hi, please refund order 1001"})["status"] == "awaiting_approval"
assert post("/threads/t-ada/approve", {"decision": "approve"})["status"] == "done"
post("/chat", {"thread_id": "t-grace", "message_id": "m-2", "text": "Please note that order 1002 was delayed"})
post("/chat", {"thread_id": "t-linus", "message_id": "m-3", "text": "What is the status of order 1003?"})
assert post("/chat", {"thread_id": "t-linus", "message_id": "m-4", "text": "OK, refund order 1003 for the broken part"})["status"] == "awaiting_approval"
assert post("/threads/t-linus/approve", {"decision": "reject"})["status"] == "done"
