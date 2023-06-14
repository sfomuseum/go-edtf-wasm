GOROOT=$(shell go env GOROOT)
GOMOD=vendor

rebuild:
	@make wasm

# Note 'wasi' requires TinyGo 0.28.0 or highter
wasi:
	tinygo build -no-debug -o www/wasi/parse.wasm -target wasi ./cmd/parse-wasi/main.go

wasm:
	GOOS=js GOARCH=wasm go build -mod $(GOMOD) -ldflags="-s -w" -o www/wasm/parse.wasm cmd/parse/main.go

cli:
	@make wasm
	go build -mod $(GOMOD) -o bin/server cmd/server/main.go
