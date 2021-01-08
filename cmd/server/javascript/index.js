if (! WebAssembly.instantiateStreaming){
	 
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();

let mod, inst;

WebAssembly.instantiateStreaming(fetch("/wasm/parse.wasm"), go.importObject).then(
    
    async result => {
	document.getElementById("button").innerText = "Parse";
	document.getElementById("button").removeAttribute("disabled");
        mod = result.module;
        inst = result.instance;
	await go.run(inst);
    }
);

async function parse() {
    
    var raw_el = document.getElementById("raw");
    var edtf_str = raw_el.value;

    var result_el = document.getElementById("result");
    result_el.innerHTML = "";
    
    try {
	
	var rsp = parse_edtf(edtf_str);

	if (! rsp){
	    result_el.innerText = "There was a problem parsing your EDTF string.";
	    return;
	}
	
	var edtf_d = JSON.parse(rsp)
	
	var pre = document.createElement("pre");
	pre.innerText = JSON.stringify(edtf_d, '', 2);
	
	result_el.appendChild(pre);
	
    } catch (err) {
	result_el.innerText = "Unable to parse your EDTF string: " + err;	
    }

    return false;
}
