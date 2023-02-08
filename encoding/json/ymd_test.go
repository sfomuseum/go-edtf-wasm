package json

import (
	"testing"

	"github.com/sfomuseum/go-edtf"
)

func TestYMD(t *testing.T) {

	ymd := &edtf.YMD{
		Year:  2022,
		Month: 1,
		Day:   15,
	}

	b, err := MarshalYMD(ymd)

	if err != nil {
		t.Fatalf("Failed to encode ymd, %v", err)
	}

	expected := `{"day": 15, "month": 1, "year": 2022}`

	if string(b) != expected {
		t.Fatalf("Unexpected marshaling, '%s'. Expected '%s'.", string(b), expected)
	}
}
