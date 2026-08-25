# AGENTS.md

## Architecture

- Two-file app: `main.go` is only a static file server (stdlib, zero deps); the entire web editor lives in `index.html` (~1600 lines of vanilla JS + xterm.js loaded from jsDelivr CDN). App-logic changes go in `index.html`, not Go.
- xterm.js/fonts come from CDNs — terminal preview won't work offline.

## Running

- `task serve` (build + run) or `task dev` (live reload; watches `main.go`, `index.html`, `test-data/`). Install go-task: `go install github.com/go-task/task/v3/cmd/task@latest`.
- Plain alternative: `go run main.go`. Port via `PORT` env (default 8080).
- **The server serves files from its current working directory** (`http.Dir(".")`) — always launch from the repo root or `index.html` will 404. It binds to `127.0.0.1` by default (`HOST` env overrides), blocks dotfiles/dot-directories, and disables directory listings.

## Verification

No tests, lint, or CI exist. Verify with:

```bash
go build ./... && go vet ./...
```

plus a manual browser check at `http://localhost:8080` using sample casts from `test-data/`.
