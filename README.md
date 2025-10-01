# Asciinema Web Editor

A web-based editor for asciinema cast files with support for v3 format, chunk-based editing, speed controls, and more.

## Features

- Support for asciinema v3 format with metadata preservation
- Chunk-based editing with collapsible sections
- Speed controls for timing modifications
- Timeline editor with synchronized raw view
- Terminal preview with xterm.js
- Upload/download .cast files
- Comment support in v3 format

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

## Usage

1. Start the server using one of the methods above
2. Open your browser to `http://localhost:8080`
3. Upload a .cast file or start editing
4. Use the Timeline Editor to modify timing and content
5. Preview changes in the terminal
6. Download the modified .cast file

## File Structure

- `index.html` - Main web application
- `main.go` - Go web server
- `go.mod` - Go module file
- `Taskfile.yml` - go-task configuration
- `bin/` - Built binaries (created by build process)
- `test-data/` - Sample cast files for testing
- `.gitignore` - Git ignore file
