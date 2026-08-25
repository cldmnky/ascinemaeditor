# Asciinema Web Editor

A browser-based editor for asciinema cast files. Edit event timing and output, preview the recording in a terminal, then download an updated `.cast` file.

## Features

- Support for asciinema v3 format with metadata preservation
- Chunk-based editing with collapsible sections
- Speed controls for timing modifications
- Timeline editor with synchronized raw view
- Terminal preview with xterm.js
- Upload/download .cast files
- Comment support in v3 format
- Configurable prompt detection for chunk splitting (e.g. `$` or `assets>`)

## Editor Overview

![Timeline editor beside the terminal preview](docs/screenshots/editor-overview.png)

The left panel is the timeline editor. The right panel provides terminal playback and a synchronized Raw `.cast` view.

## Running the Server

### Prerequisites

- Go 1.21 or later
- [go-task](https://taskfile.dev/) for task automation

### Installing go-task

Install go-task from source:

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

Make sure your `$GOPATH/bin` (or `$GOBIN`) directory is in your `$PATH` so you can run the `task` command.

### Using go-task (Recommended)

```bash
# Build and run the server
task serve

# Or just run the default task
task

# For development with file watching
task dev

# Clean build artifacts
task clean
```

### Alternative: Using Go directly

```bash
# Build the server
mkdir -p bin
go build -o bin/ascinemaeditor main.go

# Run the server
./bin/ascinemaeditor

# Or with custom port
PORT=3000 ./bin/ascinemaeditor
```

### Environment Variables

- `PORT`: Server port (default: 8080)
- `HOST`: Bind address (default: `127.0.0.1`; set `0.0.0.0` to expose on the network)

## Getting Started

1. Start the server with `task serve` or `go run main.go` from the repository root.
2. Open `http://localhost:8080` in a browser.
3. Select **Upload .cast** and choose an asciinema recording.
4. Edit the timeline, then select **Play** to check the result in the terminal preview.
5. Select **Download .cast** to save `edited-recording.cast`.

The editor accepts asciinema v3 recordings. It also imports v2 recordings and converts their absolute event timestamps to v3 interval delays when loaded.

## Using the Editor

### Edit Events

Each timeline row is one cast event.

- Change the delay field to set the interval before that event. Delays must be non-negative numbers.
- Change editable content fields to update terminal output, input, or marker text.
- Use the checkbox to choose events affected by **Speed Up**.
- Select **📍** on an event to insert a marker after it, or **Add Marker** to add one at the start of a chunk.
- Select the trash icon to remove an event. Select **Add Event** to append a blank output event.

Some control-only output events, such as a bare newline or bell, are displayed as disabled special-character rows. Output that contains visible text remains editable even when it ends with a newline.

### Preview and Raw View

- **Play** replays the edited cast with its current delays.
- Select an event row to seek the preview through that event.
- **Reset** clears playback and returns to the first event.
- Select **Raw .cast** to inspect the exact serialized event stream. Comments remain in the same position they will have in the downloaded file.

### Metadata and Comments

Open **Cast Metadata (v3)** to edit the recording title, command, and tags. Tags are entered as comma-separated values.

Chunk headings provide an editable primary comment. Existing comments from an uploaded cast are preserved in their original positions, including multiple comments before one event and comments between events.

## Prompt Detection

Prompt detection groups events into command-sized chunks. The default prompt is `$`; set a different prompt when the recording uses another shell or CLI suffix.

![Prompt detection configured for an assets prompt](docs/screenshots/prompt-detection.png)

1. Open **Cast Metadata (v3)**.
2. Enter the literal prompt suffix in **Prompt**. For example, use `assets>` for a prompt ending in `assets>`, or leave it empty for the default `$` behavior.
3. The timeline regroups immediately. Existing cast comments remain anchored to their original events.

The prompt setting is stored in the browser's local storage, so it is reused the next time the editor is opened in that browser.

## Troubleshooting

- The UI is embedded in the binary, so the server can be started from any directory. It serves only `/`; every other path returns 404.
- The terminal preview requires access to the xterm.js and font CDNs. Editing and downloading still work without a terminal preview, but the page may not initialize correctly if those scripts cannot load.
- Invalid cast headers or events are rejected without replacing the recording currently open in the editor. Check that each event is a JSON array containing a non-negative delay, string event type, and string data.

## Releases

Prebuilt binaries for linux/darwin (amd64 + arm64) are published on the [releases page](https://github.com/cldmnky/ascinemaeditor/releases). Download one for your platform and run it directly; the editor ships with built-in sample data so you can start editing immediately. Each release includes SHA256 checksums.

To cut a release, tag `v*` and push the tag:

```bash
git tag v0.1.0 && git push origin v0.1.0
```

## File Structure

- `index.html` - Main web application (embedded into the binary)
- `main.go` - Go web server serving the embedded UI
- `.github/workflows/release.yml` - Tag-triggered release build and publishing
- `go.mod` - Go module file
- `Taskfile.yml` - go-task configuration
- `bin/` - Built binaries (created by build process)
- `docs/screenshots/` - README screenshots
- `test-data/` - Sample cast files for testing
- `.gitignore` - Git ignore file
