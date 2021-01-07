package main

import (
	"context"
	"flag"
	"github.com/aaronland/go-http-server"
	"log"
	"net/http"
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
	// mux.Handle("/", NewHandler())

	log.Printf("Listening on %s", s.Address())
	s.ListenAndServe(ctx, mux)
}
