package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-server"
	"log"
	"net/http"
	"path/filepath"
)

func main() {

	root := flag.String("root", "www", "...")
	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	flag.Parse()

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	abs_root, err := filepath.Abs(*root)

	if err != nil {
		log.Fatalf("Failed to derive absolute path for root '%s', %v", abs_root, err)
	}
	
	http_root := http.Dir(abs_root)
	fs_handler := http.FileServer(http_root)
	
	mux := http.NewServeMux()
	mux.Handle("/", fs_handler)

	log.Printf("Listening on %s", s.Address())
	s.ListenAndServe(ctx, mux)
}
