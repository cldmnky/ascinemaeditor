---
layout: default
title: Metadata & Comments — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
  <a href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a href="{{ '/editing' | relative_url }}">Editing</a>
  <a href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a href="{{ '/playback' | relative_url }}">Playback</a>
  <a href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a class="active" href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Metadata & Comments

> *Archive notes survive the blast. Title it, tag it, comment it.*

## Cast Metadata (v3)

At the top of the **Timeline Editor** is the collapsible **Cast Metadata (v3)** section:

| Field | Header key | Notes |
|---|---|---|
| **Title** | `title` | Free text. Empty → key is deleted on save. |
| **Command** | `command` | Free text. Empty → deleted. |
| **Tags** | `tags` | Comma-separated. `tag1, tag2, demo` → `["tag1","tag2","demo"]`. Empty → deleted. Handles both string and array on load. |
| **Prompt pattern** | *(local only)* | Not written to the cast — stored in `localStorage` (see [Prompt Detection]({{ '/prompt-detection' | relative_url }})). |

`loadMetadataToForm()` fills the inputs from `recordingHeader` on load; `saveMetadataFromForm()` writes them back into the header at save time. `normalizeHeaderToV3()` ensures `version: 3` and a valid `term: { cols, rows }` (defaults 80×24).

## Chunk comments

Each chunk header has an editable **comment** input, initialized as `# Chunk N: …` (e.g. `# Chunk 1: Command - pwd`). The text is stored as `chunk.comment` and whether it was auto-generated as `chunk.commentIsGenerated`. Editing the input calls `setChunkComment()`:

- Non-empty → `chunk.comment = "# " + value`, `commentIsGenerated = false`, and the backing `recordingComments` entry at `chunk.startIndex` is updated or created.
- Empty → the comment entry is removed.

Chunk comments are the **primary** comment for a chunk — the one that appears right before its first event in the output file.

## `#` comments (v3)

Asciinema v3 allows `# …` comment lines interleaved with events. The editor fully preserves them:

- **On load**, `parseCastText()` splits lines into `{ type: 'comment', text }` vs `{ type: 'event', obj }`. Comments are stored in `recordingComments: [{ eventIndex, text }]` where `eventIndex` is the event they precede (0 = before first event).
- **In Raw view**, `getCommentsByEventIndex()` buckets by index and `renderRawView()` emits comments in position.
- **On save**, `buildCastContent()` / `getCommentsByEventIndex()` interleave `JSON.stringify(event)` and `comment` lines in order.
- **On edit** (add/remove/rechunk), helpers `shiftCommentsForInsertion` / `shiftCommentsForRemoval` adjust indices so comments stay anchored to the same logical point.
- **Multiple comments** before one event and comments **between events** are supported.

Chunk headings absorb the first comment at `chunk.startIndex` (if present) as their display comment; `syncChunkComments()` keeps this in sync after structural edits. If a chunk has no comment at its start, its heading is `null` and the input shows the placeholder `Chunk N`.

## Tips

- Use chunk comments to label commands (`# demo: create cluster`) — they survive round-trips and show up in Raw view.
- Check **Raw .cast** before saving — it is the exact serialization, comments included.
- Tags are meant for your own organization; `asciinema` players ignore them.
