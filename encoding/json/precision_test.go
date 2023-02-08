package json

import (
	"fmt"
	"testing"

	"github.com/sfomuseum/go-edtf"
)

func TestPrecision(t *testing.T) {

	pr := edtf.Precision(edtf.MONTH)
	b, err := MarshalPrecision(pr)

	if err != nil {
		t.Fatalf("Failed to encode precision, %v", err)
	}

	if string(b) != fmt.Sprintf("%v", edtf.MONTH) {
		t.Fatalf("Unexpected marshaling, %s", string(b))
	}
}
