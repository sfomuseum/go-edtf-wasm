package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-server"
	// "github.com/aaronland/go-http-bootstrap"
	"github.com/sfomuseum/go-edtf-wasm/www"	
	"net/http"
	"log"
)

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	flag.Parse()

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	mux := http.NewServeMux()

	http_fs := http.FS(www.FS)
	fs_handler := http.FileServer(http_fs)

	mux.Handle("/", fs_handler)

	log.Printf("Listening on %s", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
