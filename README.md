# go-edtf-wasm

![](docs/images/go-edtf-wasm-bootstrap.png)

Go package for exposing sfomuseum/go-edtf functionality as WebAssembly binaries.

## WASM

### Building

The easiest thing is to run the `wasm` Makefile target, like this:

```
$> make wasm
GOOS=js GOARCH=wasm go build -mod vendor -o www/wasm/parse.wasm cmd/parse/main.go
```

This will place a copy of the `parse.wasm` binary in `www/wasm/parse.wasm`.

The binary exposes a single `parse_edtf` function that takes a single string as its input and returns a JavaScript promise. The promise returns a JSON-encoded [edtf.EDTFDate](https://github.com/sfomuseum/go-edtf#date-spans-or-edtfedtfdate) if successful and an error string if not.

For example (with error handling omitted for the sake of brevity):

```
var raw_el = document.getElementById("raw");
var edtf_str = raw_el.value;

var result_el = document.getElementById("result");
result_el.innerHTML = "";
    
parse_edtf(edtf_str).then(rsp => {

	var edtf_d = JSON.parse(rsp)
	
	var pre = document.createElement("pre");
	pre.innerText = JSON.stringify(edtf_d, '', 2);
	
	result_el.appendChild(pre);
	
}).catch(err => {
	console.log("Failed to parse EDTF string", err)
});
```

### Serving go-edtf-wasm

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

## WASI

### Building

The easiest thing is to run the `wasi` Makefile target, like this:

```
$> make wasi
tinygo build -no-debug -o www/wasi/parse.wasm -target wasi ./cmd/parse-wasi/main.go
```

This will place a copy of the `parse.wasm` binary in `www/wasi/parse.wasm`.

Note that this requires having a copy of [TinyGo](https://tinygo.org/) installed and findable in your local path.

### Python

There is a still-experimental Python script for running the ``www/wasi/parse.wasm` binary in the `python` directory. It depends on the [wasmer-python](https://github.com/wasmerio/wasmer-python) libraries already being installed.

```
$> /usr/local/opt/python@3.10/bin/python3.10 ./python/parse.py -h
Usage: parse.py [options]

Options:
  -h, --help            show this help message and exit
  -e EDTF, --edtf=EDTF  The EDTF string to parse
  -w WASI, --wasi=WASI  The path to the WASI binary to compile
```

If successful the program will emit a JSON-encoded [edtf.EDTFDate](https://pkg.go.dev/github.com/sfomuseum/go-edtf#EDTFDate) struct to STDOUT. If not successful the program will emit an error message to STDOUT. Better error handling and reporting is expected to follow shortly.

For example:

```
$> /usr/local/opt/python@3.10/bin/python3.10 ./python/parse.py \
	--wasi ./www/wasi/parse.wasm \
	--edtf '2022-05~'
	
	| jq
	
{
  "edtf": "2022-05~",
  "end": {
    "edtf": "2022-05-31",
    "lower": {
      "approximate": 0,
      "datetime": "2022-05-31T00:00:00Z",
      "inclusivity": 0,
      "open": false,
      "precision": 64,
      "timestamp": 1653955200,
      "uncertain": 0,
      "unknown": false,
      "unspecified": 0,
      "ymd": {
        "day": 31,
        "month": 5,
        "year": 2022
      }
    },
    "upper": {
      "approximate": 0,
      "datetime": "2022-05-31T23:59:59Z",
      "inclusivity": 0,
      "open": false,
      "precision": 64,
      "timestamp": 1654041599,
      "uncertain": 0,
      "unknown": false,
      "unspecified": 0,
      "ymd": {
        "day": 31,
        "month": 5,
        "year": 2022
      }
    }
  },
  "feature": "Seasons",
  "level": 1,
  "start": {
    "edtf": "2022-05-01",
    "lower": {
      "approximate": 0,
      "datetime": "2022-05-01T00:00:00Z",
      "inclusivity": 0,
      "open": false,
      "precision": 64,
      "timestamp": 1651363200,
      "uncertain": 0,
      "unknown": false,
      "unspecified": 0,
      "ymd": {
        "day": 1,
        "month": 5,
        "year": 2022
      }
    },
    "upper": {
      "approximate": 0,
      "datetime": "2022-05-01T23:59:59Z",
      "inclusivity": 0,
      "open": false,
      "precision": 64,
      "timestamp": 1651449599,
      "uncertain": 0,
      "unknown": false,
      "unspecified": 0,
      "ymd": {
        "day": 1,
        "month": 5,
        "year": 2022
      }
    }
  }
}
```

## See also

* https://github.com/sfomuseum/go-edtf
* https://github.com/golang/go/wiki/WebAssembly
* https://github.com/aaronland/go-http-server
* https://github.com/aaronland/go-http-bootstrap