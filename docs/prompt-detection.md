---
layout: default
title: Prompt Detection — Asciinema Editor
---

<div class="doc-nav" markdown="0">
  <a href="{{ '/' | relative_url }}">Overview</a>
  <a href="{{ '/installation' | relative_url }}">Install</a>
  <a href="{{ '/quickstart' | relative_url }}">Quickstart</a>
  <a href="{{ '/editing' | relative_url }}">Editing</a>
  <a href="{{ '/timeline' | relative_url }}">Timeline</a>
  <a href="{{ '/playback' | relative_url }}">Playback</a>
  <a class="active" href="{{ '/prompt-detection' | relative_url }}">Prompt</a>
  <a href="{{ '/metadata' | relative_url }}">Metadata</a>
  <a href="{{ '/file-format' | relative_url }}">File Format</a>
  <a href="{{ '/troubleshooting' | relative_url }}">Troubleshooting</a>
  <a href="{{ '/releases' | relative_url }}">Releases</a>
</div>

# Prompt Detection

> *Teach the editor to recognize your prompt. It will carve the wasteland into manageable sectors.*

Prompt detection groups raw events into **chunks** — one chunk per shell command. Each chunk gets its own collapsible card, speed control, and comment. Without it, the whole recording is one giant chunk.

![Prompt detection configured for a bash-5.3$ prompt]({{ '/screenshots/prompt-detection.png' | relative_url }})
*`bash-5.3$` in **Prompt pattern** regrouped the sample into 4 chunks (`pwd` → `ls` → `echo "hello"` → `exit`). The filmstrip, chunk segments, and vertical list all update.*

## How it works

The editor looks at every `o` (output) event's `data` and tests it with a single regex:

```js
// index.html:364 — promptCoreSource()
pattern === '' ? '\\$' : escapeRegExp(pattern.trim())
buildPromptRegex() // => new RegExp(core + '\\s*(?:\\x1b\\[[\\d;]*m)*\\s*$')
```

- The **core** is the literal text you type in **Prompt pattern** (escaped so `>`, `$`, `[` etc. are not regex).
- If the field is empty, the core is `\$` — matches any prompt ending in `$`.
- The regex allows trailing whitespace and ANSI color escapes (`\x1b[…m`) before end-of-line, so colored prompts still match.

`isPromptOutput(data)` tests each output event. Editors call `groupEventsIntoChunks(events)`, which starts a new chunk every time the test matches (except for the very first event). The result is `recordingChunks: [{ events, comment, commentIsGenerated, startIndex, speed }, …]`.

## Setting the prompt

1. Open **Cast Metadata (v3)** (collapsed section at the top of the Timeline Editor).
2. Type the literal suffix of your prompt in **Prompt pattern** — e.g. `assets>` for `myapp assets>`, or `❯` for a `❯` prompt.
3. The timeline **regroups immediately**. To tab out, the field listens to `change` (not `input`), so press <kbd>Tab</kbd> or click elsewhere.

Examples:

| Your prompt looks like | Enter in Prompt pattern |
|---|---|
| `bash-5.3$` | `bash-5.3$` |
| `user@host:~$` | `$` (default already matches) |
| `myapp assets>` | `assets>` |
| `❯` | `❯` |
| `PS C:\>` | `>` |

Leave it empty for the default `$` behavior.

## Persistence

The value is stored in `localStorage` under `ascinemaeditor.promptPattern`:

```js
localStorage.getItem('ascinemaeditor.promptPattern')
localStorage.setItem('ascinemaeditor.promptPattern', promptPattern)
```

It survives reloads in the same browser or desktop WebView. It is per-origin — the desktop WebView has its own storage.

## Comments stay anchored

Regrouping does not move `# …` comment lines. The editor keeps `recordingComments: [{ eventIndex, text }]` indexed by event position and re-derives chunk headings via `syncChunkComments()` / `getCommentsByEventIndex()`. So existing comments from an uploaded cast remain at their original event indices even after you change the prompt.

## Tips

- Set the prompt **before** editing — otherwise you will edit one big chunk and then have to re-find your edits after the regroup.
- If your recording has no recognizable prompt (e.g. a raw `script` capture), leave the prompt empty and edit the single chunk, or insert `m` markers to split it manually.
- Colored prompts (with ANSI escapes before the newline) are handled — you do not need to include escape codes in the pattern.

<div class="vault-card" markdown="1">
**Vault-Tec tip:** The sample recording's prompt is `bash-5.3$ `. Type exactly that to see 4 clean chunks. The live screenshot above was taken with that value.
</div>
