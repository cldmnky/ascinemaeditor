---
layout: default
title: Playback — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
  <a href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a href="{{ '/editing' | relative_url }}">Editing</a>
  <a href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a class="active" href="{{ '/playback' | relative_url }}">Playback</a>
  <a href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Playback

> *Hit Play. Watch the wasteland replay exactly as you edited it.*

## Preview tab

The **Preview** tab is a real terminal powered by **ghostty-web** (WASM VT100, xterm.js-compatible API). It is 80×24 by default (from the cast header) and resizes via `r` events and on `seekTo`.

### Controls

| Control | What it does |
|---|---|
| **▶ Play / ⏸ Pause** | Start or pause playback. While playing the button is amber (`Pause`); while stopped it is sky (`Play`). |
| **↺ Loop** | When on, reaching the end clears the terminal and restarts from event 0. Toggle with the Loop button (`aria-pressed` reflects state). |
| **⟲ Reset** | Pauses, resets `currentIndex` to 0, clears the terminal, and re-renders Raw view. |
| **Preview / Raw .cast** tabs | Switch between the terminal and the serialized view. Re-entering Preview calls `fitAddon.fit()` (twice, with a delay) to reflow the terminal. |
| **Terminal chrome** | `bash — ghostty-web` label and `80×24` size badge. Traffic-light dots are decorative. |

![Terminal preview during playback]({{ '/screenshots/playback.png' | relative_url }})
*Playing at 6.069s — the progress bar and filmstrip highlight show position; the terminal replays the edited delays.*

### How playback works

- `play()` snapshots `getFlattenedEvents()` and `recomputeTiming()`, then schedules `setTimeout(delay * 1000)` per event.
- Only `o` (output) and `r` (resize) events affect the terminal; `i`/`m`/`x` are logged to the console.
- After each event fires, `currentIndex` advances and `updateActiveEventMarker()` highlights the active row and Raw line and updates the scrubber.
- If `loopEnabled` and the end is reached, playback clears the terminal and continues from 0.
- Any seek or edit pauses before re-rendering.

### Seeking updates the terminal

`seekTo(index)` pauses, clamps the index, clears the terminal, replays events `0..index` via `applyTerminalEvent`, and fits the terminal. Both the filmstrip and the Raw view scroll the active line into view.

## Raw .cast tab

The **Raw .cast** tab is the ground truth — what **Download .cast** / **Save** will write.

- Each line is `JSON.stringify([delay, type, data])`.
- `# …` comment lines appear exactly where they will be in the file, interleaved by `recordingComments` (see [Metadata & Comments]({{ '/metadata' | relative_url }})).
- The active event's line is highlighted (`bg-sky-900/40`) and scrolled into view during playback and after seeking.
- `renderRawView()` rebuilds the tab after every edit, seek, or rechunk.

![Raw .cast view with serialized events and comments]({{ '/screenshots/raw-view.png' | relative_url }})
*Raw view showing `# Chunk 1: Command — pwd` and serialized events like `[0.000296,"o","/tmp/asciinema-demo\\r\\n"]`.*

## What the player does with each type

| Type | Terminal | Raw view |
|---|---|---|
| `o` | `term.write(data)` (with resize handling for `r` payloads that slip through) | One JSON line |
| `r` | `term.resize(cols, rows)` parsed from `data` (`"80x24"`) | One JSON line |
| `i` | — (console log) | One JSON line |
| `m` | — (console log `Marker:`) | One JSON line |
| `x` | — (console log `Process exit:`) | One JSON line |

## Tips

- Use **Raw .cast** to verify comments and exact delays before saving — if it looks wrong there, the file will be wrong.
- **Loop** is useful when trimming pauses: play, listen, pause, tweak, loop again.
- If the preview is blank, check [Troubleshooting]({{ '/troubleshooting' | relative_url }}) — the ghostty-web/Tailwind/font CDNs may be blocked. Editing still works.
