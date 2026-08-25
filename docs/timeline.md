---
layout: default
title: Timeline — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
  <a href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a href="{{ '/editing' | relative_url }}">Editing</a>
  <a class="active" href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a href="{{ '/playback' | relative_url }}">Playback</a>
  <a href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Timeline

> *The vault corridor — it all connects. Slider, ticks, chunks, filmstrip.*

![Horizontal timeline with chunk segments, event ticks, and filmstrip]({{ '/screenshots/horizontal-timeline.png' | relative_url }})
*The horizontal timeline sits above the editor and preview. Drag the slider or click any element to seek.*

## Anatomy

| Part | What it is | Interaction |
|---|---|---|
| **Event counter** (`17 / 33`) | Current event index / total events | Read-only — updates as you seek or play |
| **Chunk badge** (`4 chunks`) | Prompt-detected chunk count | Read-only |
| **Current / total time** (`6.069s / 9.489s`) | Playback time at the playhead / total duration | Read-only |
| **Slider** (hidden range input over the track) | Time-based scrubber (0–100% of duration) | Drag to seek; pauses playback |
| **Track** (rounded bar) | Seek surface | Click anywhere on the track to seek |
| **Progress fill** (blue) | Elapsed time | Visual only |
| **Chunk segments** (alternating faint bands in the track) | Width = chunk duration / total duration | Click a segment to seek to that chunk |
| **Ticks** (thin vertical lines) | One per event (capped at ~300 for performance), colored by type | Click a tick to seek to that event |
| **Prev / Next chunk** | Jump between chunk starts | Click to seek |
| **Filmstrip** (row of blocks below) | One block per event, width ∝ delay (8–72px, scaled by max delay) | Click a block to seek |

## Tick colors

Ticks use the event type to pick a color (same mapping as the filmstrip):

- <span style="display:inline-block;width:10px;height:10px;background:#D1D5DB;border:1px solid #9CA3AF;border-radius:2px;vertical-align:middle"></span> `o` output (text) — zinc
- <span style="display:inline-block;width:10px;height:10px;background:#6B7280;border:1px solid #4B5563;border-radius:2px;vertical-align:middle"></span> `o` output (control-only, e.g. bare newline) — darker zinc
- <span style="display:inline-block;width:10px;height:10px;background:#10B981;border:1px solid #059669;border-radius:2px;vertical-align:middle"></span> `i` input — emerald
- <span style="display:inline-block;width:10px;height:10px;background:#F97316;border:1px solid #EA580C;border-radius:2px;vertical-align:middle"></span> `r` resize — orange
- <span style="display:inline-block;width:10px;height:10px;background:#38BDF8;border:1px solid #0EA5E9;border-radius:2px;vertical-align:middle"></span> `m` marker — sky
- <span style="display:inline-block;width:10px;height:10px;background:#EF4444;border:1px solid #DC2626;border-radius:2px;vertical-align:middle"></span> `x` exit — red

Hover a tick to see `#index [type] @ time — preview`.

Filmstrip blocks use the same palette, plus `⌨` / `↔` / `📍` / `🚪` icons for `i`/`r`/`m`/`x`, and `⏎` for control-only output. Prompt-like outputs get a small sky-blue dot.

## Filmstrip details

- Empty state: *No events — upload a .cast or add events*.
- Active block is highlighted with a sky ring and auto-scrolled into view.
- Clicking a block seeks to that event (`seekTo(index + 1)`).

## Seeking: time vs index

- The slider and track map **time** (`currentTime / totalDuration`) when a duration exists; otherwise they fall back to **index** (`currentIndex / length`).
- `timeToIndex` is cumulative: it finds the first event whose end time ≥ the target time.
- Ticks, filmstrip blocks, chunk segments, and event rows all seek via `seekTo`.
- Seeking clears the terminal and replays events up to the target index, then fits the terminal.

## Keyboard

When focus is not inside an `<input>` or `<textarea>`:

- <kbd>←</kbd> — seek to previous event.
- <kbd>→</kbd> — seek to next event.
- <kbd>Space</kbd> — toggle Play/Pause.

These also work while the Raw tab is active, as long as no input has focus.

## Under the hood (for contributors)

- `recomputeTiming()` builds `cumulativeTimes` and `totalDuration` from `getFlattenedEvents()`.
- `renderTimelineTicks()`, `renderChunkSegments()`, `renderHorizontalFilmstrip()`, and `updateScrubberUI()` are the four renderers; `updateHorizontalTimeline()` calls them after `recomputeTiming()`.
- `index.html:394` is the `HORIZONTAL TIMELINE / SLIDER LOGIC` section.
