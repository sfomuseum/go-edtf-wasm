GOROOT=$(shell go env GOROOT)
GOMOD=vendor

rebuild:
	@make wasm

# Note 'wasi' requires Go 1.18-1.19 (1.20 is not supported by TinyGo yet)
wasi:
	tinygo build -no-debug -o www/wasi/parse.wasm -target wasi ./cmd/parse-wasi/main.go

wasip:
	GOARCH=wasm GOOS=wasip1 ~/go/bin/gotip build -mod $(GOMOD) -o www/wasi/parse.wasm ./cmd/parse-wasi/main.go

wasm:
	GOOS=js GOARCH=wasm go build -mod $(GOMOD) -ldflags="-s -w" -o www/wasm/parse.wasm cmd/parse/main.go

cli:
	@make wasm
	go build -mod $(GOMOD) -o bin/server cmd/server/main.go
