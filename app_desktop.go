//go:build desktop || dev || production || bindings

package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// runWailsIfTagged starts the Wails app when built with any Wails mode tag
// (desktop, dev, production, bindings). Returns false when built without them
// so main() can fall back to HTTP server mode. With the `bindings` tag,
// wails.Run emits the bindings JSON and exits instead of opening a window.
func runWailsIfTagged() bool {
	app := NewApp()
	err := wails.Run(&options.App{
		Title:  "Asciinema Editor",
		Width:  1400,
		Height: 900,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 24, G: 24, B: 27, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind:             []interface{}{app},
	})
	if err != nil {
		log.Fatal(err)
	}
	return true
}
