build:
	GOOS=js GOARCH=wasm go build -mod vendor -o www/parse.wasm cmd/parse/main.go
