# AGENTS.md

## Architecture

- Two-file app: `main.go` serves only the embedded UI (`//go:embed index.html`, stdlib, zero deps); the entire web editor lives in `index.html` (~1600 lines of vanilla JS + xterm.js loaded from jsDelivr CDN). App-logic changes go in `index.html`, not Go.
- The editor ships with built-in sample data (`initialRecordingData` in `index.html`). Do not embed `test-data/` casts into the binary; they are dev fixtures only, served by nothing.
- xterm.js/fonts come from CDNs — terminal preview won't work offline.

## Running

- `task serve` (build + run) or `task dev` (live reload; watches `main.go`, `index.html`, `test-data/`). Install go-task: `go install github.com/go-task/task/v3/cmd/task@latest`.
- Plain alternative: `go run main.go`. Port via `PORT` env (default 8080).
- The UI is embedded in the binary, so it can be launched from any directory. It binds to `127.0.0.1` by default (`HOST` env overrides) and serves only `/`; every other path 404s.

## Releases

- Tag a release with `v*` (e.g. `v0.1.0`) and push the tag; `.github/workflows/release.yml` builds linux/darwin amd64+arm64 binaries and publishes them to a GitHub release with SHA256 checksums.

## Verification

No tests or lint exist. Verify with:

```bash
go build ./... && go vet ./...
```

plus a manual browser check at `http://localhost:8080` using sample casts from `test-data/`.
