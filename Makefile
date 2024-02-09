GOROOT=$(shell go env GOROOT)
GOMOD=$(shell test -f "go.work" && echo "readonly" || echo "vendor")

# TINYGO=tinygo
TINYGO=/usr/local/src/tinygo/build/tinygo

rebuild:
	@make wasm

wasi:
	$(TINYGO) build \
		-target wasi \
		-no-debug \
		-o www/wasi/parse.wasm \
		./cmd/parse-wasi/main.go

pdk:
	$(TINYGO) build \
		-target wasi \
		-no-debug \
		-o www/pdk/parse.wasm \
		cmd/parse-pdk/main.go

# This requires Go 1.21.0 or higher and produces larger binaries that TinyGo (wasi)

wasip:
	GOARCH=wasm GOOS=wasip1 \
		go build -mod $(GOMOD) -ldflags="-s -w" \
		-o www/wasip/parse.wasm \
		./cmd/parse-wasi/main.go

wasm:
	GOOS=js GOARCH=wasm \
		go build -mod $(GOMOD) -ldflags="-s -w" \
		-o www/wasm/parse.wasm \
		cmd/parse/main.go

cli:
	@make wasm
	go build -mod $(GOMOD) -o bin/server cmd/server/main.go

server:
	go run -mod $(GOMOD) cmd/server/main.go
