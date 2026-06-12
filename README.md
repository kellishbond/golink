# GoLink 🔗

A fast, lightweight URL shortener API built from scratch in Go.

No frameworks. Pure Go standard library.

**Live API:** https://golink-production-8f50.up.railway.app/health

---

## What it does

Send GoLink a long URL → get back a short code → share it → track every click.

---

## Endpoints

### Health Check
```
GET /health
```
Response:
```json
{"status": "GoLink is running"}
```

---

### Shorten a URL
```
POST /shorten
```
Body:
```json
{"url": "https://www.youtube.com"}
```
Response:
```json
{
  "code": "Te6Db1",
  "short_url": "https://golink-production-8f50.up.railway.app/Te6Db1"
}
```

---

### Redirect
```
GET /:code
```
Redirects the user to the original URL.

Example:
```
https://golink-production-8f50.up.railway.app/Te6Db1
→ https://www.youtube.com
```

---

### View Stats
```
GET /stats/:code
```
Response:
```json
{
  "code": "Te6Db1",
  "url": "https://www.youtube.com",
  "clicks": 3
}
```

---

## How it works

1. You send a long URL to `POST /shorten`
2. GoLink generates a random 6-character code and stores the mapping in memory
3. Anyone who visits `/:code` gets redirected to the original URL
4. Every redirect increments a click counter
5. You can check click stats anytime via `GET /stats/:code`

---

## Tech Stack

- **Language:** Go
- **Router:** net/http (standard library only)
- **Storage:** In-memory map (no database)
- **Deploy:** Docker + Railway

---

## Run Locally

**Requirements:** Go 1.23+

```bash
git clone https://github.com/kellishbond/golink.git
cd golink
go run .
```

Server starts on `http://localhost:8080`

---

## Known Limitations

- Storage is in-memory — data resets on every redeploy
- No authentication — anyone can shorten URLs
- No custom short codes yet

**v2 will add:** Redis for persistent storage, custom codes, and an API key system.

---

## Author

Built by **Ishola Kolawole** (Koldev)

- GitHub: [@kellishbond](https://github.com/kellishbond)
- Twitter/X: [@kellishbond](https://x.com/kellishbond)
