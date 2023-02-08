package tinygo

import (
	"fmt"

	"github.com/sfomuseum/go-edtf"
)

func MarshalYMD(ymd *edtf.YMD) ([]byte, error) {
	v := fmt.Sprintf(`{"day": %d, "month": %d, "year": %d}`, ymd.Day, ymd.Month, ymd.Year)
	return []byte(v), nil
}
