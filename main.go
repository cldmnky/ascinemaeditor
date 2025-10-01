package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get the directory where the executable is located
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files from the current directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	fmt.Printf("Starting server on http://localhost:%s\n", port)
	fmt.Printf("Serving files from: %s\n", dir)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
