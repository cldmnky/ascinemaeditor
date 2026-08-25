---
layout: default
title: Troubleshooting — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
  <a href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a href="{{ '/editing' | relative_url }}">Editing</a>
  <a href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a href="{{ '/playback' | relative_url }}">Playback</a>
  <a href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a class="active" href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Troubleshooting

> *The vault door stuck? We have a manual override.*

## Preview is blank but editing works

The terminal preview depends on CDNs:

- **ghostty-web** — `cdn.jsdelivr.net/npm/ghostty-web@0.4.0/+esm`
- **Tailwind** — `cdn.tailwindcss.com`
- **Fonts** — `fonts.googleapis.com` / `fonts.gstatic.com` (Fira Code + Inter)
- **es-module-shims** — `unpkg.com`

If any are blocked (offline, firewall, corporate proxy), editing still works but the ghostty canvas will not initialize. Check the browser console for `ghostty-web init failed` and network errors. The same applies to the desktop app — it loads the same CDNs via the WebView.

Fix: allow those domains, or serve the assets locally (not built-in — you would need to patch `index.html`).

## “Invalid header” / “Line N: …” toast on upload

The file failed `parseCastText` validation. Common causes:

- First line is not a JSON object with `version: 2` or `3`.
- An event line is not `[delay, type, data]` with `delay` a finite `≥ 0` number and `type`/`data` strings.
- Invalid JSON on a line.

The currently open recording is **not** replaced — fix the file and re-upload. Check `Line N: …` in the toast for the first error; `JSON.parse` errors point at the header, per-line errors point at an event.

## “File too large (limit 50 MB)” or “Too many events (limit 100000)”

Hard limits: 50 MB file size, 100,000 events. Enforced on Upload / drop / Wails `OpenCast`. Trim the recording before editing (e.g. `asciinema` rec with a smaller terminal or shorter session).

## Drag-and-drop says “Please drop a .cast file”

The drop handler checks `file.name.endsWith('.cast')`. Rename the file to end in `.cast`, or use **Upload .cast**.

## Nothing happens on `Upload .cast` in desktop

The Wails binding `window.go.main.App.OpenCast()` is only present in the desktop build (`-tags desktop` / `wails build`). In browser/`serve` mode the button falls back to the hidden `<input>` picker. If desktop shows no dialog, you launched the server binary, not the desktop `.app`/`.exe`.

## Keyboard shortcuts do nothing

<kbd>←</kbd>, <kbd>→</kbd>, and <kbd>Space</kbd> are ignored when focus is inside an `<input>` or `<textarea>` (including delay/content fields and the Prompt pattern). Click outside the inputs or press <kbd>Esc</kbd> to unfocus, then try again.

## `xattr -cr` / Gatekeeper on macOS

The macOS `.app` is ad-hoc signed. On first launch macOS may say it is damaged or from an unidentified developer:

- Right-click the app → **Open** → **Open** again, or
- `xattr -cr /Applications/Asciinema\ Editor.app` (adjust the path).

## `go build` or desktop build fails

- `go 1.25+` is required (`go.mod` says `go 1.25.0`). CI uses `go-version-file: go.mod`.
- Desktop on macOS needs Xcode CLT for CGO (Cocoa/WebKit). Desktop on Windows needs a Windows runner — cross-compiling desktop to Windows from macOS/Linux is not supported by Wails.
- If `frontend/dist/index.html` is missing, run `task build` or `task wails:build` — they sync `index.html` → `frontend/dist/index.html`.
- `task` itself: `go install github.com/go-task/task/v3/cmd/task@latest`.

## Every path 404s

The server intentionally serves only `/`. Everything else returns 404 — there is no router, no static dir, no API. Launch it from any directory; the UI is embedded via `//go:embed`.

## Still stuck?

- Run `go build ./... && go vet ./...` and try a sample from `test-data/`.
- Check the browser console and `gh` logs for the Pages/Release workflows.
- [Open an issue](https://github.com/cldmnky/ascinemaeditor/issues) with the file (or a redacted excerpt) and the exact toast text.
