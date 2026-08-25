---
layout: default
title: Editing — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
  <a href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a class="active" href="{{ '/editing' | relative_url }}">Editing</a>
  <a href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a href="{{ '/playback' | relative_url }}">Playback</a>
  <a href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Editing

> *Tinker with time and text. The wasteland rewards the patient.*

![Timeline editor beside the terminal preview]({{ '/screenshots/editor-overview.png' | relative_url }})
*Left: the timeline editor (collapsible chunks). Right: Preview / Raw .cast.*

## Events are rows

Each row is one asciinema event: `[delay, type, data]` (see [File Format]({{ '/file-format' | relative_url }})). The editor groups events into **chunks** by prompt detection (see [Prompt Detection]({{ '/prompt-detection' | relative_url }})).

| Control | What it does |
|---|---|
| **Delay** (number field) | Interval *before* this event, in seconds. Must be `≥ 0`. Invalid entries are ignored; the field reverts. |
| **Content** (text field) | Event payload — output, input, or marker text. Preserves the event's `type`; only `data` changes. |
| **Checkbox** | Whether **Speed 1x** affects this event. |
| **📍** (pin on a row) | Insert a marker (`[0.1, "m", "Marker after event N"]`) after that event. |
| **Trash** | Remove the event. Comments are shifted to stay anchored. |
| **Select** (chunk header) | Master checkbox — check/uncheck all events in the chunk. |
| **Speed 1x** (chunk header) | Cycle selected events through `1x` → `2x` → `4x` → `8x` by scaling delays. |
| **📍 Marker** (chunk header) | Add a marker at the start of the chunk. |
| **+ Add Event** (bottom) | Append a blank output event (`[0.1, "o", ""]`) to the last chunk. |
| **Collapsible header** (▼/▶) | Collapse/expand a chunk. Click the header (not the buttons/inputs). |

### Speed

`Speed` cycles `1x` → `2x` → `4x` → `8x` → `1x` per chunk. For each selected event it divides the delay by the speed factor (e.g. `2x` halves the delays). The button label and `data-speed` update, the delay inputs refresh, and checkboxes keep their state. If playback is running, it pauses and resumes at the same index.

### Markers

Markers are editor-only events: `[delay, "m", text]`. The player logs them to the console and the terminal ignores them. They are serialized in the output `.cast` and round-trip through comments/positions like any other event.

### Add / Remove

- **Remove** shifts `recordingComments` down for indices after the removed event and calls `syncChunkComments` so chunk headings stay consistent.
- **Add Event** appends to the last chunk (or creates chunk 1 if empty). Comments are shifted for insertion at the end, chunk indexes are refreshed, and the editor scrolls to the new event.

### Special rows

Control-only outputs (bare `\r`, `\n`, `\b`, bell) with no visible text render as **disabled** rows — you can edit the delay but not the content. Any `o` event that contains visible text is editable even if it ends with a newline.

## Event types

| Type | Meaning | Editable | Playback |
|---|---|---|---|
| `o` | Terminal output | Yes (unless control-only) | Written to the ghostty-web terminal |
| `i` | Input / keystrokes | Yes | Logged to console |
| `r` | Terminal resize (`"80x24"` etc.) | Yes | `term.resize(cols, rows)` + `fitAddon.fit()` |
| `m` | Editor marker | Yes | Logged to console |
| `x` | Process exit | Yes | Logged to console |

`type` itself is not editable — changing content keeps the original type.

## Rechunking

Changing **Prompt pattern** (in [Cast Metadata]({{ '/metadata' | relative_url }})) regroups events via `rechunkWithCurrentPrompt()` without losing comments. The active `currentIndex` is not reset — the timeline re-renders around the same position.

## Live previews

After every structural change (remove, add, marker, speed, delay/data edit), the editor calls:

- `renderTimeline()` — rebuild chunk/event DOM
- `renderRawView()` — rebuild the Raw tab
- `updateHorizontalTimeline()` — rebuild ticks/segments/filmstrip and scrubber state

Edits live in memory until you **Download .cast** / **Save**. Closing the window without saving loses them.

## Tips

- Use **Prompt Detection** first, then edit chunks — otherwise you will edit one giant chunk.
- Use **Speed** to tighten pauses without hand-editing every delay.
- Drop a marker before a tricky splice so you can find it again in Raw view.
- Check Raw view before saving — it is the ground truth.
