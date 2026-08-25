---
layout: default
title: Quickstart — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
  <a class="active" href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a href="{{ '/editing' | relative_url }}">Editing</a>
  <a href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a href="{{ '/playback' | relative_url }}">Playback</a>
  <a href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Quickstart

> *Vault orientation, 60 seconds. No RadAway required.*

## 1. Launch

**Desktop:**

```bash
task desktop
# a native window opens — no browser needed
```

**Browser:**

```bash
task serve
# open http://localhost:8080
```

You will see the built-in sample recording right away — no file needed.

## 2. Load a recording (optional)

- Click **Upload .cast** (native dialog on desktop, file picker in browser) or **drag a `.cast` file** onto the window.
- Accepts asciinema **v3** and **v2** (v2 timestamps are converted to v3 interval delays on load).
- Limit: **50 MB**, **100,000 events**. Bigger files are rejected with a toast.

Try `test-data/` if you built from source.

## 3. Scrub

Pick any of these — they all move the same playhead:

- **Drag the horizontal slider** or **click the track**.
- **Click a colored tick** (event), a **filmstrip block**, an **event row**, or a **chunk segment**.
- **Prev chunk** / **Next chunk** to jump between prompt-detected chunks.
- <kbd>←</kbd> / <kbd>→</kbd> to step one event (when not typing in an input).
- <kbd>Space</kbd> to toggle Play/Pause (when not typing in an input).

![Horizontal timeline with chunk segments, event ticks, and filmstrip]({{ '/screenshots/horizontal-timeline.png' | relative_url }})

## 4. Edit a delay

In the **Timeline Editor** on the left, change any delay field — it is the interval *before* that event, in seconds. Must be a non-negative number. Invalid values are ignored.

Click **Speed 1x** on a chunk header to cycle selected events through `1x` → `2x` → `4x` → `8x` (their delays are scaled; checkboxes choose which events are affected).

## 5. Edit output

Change the content field for output/input/marker events. The editor preserves the event type:

| Type | Meaning |
|---|---|
| `o` | Terminal output |
| `i` | Input / keystrokes |
| `r` | Resize |
| `m` | Marker |
| `x` | Exit |

Control-only outputs (bare newline, bell) show as disabled special rows — only the delay is editable there.

## 6. Preview and save

- **Play** replays the edited delays in the terminal on the right. Use **Loop** to restart at the end, **Reset** to go back to the start.
- **Raw .cast** shows the exact serialized stream you will save, comments included.
- **Download .cast** — browser: downloads `edited-recording.cast`; desktop: native Save dialog.

![Terminal preview during playback]({{ '/screenshots/playback.png' | relative_url }})

<div class="vault-card" markdown="1">
**Remember:** edits live in memory until you save. Download/Save before closing the window.
</div>

## Next

- **[Editing →]({{ '/editing' | relative_url }})** — delays, content, speed, markers, add/remove in depth.
- **[Timeline →]({{ '/timeline' | relative_url }})** — what every part of the scrub bar means.
- **[Prompt Detection →]({{ '/prompt-detection' | relative_url }})** — chunk your timeline by shell prompt.
