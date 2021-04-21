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
    result_el.style.display = "none";
    
    result_el.innerHTML = "";
    
    parse_edtf(edtf_str).then(rsp => {
	
	try {
	    var edtf_d = JSON.parse(rsp)
	} catch(e){
	    result_el.innerText = "Unable to parse your EDTF string: " + e;
	    
	    result_el.style.display = "block";
	    return;
	}
	
	var pre = document.createElement("pre");
	pre.innerText = JSON.stringify(edtf_d, '', 2);
	
	result_el.appendChild(pre);
	result_el.style.display = "block";	    	    	
	
    }).catch(err => {
	result_el.innerText = "There was a problem parsing your EDTF string:" + err;
	result_el.style.display = "block";    	
    });
    
    return false;
}
