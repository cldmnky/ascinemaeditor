package main

import (
	_ "embed"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

//go:embed index.html
var indexHTML []byte

func main() {
	// Browser mode is kept for quick iteration: `go run . serve` or `PORT=... go run . serve`.
	// `go run -tags desktop .` or `wails dev`/`wails build` starts the desktop app — no browser needed.
	if isServeMode() {
		runServer()
		return
	}
	if runWailsIfTagged() {
		return
	}
	runServer()
}

func isServeMode() bool {
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		return true
	}
	return os.Getenv("ASCINEMAEDITOR_MODE") == "serve"
}

func runServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		_, _ = w.Write(indexHTML)
	})

	server := &http.Server{
		Addr:              net.JoinHostPort(host, port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	fmt.Printf("Starting server on http://%s  (desktop is the default with -tags desktop — `wails dev`/`wails build` or `go run -tags desktop .`; `go run . serve` for browser mode)\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
