package json

import (
	gojson "encoding/json"
	"testing"

	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
)

func TestMarshalEDTFDate(t *testing.T) {

	edtf_str := "2000-05~"

	d, err := parser.ParseString(edtf_str)

	if err != nil {
		t.Fatalf("Failed to parse '%s', %v", edtf_str, err)
	}

	enc, err := MarshalEDTFDate(d)

	if err != nil {
		t.Fatalf("Failed to marshal '%s', %v", edtf_str, err)
	}

	var d2 *edtf.EDTFDate

	err = gojson.Unmarshal(enc, &d2)

	if err != nil {
		t.Fatalf("Failed to unmarshal custom JSON, %v", err)
	}
}
