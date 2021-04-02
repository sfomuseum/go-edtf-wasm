# go-edtf-wasm

![](docs/images/edtf-parse-wasm.png)

Go package for exposing sfomuseum/go-edtf functionality as WebAssembly binaries.

## Building go-edtf-wasm

The easiest thing is to run the `wasm` Makefile target, like this:

```
$> make wasm
GOOS=js GOARCH=wasm go build -mod vendor -o www/wasm/parse.wasm cmd/parse/main.go
```

This will place a copy of the `parse.wasm` binary in `www/wasm/parse.wasm`.

The binary exposes a single `parse_edtf` function that takes a single string as its input and returns a JSON-encoded [edtf.EDTFDate](https://github.com/sfomuseum/go-edtf#date-spans-or-edtfedtfdate), or nil, as its response.

For example (with error handling omitted for the sake of brevity):

```
var raw_el = document.getElementById("raw");
var edtf_str = raw_el.value;

var result_el = document.getElementById("result");
result_el.innerHTML = "";
    
var rsp = parse_edtf(edtf_str);

if (! rsp){
    return;
}
	
var edtf_d = JSON.parse(rsp)
	
var pre = document.createElement("pre");
pre.innerText = JSON.stringify(edtf_d, '', 2);
	
result_el.appendChild(pre);
```

Obviously returning nil when there is a problem isn't great so the final return value remains to be determined.

## Serving go-edtf-wasm

The package comes with a handy `server` tool for serving the `parse.wasm` binary and a simple web page for parsing EDTF date strings.

```
$> make cli
go build -mod vendor -o bin/server cmd/server/main.go
```

The to run the server:

```
$> ./bin/server 
2021/01/07 17:56:48 Listening on http://localhost:8080
```

All of the HTML, JavaScript and `.wasm` files for the web application are kept in the `www` folder and are bundled with the `bin/server` binary using Go 1.16's [embed](https://tip.golang.org/pkg/embed/) directive.

## See also

* https://github.com/sfomuseum/go-edtf