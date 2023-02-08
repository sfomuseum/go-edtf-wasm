package json

import (
	"github.com/sfomuseum/go-edtf"
)

func MarshalTimestamp(ts *edtf.Timestamp) ([]byte, error) {
	return ts.MarshalJSON()
}
