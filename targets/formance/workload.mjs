// Labelled synthetic workload for Formance Ledger v2: the order/fee/seller
// numscript from the ledger's own tools/generator/examples/example1.js, sent
// as single keyed POSTs (stutter does not fault _bulk: Formance ignores the
// Idempotency-Key header there). Then reverts and metadata writes by id.
// The template already holds one setup transaction (see stutter.json): a
// ledger's first write resets its id sequences, which would undo banding.
// usage: node workload.mjs <base url>
const base = process.argv[2];
const plain = `vars {
    account $order
    account $seller
}
send [USD/2 100] (
    source = @world
    destination = $order
)
send [USD/2 1] (
    source = $order
    destination = @fees
)
send [USD/2 99] (
    source = $order
    destination = $seller
)`;
async function call(method, path, key, body) {
  const res = await fetch(base + path, {
    method, body: body === undefined ? undefined : JSON.stringify(body),
    headers: { 'Content-Type': 'application/json', 'Idempotency-Key': key },
  });
  const text = await res.text();
  if (res.status >= 300) {
    console.error(`${method} ${path}: HTTP ${res.status} ${text}`);
    process.exit(1);
  }
  return text ? JSON.parse(text) : null;
}
const ids = [];
for (let i = 0; i < 30; i++) {
  const r = await call('POST', '/v2/default/transactions', `tx-${i}`, {
    script: { plain, vars: { order: `orders:${String(i).padStart(4, '0')}`, seller: `sellers:${i % 5}` } },
  });
  ids.push(r.data.id);
}
for (let i = 0; i < 5; i++) {
  await call('POST', `/v2/default/transactions/${ids[i * 6]}/metadata`, `md-${i}`, { note: `checked-${i}` });
}
for (let i = 0; i < 5; i++) {
  await call('POST', `/v2/default/transactions/${ids[i * 6 + 1]}/revert`, `rv-${i}`);
}
for (let i = 0; i < 3; i++) {
  await call('POST', `/v2/default/accounts/sellers:${i}/metadata`, `am-${i}`, { tier: `gold-${i}` });
}
console.log(`sent ${30 + 5 + 5 + 3} requests`);
