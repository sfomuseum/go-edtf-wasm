package tinygo

import (
	"fmt"

	"github.com/sfomuseum/go-edtf"
)

func MarshalPrecision(p edtf.Precision) ([]byte, error) {
	v := fmt.Sprintf("%v", p)
	return []byte(v), nil
}
