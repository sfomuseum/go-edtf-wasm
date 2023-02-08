package main

import (
        "flag"
        "fmt"
        _ "encoding/json"
	
        "github.com/sfomuseum/go-edtf/parser"
        "github.com/sfomuseum/go-edtf-wasm/encoding/json"	
)

func main(){

        flag.Parse()

        for _, raw := range flag.Args(){
                fmt.Println(parse(raw))
        }
}

//export parse
func parse(raw string) string {

        d, err := parser.ParseString(raw)

        if err != nil {
                return err.Error()
        } else {
	
		v, err := json.MarshalEDTFDate(d)

		if err != nil {
			return err.Error()
		}

		return string(v)
		
		/*

		return d.String()

		Basically the go-edtf package needs to export it's own JSON because encoding/json is
		not supported by tinygo yet

> /usr/local/opt/python@3.10/bin/python3.10 ./test-wasmer.py
Traceback (most recent call last):
  File "/usr/local/sfomuseum/go-edtf-wasm/python/./test-wasmer.py", line 111, in <module>
    instance.exports._start()
RuntimeError: RuntimeError: unreachable
    at runtime._panic (<module>[28]:0x184d)
    at (reflect.rawType).Name (<module>[52]:0x4524)

		enc, err := json.Marshal(d)

		if err != nil {
			return err.Error()
		}
		
                return string(enc)
		*/
        }

}
