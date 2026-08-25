---
layout: default
title: File Format — Asciinema Editor
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
  <a class="active" href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# File Format

> *Know your scrolls. Vault-Tec archives in v3.*

## Asciinema `.cast` in 30 seconds

A `.cast` file is **JSON Lines** — one JSON value per line:

```
{"version":3,"term":{"cols":80,"rows":24},"title":"demo","timestamp":1714000000}
[0.023447,"o","bash-5.3$ "]
[2.946602,"o","p"]
[0.083101,"o","w"]
# Chunk 1: Command - pwd        ← v3 comment
[0.000296,"o","/tmp/asciinema-demo\r\n"]
…
```

- **Line 1** — header object. Must have `version: 2` or `3`; anything else is rejected.
- **Lines 2+** — events as `[delay, type, data]` **or** `# …` comment lines (v3 only).
- Blank lines are ignored.

## v2 vs v3

|  | v3 (preferred) | v2 (imported) |
|---|---|---|
| **Time** | `delay` — interval *before* this event (seconds, `≥ 0`) | Absolute timestamp since start — converted to interval on load: `interval = ts - prevTs` |
| **Header** | `{ version: 3, term: { cols, rows }, title?, command?, tags?, timestamp? }` | `{ version: 2, width, height, timestamp?, title?, command? }` — normalized to v3 via `normalizeHeaderToV3()` |
| **Comments** | Yes — `# …` lines, preserved with position | Not part of spec — not preserved |

Uploading a v2 file shows the toast *Cast loaded (v2 converted to v3)*. The saved file is always v3.

## Event tuple

Each event is a JSON array of three elements:

```js
[delay, type, data]
```

| Field | Type | Constraints |
|---|---|---|
| `delay` | number | Finite, `≥ 0`. Seconds before this event fires. Rejected otherwise. |
| `type` | string | One of `o`, `i`, `r`, `m`, `x` (see below). Editor preserves the original type on edit. |
| `data` | string | Payload. For `r`, a size like `"80x24"`; for `m`, marker text; for `o`/`i`/`x`, terminal or process data. |

Validation on load (`parseCastText`) rejects missing fields, wrong types, negative delays, and too many events. The currently open recording is **not** replaced on error — a toast explains the first error.

| Type | Meaning | Example `data` |
|---|---|---|
| `o` | Output to terminal | `"hello\r\n"` |
| `i` | Input / keystrokes | `"ls\r"` |
| `r` | Resize | `"120x30"` |
| `m` | Marker (editor extension) | `"Marker for Chunk 2"` |
| `x` | Exit code / process exit | `"0"` |

The editor's sample recording uses only `o` events (see `initialRecordingData` in `index.html`).

## Header

The editor normalizes and preserves all header keys except the v2→v3 renames. Minimal v3 header:

```json
{"version":3,"term":{"cols":80,"rows":24}}
```

Optional keys the UI edits:

- `title` — string (see [Metadata]({{ '/metadata' | relative_url }})).
- `command` — string.
- `tags` — `string[]` (stored as comma-separated in the form).
- `timestamp` — number (seconds since epoch). Kept as-is; not edited in the UI.

`normalizeHeaderToV3` defaults `cols` to 80 and `rows` to 24 if missing or non-finite, and moves legacy `width`/`height`/`rows` into `term`.

## Comments

`# …` lines are v3-only (see [Metadata & Comments]({{ '/metadata' | relative_url }})). They round-trip losslessly and participate in Raw view and output. The editor tracks them by `eventIndex` and emits them interleaved with events on save.

## Limits

| Limit | Value | Enforced |
|---|---|---|
| File size | 50 MB | On Upload / drop / Wails `OpenCast` |
| Event count | 100,000 | On parse (`MAX_EVENTS`) |
| Delay | `≥ 0`, finite | On parse + on delay edits |

Exceeding file size shows *File too large (limit 50 MB)*; too many events shows *Too many events (limit …)*.

## What gets written

`buildCastContent()` is the serializer:

1. `normalizeHeaderToV3({ ...recordingHeader })` + `saveMetadataFromForm()` → `JSON.stringify(header) + "\n"`.
2. Interleave `recordingComments` and `getFlattenedEvents()` by `eventIndex` → `comment + "\n"` then `JSON.stringify(event) + "\n"`.

**Raw .cast** (`renderRawView`) renders this same interleaving as HTML so what you see there is what you will download — check it before saving.

## Compatibility

Files saved by the editor are standard asciinema v3 and play with `asciinema play`. Markers (`m`) are an editor convention — `asciinema` ignores unknown types, so they do not affect playback elsewhere, but they do appear in the file.
