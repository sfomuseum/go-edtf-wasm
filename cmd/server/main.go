package main

import (
	"context"
	"embed"
	"flag"
	"github.com/aaronland/go-http-server"
	"log"
	"net/http"
)

//go:embed index.html
//go:embed wasm
//go:embed javascript
var web_app embed.FS

func main() {

	server_uri := flag.String("server-uri", "http://localhost:8080", "A valid aaronland/go-http-server URI.")

	flag.Parse()

	ctx := context.Background()

	s, err := server.NewServer(ctx, *server_uri)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	mux := http.NewServeMux()

	http_fs := http.FS(web_app)
	fs_handler := http.FileServer(http_fs)

	mux.Handle("/", fs_handler)

	log.Printf("Listening on %s", s.Address())
	s.ListenAndServe(ctx, mux)
}
