package main

import (
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// restrictedFileSystem wraps http.Dir to block dotfiles, dot directories
// (.git in particular) and directory listings. Directories that contain an
// index.html are still served normally so "/" resolves to index.html.
type restrictedFileSystem struct {
	root http.FileSystem
}

func (r restrictedFileSystem) Open(name string) (http.File, error) {
	clean := path.Clean("/" + name)
	for _, segment := range strings.Split(clean, "/") {
		if strings.HasPrefix(segment, ".") {
			return nil, fs.ErrPermission
		}
	}

	file, err := r.root.Open(clean)
	if err != nil {
		return nil, err
	}

	info, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}
	if info.IsDir() {
		index, indexErr := r.root.Open(path.Join(clean, "index.html"))
		if index != nil {
			index.Close()
		}
		if indexErr != nil {
			file.Close()
			return nil, fs.ErrNotExist
		}
	}

	return file, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	root, err := filepath.Abs(".")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(restrictedFileSystem{root: http.Dir(root)}))

	server := &http.Server{
		Addr:              net.JoinHostPort(host, port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	fmt.Printf("Starting server on http://%s\n", server.Addr)
	fmt.Printf("Serving files from: %s\n", root)

	log.Fatal(server.ListenAndServe())
}
