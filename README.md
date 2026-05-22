# walkquest-service
Location-based exploration backend service for WalkQuest

## Run

```bash
go run ./cmd/server
```

The server listens on port `8080` by default.

```bash
curl http://localhost:8080/health
```

Expected response:

```json
{
  "status": "ok",
  "service": "walkquest-service"
}
```

To use a different port, set the `PORT` environment variable.
