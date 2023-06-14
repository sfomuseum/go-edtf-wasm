GOROOT=$(shell go env GOROOT)
GOMOD=vendor

rebuild:
	@make wasm

wasi:
	~/sfomuseum/tinygo/build/tinygo build -no-debug -o www/wasi/parse.wasm -target wasi ./cmd/parse-wasi/main.go

wasip:
	GOARCH=wasm GOOS=wasip1 ~/go/bin/gotip build -mod $(GOMOD) -o www/wasi/parse.wasm ./cmd/parse-wasi/main.go

wasm:
	GOOS=js GOARCH=wasm go build -mod $(GOMOD) -ldflags="-s -w" -o www/wasm/parse.wasm cmd/parse/main.go

cli:
	@make wasm
	go build -mod $(GOMOD) -o bin/server cmd/server/main.go

server:
	go run -mod $(GOMOD) cmd/server/main.go
