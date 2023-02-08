GOROOT=$(shell go env GOROOT)

rebuild:
	@make wasmjs
	@make wasm

# Note 'wasi' requires Go 1.18-1.19 (1.20 is not supported by TinyGo yet)
wasi:
	tinygo build -no-debug -o www/wasi/parse.wasm -target wasi ./cmd/parse-wasi/main.go

wasm:
	GOOS=js GOARCH=wasm go build -mod vendor -ldflags="-s -w" -o www/wasm/parse.wasm cmd/parse/main.go

wasmjs:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" www/javascript/

cli:
	@make wasm
	go build -mod vendor -o bin/server cmd/server/main.go
