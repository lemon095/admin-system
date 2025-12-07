const BASE = import.meta.env.VITE_API_BASE || ''

async function request(path, opts={}) {
  const res = await fetch(BASE + path, {
    headers: {'Content-Type':'application/json'},
    ...opts
  })
  const text = await res.text()
  let data = null
  try { data = text ? JSON.parse(text) : null } catch(e) { data = text }
  return { ok: res.ok, status: res.status, data }
}

export default {
  get: (path) => request(path, { method: 'GET' }),
  post: (path, body) => request(path, { method: 'POST', body: JSON.stringify(body) })
}
