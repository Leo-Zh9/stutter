"""paid-media-agent workload: three change requests through the self-hosted API,
two approved and one rejected. usage: python3 workload.py BASE_URL"""
import json
import sys
import urllib.request

base = sys.argv[1].rstrip("/")
H = {"Content-Type": "application/json", "Authorization": "Bearer stutter-token"}


def post(path, body):
    req = urllib.request.Request(base + path, data=json.dumps(body).encode(), headers=H)
    with urllib.request.urlopen(req) as r:
        out = json.loads(r.read())
    print(path, json.dumps(out)[:300])
    return out


def ask(thread, text):
    r = post(f"/threads/{thread}/messages", {"text": text})
    assert r["interrupted"] and r["proposal"], r
    return r["proposal"]["proposal_id"]


p = ask("t-google", "Lower the daily budget of the Brand Search - US campaign in demo-google to 150.")
r = post(f"/proposals/{p}/approve", {})
assert r.get("receipt"), r
p = ask("t-meta", "Pause the Prospecting - Lookalike 1% campaign in demo-meta.")
post(f"/proposals/{p}/reject", {"message": "Not this week."})
p = ask("t-reddit", "Raise the daily budget of the Community Targeting - Dev Tools campaign in demo-reddit to 90.")
r = post(f"/proposals/{p}/approve", {})
assert r.get("receipt"), r
