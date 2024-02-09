package main

/*

> extism call www/pdk/parse.wasm parse --input "2022-05~" --wasi
Error: wasm error: unreachable
wasm stack trace:
	.runtime.runtimePanicAt(i32,i32)
	.byn$mgfn-shared$runtime.lookupPanic(i32,i32)
	.runtime.nilPanic()
	.(*regexp.Regexp).doExecute(i32,i32,i32,i32,i32,i32,i32)
	.(*regexp.Regexp).MatchString(i32,i32,i32) i32
	.parse() i32

*/

import (
	"log/slog"

	"github.com/extism/go-pdk"
	"github.com/sfomuseum/go-edtf/parser"
)

func main() {}

//export parse
func parse() int32 {

	raw := pdk.InputString()
	d, err := parser.ParseString(raw)

	if err != nil {
		slog.Error("Failed to parse string", "error", err)
		pdk.SetError(err)
		return 1
	}

	pdk.OutputJSON(d)
	return 0
}
