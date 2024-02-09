package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/aaronland/go-http-bootstrap"
	"github.com/aaronland/go-http-server"
	"github.com/sfomuseum/go-edtf-wasm/www"
	"github.com/sfomuseum/go-http-wasm/v2"
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

	bootstrap_opts := bootstrap.DefaultBootstrapOptions()

	err = bootstrap.AppendAssetHandlers(mux, bootstrap_opts)

	if err != nil {
		log.Fatalf("Failed to append Bootstrap asset handlers, %v", err)
	}

	wasm_opts := wasm.DefaultWASMOptions()

	err = wasm.AppendAssetHandlers(mux, wasm_opts)

	if err != nil {
		log.Fatalf("Failed to append wasm assets handler, %v", err)
	}

	http_fs := http.FS(www.FS)
	fs_handler := http.FileServer(http_fs)

	fs_handler = bootstrap.AppendResourcesHandler(fs_handler, bootstrap_opts)

	fs_handler = wasm.AppendResourcesHandler(fs_handler, wasm_opts)

	mux.Handle("/", fs_handler)

	log.Printf("Listening on %s", s.Address())
	err = s.ListenAndServe(ctx, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
