GOROOT=$(shell go env GOROOT)

rebuild:
	@make wasmjs
	@make wasm

wasm:
	GOOS=js GOARCH=wasm go build -mod vendor -o www/wasm/parse.wasm cmd/parse/main.go

wasmjs:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" www/javascript/

cli:
	@make wasm
	go build -mod vendor -o bin/server cmd/server/main.go
