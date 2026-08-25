---
layout: default
title: Asciinema Editor — Vault-Tec Approved
description: "Edit asciinema .cast files in a Vault-Tec approved terminal. Desktop + browser, ghostty-web + Tailwind."
---

<div class="doc-nav" markdown="0">
  <a class="active" href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
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

# Welcome to Vault 3, Editor

**Asciinema Editor** is a desktop (Wails) + browser editor for [asciinema](https://asciinema.org) `.cast` files. Scrub a recording on a horizontal timeline, edit delays and output per-event, preview it in a real terminal ([ghostty-web](https://github.com/mitchellh/ghostty-web)), then save a clean `.cast` you can publish anywhere.

> **Vault-Tec motto:** *Edit. Preview. Export. Survive the wasteland with perfect demos.*

<div class="vault-card" markdown="1">
**New here?** Start with **[Quickstart →]({{ '/quickstart' | relative_url }})** (60 seconds to your first edit), then read **[Timeline]({{ '/timeline' | relative_url }})** and **[Editing]({{ '/editing' | relative_url }})**.
</div>

![Timeline editor beside the terminal preview]({{ '/screenshots/editor-overview.png' | relative_url }})
*The editor on launch — built-in sample recording, horizontal timeline on top, timeline editor on the left, terminal preview + Raw view on the right.*

## What you get

- **Asciinema v3** with full header + comment preservation; **v2 import** (absolute timestamps → interval delays)
- **Chunk-based editing** — prompt detection groups events into command-sized chunks with collapsible sections
- **Horizontal timeline** — slider + colored ticks + chunk segments + delay-scaled filmstrip, all seekable
- **Terminal preview** — ghostty-web (WASM VT100, xterm.js API), 80×24 by default, resizes with the header
- **Tailwind CSS UI** — dark vault UI, responsive, works offline except for the preview CDN
- **Desktop app** — Wails window with native Open/Save dialogs and drag-and-drop; no browser needed
- **Raw `.cast` view** — see the exact serialized stream you will download

### Timeline at a glance

![Horizontal timeline with chunk segments, event ticks, and filmstrip]({{ '/screenshots/horizontal-timeline.png' | relative_url }})
*Top bar: event counter, chunk count, current/total time, slider, track, chunk segments, colored ticks, Prev/Next chunk, and the filmstrip. Every tick and block seeks.*

## Two ways to run

| Mode | Command | What it does |
|---|---|---|
| **Desktop** (recommended) | `task desktop` | Native window via Wails. Native dialogs, drag-and-drop, embedded UI. |
| **Browser** | `task serve` then open `http://localhost:8080` | Static Go server (`go run . serve`). Same UI, file picker + blob download. |

Both embed the full UI — you can start editing immediately with the built-in sample recording, no file needed.

```bash
# Desktop (Go 1.25+ — Wails CLI auto-installs to ./bin/wails)
task desktop          # build + run native app
task wails:dev        # live-reload dev
task wails:build      # production bundle → build/bin/

# Browser — quick iteration without Wails
task serve            # build + serve on :8080
task dev              # with file watching
go run . serve        # plain Go alternative
```

See **[Installation →]({{ '/installation' | relative_url }})** for requirements, prebuilt binaries, and `PORT`/`HOST` options.

## How a session goes

1. **Open** — `task desktop` or `task serve` → `http://localhost:8080`.
2. **Load** — **Upload .cast** or drop a `.cast` onto the window. Or just use the sample.
3. **Scrub** — drag the slider, click the track/ticks/filmstrip, or use the vertical list.
4. **Edit** — tweak delays and output per event, use **Speed 1x** and markers, remove/add events.
5. **Preview** — **Play** in the terminal; **Loop** / **Reset** as needed. Check **Raw .cast** for the exact output.
6. **Save** — **Download .cast** (browser → `edited-recording.cast`) or native Save dialog (desktop).

## The vault map

| Guide | What it covers |
|---|---|
| [Installation]({{ '/installation' | relative_url }}) | Requirements, Taskfile, prebuilt releases, building from source |
| [Quickstart]({{ '/quickstart' | relative_url }}) | Your first edit in 60 seconds |
| [Editing]({{ '/editing' | relative_url }}) | Delays, content, speed, markers, add/remove, event types |
| [Timeline]({{ '/timeline' | relative_url }}) | Slider, ticks, chunk segments, filmstrip, seeking |
| [Playback]({{ '/playback' | relative_url }}) | Preview, Raw view, Loop/Reset, terminal behavior |
| [Prompt Detection]({{ '/prompt-detection' | relative_url }}) | How chunks are formed, custom prompts, `bash-5.3$` example |
| [Metadata & Comments]({{ '/metadata' | relative_url }}) | Title/command/tags, chunk comments, `#` comment preservation |
| [File Format]({{ '/file-format' | relative_url }}) | v2 vs v3, event tuple, header fields, limits |
| [Troubleshooting]({{ '/troubleshooting' | relative_url }}) | CDN, validation, 404s, desktop build notes |
| [Releases]({{ '/releases' | relative_url }}) | Artifact table, checksums, how to cut a release |

## Screenshots

### Prompt detection

![Prompt detection configured for a bash-5.3$ prompt]({{ '/screenshots/prompt-detection.png' | relative_url }})
*Set **Prompt pattern** to `bash-5.3$` in **Cast Metadata (v3)** — the timeline regroups instantly into 4 chunks, anchored to the original comments.*

### Playback

![Terminal preview during playback]({{ '/screenshots/playback.png' | relative_url }})
*Mid-playback at 6.069s. Progress bar and scrubber show position; the ghostty-web terminal replays the edited delays.*

### Raw view

![Raw .cast view with serialized events and comments]({{ '/screenshots/raw-view.png' | relative_url }})
*Raw .cast tab — the exact JSONL you will save, with `# Chunk …` comments in place and the active line highlighted.*

---

<div class="vault-card" markdown="1">
**Vault-Tec reminder:** This editor keeps your edits in memory until you save. Download or Save before closing the window, wanderer.
</div>

*Built with [Wails](https://wails.io), [ghostty-web](https://github.com/mitchellh/ghostty-web), and Tailwind CSS. Logo: Fallout-inspired vault gear + Pip-Boy CRT — Vault-Tec approved, not affiliated with Bethesda.*
