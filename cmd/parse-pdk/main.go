package main

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
