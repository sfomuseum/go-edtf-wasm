GO=~/src/go/bin/go

cli:
	@make wasm
	$(GO) build -mod vendor -o bin/server cmd/server/main.go

wasm:
	GOOS=js GOARCH=wasm $(GO) build -mod vendor -o www/wasm/parse.wasm cmd/parse/main.go
