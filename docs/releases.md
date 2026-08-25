---
layout: default
title: Releases — Asciinema Editor
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
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a class="active" href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Releases

> *Fresh supplies from the vault — every tag is a drop.*

## Which download do I want?

| File | Platform | How to run |
|---|---|---|
| `ascinemaeditor-desktop-macos-universal.zip` | macOS 11+ — Intel + Apple Silicon (one bundle) | Unzip → **Asciinema Editor.app** |
| `ascinemaeditor-desktop-windows-amd64.zip` | Windows 10/11 64-bit | Unzip → `ascinemaeditor.exe` |
| `ascinemaeditor-server-linux-amd64` / `-arm64` | Linux | `chmod +x … && ./… serve` → `http://localhost:8080` |
| `ascinemaeditor-server-darwin-amd64` / `-arm64` | macOS (server mode) | `chmod +x … && ./… serve` → `http://localhost:8080` |

- **Gatekeeper (macOS):** ad-hoc signed. If blocked: right-click → **Open**, or `xattr -cr /Applications/Asciinema\ Editor.app`.
- **Windows:** needs WebView2 (preinstalled on most Win 10/11).
- **Linux desktop:** not published — use a server binary + any browser.
- **Checksums:** every release ships `checksums-sha256.txt` — `shasum -a 256 -c checksums-sha256.txt`.
- Old releases (before v0.4.0) named artifacts without the `server-`/`desktop-` prefix.

[View all releases on GitHub →](https://github.com/cldmnky/ascinemaeditor/releases)

## Two kinds of build

### Desktop apps (Wails) — recommended

Native window, no browser needed. Open/Save use OS file dialogs, drag-and-drop is disk-backed.

- macOS universal (Intel + Apple Silicon) — built on `macos-latest` via CGO against Cocoa/WebKit.
- Windows 10/11 64-bit — built on `windows-latest` (WebView2, no CGO).

### Server binaries (browser mode)

Single static executable — `CGO_ENABLED=0`, embedded UI, zero deps. Serves only `/` on `http://localhost:8080` (`PORT`/`HOST` configurable).

```bash
chmod +x ascinemaeditor-server-linux-amd64
./ascinemaeditor-server-linux-amd64 serve   # then open http://localhost:8080
```

Both build types contain the full editor and built-in sample data — you can start editing immediately after download.

## Cutting a release

Tags drive releases. Push a `v*` tag and [`.github/workflows/release.yml`](https://github.com/cldmnky/ascinemaeditor/blob/main/.github/workflows/release.yml) builds all platforms and publishes the GitHub Release with generated notes and SHA256 checksums.

```bash
git tag vX.Y.Z && git push origin vX.Y.Z
# e.g. git tag v0.4.0 && git push origin v0.4.0
```

The workflow builds:

- 4× `Server {linux,darwin}×{amd64,arm64}` on `ubuntu-latest`
- `desktop-darwin` (universal) on `macos-latest` — `wails build -platform darwin/universal`
- `desktop-windows` (amd64) on `windows-latest` — `wails.exe build -platform windows/amd64 -H windowsgui`
- `release` — downloads all artifacts, writes `checksums-sha256.txt`, publishes with `softprops/action-gh-release` and a “Which download do I want?” body.

Releases are the source of truth for downloadable artifacts; the docs site does not host binaries.

## Docs site

This site is published via **[`.github/workflows/pages.yml`](https://github.com/cldmnky/ascinemaeditor/blob/main/.github/workflows/pages.yml)** on every `push` to `main` that touches `docs/**`, `index.html`, or the workflow itself. It builds `docs/` with `actions/jekyll-build-pages` (Cayman theme) and deploys to GitHub Pages.

- Source: `docs/` on `main`
- Theme: `pages-themes/cayman@v0.2.0` via `jekyll-remote-theme`
- URL: `https://cldmnky.github.io/ascinemaeditor/` (`url` + `baseurl` in `docs/_config.yml`)
- Logo: `docs/assets/img/logo.svg` — Fallout vault gear + CRT, used in the header.
