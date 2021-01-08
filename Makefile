cli:
	go build -mod vendor -o bin/server cmd/server/main.go

wasm:
	GOOS=js GOARCH=wasm go build -mod vendor -o www/wasm/parse.wasm cmd/parse/main.go
