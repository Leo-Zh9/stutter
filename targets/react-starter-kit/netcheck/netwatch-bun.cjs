// Observation only: logs every outbound socket connect, DNS lookup, fetch,
// http(s) request and Bun.connect made by this process to $NETWATCH_LOG.
const fs = require('node:fs'), net = require('node:net'), dns = require('node:dns'), tls = require('node:tls'), http = require('node:http'), https = require('node:https');
const log = (m) => { try { fs.appendFileSync(process.env.NETWATCH_LOG, `${process.pid} ${process.argv.slice(1).join(' ').slice(0, 60)} ${m}\n`); } catch {} };
log('start');
const desc = (o, p) => { try { if (typeof o === 'object' && o) { if (o.path) return `path=${o.path}`; if (o.port !== undefined) return `${o.host}:${o.port}`; return 'obj ' + JSON.stringify(o, (k, v) => typeof v === 'function' ? '[fn]' : v).slice(0, 200); } return `${typeof o}:${o} ${typeof p}:${p}`; } catch (e) { return 'desc-err ' + e; } };
const oc = net.Socket.prototype.connect;
net.Socket.prototype.connect = function (...a) { let o = a[0]; if (Array.isArray(o)) o = o[0]; log(`socket.connect ${desc(o, a[1])}`); return oc.apply(this, a); };
for (const [mod, name] of [[net, 'net'], [tls, 'tls']]) {
  const o = mod.connect; mod.connect = function (...a) { log(`${name}.connect ${desc(a[0], a[1])}`); return o.apply(this, a); };
  if (mod.createConnection) { const c = mod.createConnection; mod.createConnection = function (...a) { log(`${name}.createConnection ${desc(a[0], a[1])}`); return c.apply(this, a); }; }
}
for (const [mod, name] of [[http, 'http'], [https, 'https']]) {
  for (const f of ['request', 'get']) { const o = mod[f]; mod[f] = function (...a) { const u = a[0]; log(`${name}.${f} ${typeof u === 'string' || u instanceof URL ? String(u) : (u && (u.hostname || u.host))}`); return o.apply(this, a); }; }
}
const ol = dns.lookup; dns.lookup = function (h, ...r) { log(`dns.lookup ${h}`); return ol.call(this, h, ...r); };
if (dns.promises) { const pl = dns.promises.lookup; dns.promises.lookup = function (h, ...r) { log(`dns.promises.lookup ${h}`); return pl.call(this, h, ...r); }; }
if (typeof globalThis.fetch === 'function') { const of = globalThis.fetch; globalThis.fetch = function (i, ...r) { log(`fetch ${typeof i === 'string' || i instanceof URL ? String(i) : i && i.url}`); return of.call(this, i, ...r); }; }
if (typeof Bun !== 'undefined') {
  for (const f of ['connect', 'listen']) { const o = Bun[f]; if (o) try { Bun[f] = function (opts, ...r) { log(`Bun.${f} ${opts && (opts.unix || `${opts.hostname}:${opts.port}`)}`); return o.call(this, opts, ...r); }; } catch (e) { log(`cannot wrap Bun.${f}: ${e}`); } }
  if (Bun.dns && Bun.dns.lookup) try { const o = Bun.dns.lookup; Bun.dns.lookup = function (h, ...r) { log(`Bun.dns.lookup ${h}`); return o.call(this, h, ...r); }; } catch {}
}
