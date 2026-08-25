---
layout: default
title: Installation — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a class="active" href="{{ '/installation' | relative_url }}">Install</a>
  <a href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a href="{{ '/editing' | relative_url }}">Editing</a>
  <a href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a href="{{ '/playback' | relative_url }}">Playback</a>
  <a href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Installation

> *Vault-Tec step one: get your Pip-Boy — er, editor — powered up.*

Two ways to run the same editor. Pick the one that fits your wasteland.

## Option A — Prebuilt binaries (fastest)

Grab the latest release from **[Releases on GitHub](https://github.com/cldmnky/ascinemaeditor/releases)**. Every tag ships checksums in `checksums-sha256.txt`.

| File | Platform | How to run |
|---|---|---|
| `ascinemaeditor-desktop-macos-universal.zip` | macOS 11+ — Intel + Apple Silicon | Unzip → double-click **Asciinema Editor.app** |
| `ascinemaeditor-desktop-windows-amd64.zip` | Windows 10/11 64-bit | Unzip → run `ascinemaeditor.exe` |
| `ascinemaeditor-server-linux-amd64` / `-arm64` | Linux | `chmod +x … && ./… serve` → open `http://localhost:8080` |
| `ascinemaeditor-server-darwin-amd64` / `-arm64` | macOS (server mode) | `chmod +x … && ./… serve` → open `http://localhost:8080` |

- **macOS Gatekeeper:** the `.app` is ad-hoc signed. If macOS blocks it, right-click → **Open**, or run `xattr -cr /Applications/Asciinema\ Editor.app`.
- **Windows:** requires the WebView2 runtime (preinstalled on most Windows 10/11). The window is WebView2 — no console.
- **Linux desktop:** not published yet — use a server binary + any browser.
- Old releases (before v0.4.0) named these without the `server-`/`desktop-` prefix.

Verify downloads:

```bash
shasum -a 256 -c checksums-sha256.txt
```

## Option B — Build from source

### Requirements

- **Go 1.25+** — `go version` should say `go1.25` or newer. The module pins `go 1.25.0`; `actions/setup-go` uses `go-version-file: go.mod` in CI.
- **go-task** (aka `task`) — `go install github.com/go-task/task/v3/cmd/task@latest`
- **Wails CLI** — only for the desktop app. Auto-installed to `./bin/wails` by `task wails:*` / `task desktop`, or manually: `GOBIN=$(pwd)/bin go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- **Platform deps for the desktop build:**
  - macOS: Xcode CLT (for CGO against Cocoa/WebKit). Windows build needs a Windows runner (CI handles it); cross-compiling desktop to Windows from Linux is not supported by Wails.

### Clone and run

```bash
git clone https://github.com/cldmnky/ascinemaeditor.git
cd ascinemaeditor

# Desktop — native window, no browser needed
task desktop          # build + run the .app / .exe
task wails:dev        # live-reload dev (Wails)
task wails:build      # production bundle → build/bin/

# Browser — same UI, served on :8080
task serve            # build + serve
task dev              # with file watching (watches main.go, index.html, test-data/)
go run . serve        # plain Go, no Taskfile
```

`task` with no args runs the desktop app (`task desktop`).

### Environment

| Var | Default | Applies to |
|---|---|---|
| `PORT` | `8080` | `serve` / server binaries |
| `HOST` | `127.0.0.1` | `serve` / server binaries |
| `ASCINEMAEDITOR_MODE=serve` | — | Force server mode even with desktop tags |

The server is single-binary, zero deps, and serves only `/` (everything else 404s). The UI is embedded via `//go:embed index.html` so you can launch it from any directory. For Wails, `frontend/dist/index.html` is synced from `index.html` at build time.

## Verify the build

```bash
go build ./... && go vet ./...
# then open http://localhost:8080 and try a sample from test-data/
```

Sample casts in `test-data/` are dev fixtures — they are not embedded in the binary.

## Next

- **[Quickstart →]({{ '/quickstart' | relative_url }})** — your first edit in 60 seconds.
- **[Releases →]({{ '/releases' | relative_url }})** — artifact table and how to cut a release.
