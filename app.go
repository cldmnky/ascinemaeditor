//go:build desktop || dev || production || bindings

package main

import (
	"context"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App holds Wails lifecycle state and native dialog bindings.
type App struct {
	ctx context.Context
}

// NewApp creates the application struct.
func NewApp() *App { return &App{} }

func (a *App) startup(ctx context.Context)  { a.ctx = ctx }
func (a *App) shutdown(_ context.Context) {}

// OpenCast opens a native file chooser for *.cast and returns the file content.
// Empty string means the user cancelled. Errors are surfaced to the frontend.
func (a *App) OpenCast() (string, error) {
	if a.ctx == nil {
		return "", nil
	}
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Open Asciinema Cast",
		Filters: []runtime.FileFilter{
			{DisplayName: "Asciinema Cast (*.cast)", Pattern: "*.cast"},
			{DisplayName: "All Files (*.*)", Pattern: "*.*"},
		},
	})
	if err != nil {
		return "", err
	}
	if path == "" {
		return "", nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// SaveCast opens a native save dialog and writes content to the chosen path.
// Returns the saved path or empty string on cancel.
func (a *App) SaveCast(content string) (string, error) {
	if a.ctx == nil {
		return "", nil
	}
	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Save Asciinema Cast",
		DefaultFilename: "edited-recording.cast",
		Filters: []runtime.FileFilter{
			{DisplayName: "Asciinema Cast (*.cast)", Pattern: "*.cast"},
		},
	})
	if err != nil {
		return "", err
	}
	if path == "" {
		return "", nil
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return "", err
	}
	return path, nil
}
