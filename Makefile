GOROOT=$(shell go env GOROOT)
GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

rebuild:
	@make wasm

wasi:
	tinygo build -no-debug -o www/wasi/parse.wasm -target wasi ./cmd/parse-wasi/main.go

wasip:
	GOARCH=wasm GOOS=wasip1 go build -mod $(GOMOD) -ldflags="-s -w" -o www/wasi/parse.wasm ./cmd/parse-wasi/main.go

wasm:
	GOOS=js GOARCH=wasm go build -mod $(GOMOD) -ldflags="-s -w" -o www/wasm/parse.wasm cmd/parse/main.go

cli:
	@make wasm
	go build -mod $(GOMOD) -o bin/server cmd/server/main.go

server:
	go run -mod $(GOMOD) cmd/server/main.go
