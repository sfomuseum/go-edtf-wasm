cli:
	~/src/go/bin/go build -mod vendor -o bin/server cmd/server/main.go

wasm:
	GOOS=js GOARCH=wasm ~/src/go/bin/go build -mod vendor -o www/wasm/parse.wasm cmd/parse/main.go
