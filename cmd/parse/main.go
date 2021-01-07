package main

import (
	"encoding/json"
	"github.com/sfomuseum/go-edtf/parser"
	"log"
	"syscall/js"
)

var parse_func js.Func

func main() {

	parse_func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		if len(args) != 1 {
			log.Println("Invalid arguments")
			return nil
		}

		edtf_str := args[0].String()

		edtf_d, err := parser.ParseString(edtf_str)

		if err != nil {
			log.Printf("Failed to parse '%s', %v\n", edtf_str, err)
			return nil
		}

		enc, err := json.Marshal(edtf_d)

		if err != nil {
			log.Printf("Failed to marshal result for '%s', %v\n", edtf_str, err)
			return nil
		}

		return string(enc)
	})

	defer parse_func.Release()

	js.Global().Set("parse_edtf", parse_func)

	c := make(chan struct{}, 0)

	log.Println("WASM EDTF parser initialized")
	<-c
}
