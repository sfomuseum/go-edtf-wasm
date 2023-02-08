package tinygo

import (
	"fmt"

	"github.com/sfomuseum/go-edtf"
)

func MarshalDateRange(d *edtf.DateRange) ([]byte, error) {

	lower, err := MarshalDate(d.Lower)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal lower date, %w", err)
	}

	upper, err := MarshalDate(d.Upper)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal upper date, %w", err)
	}

	v := fmt.Sprintf(`{"edtf": %s, "lower": %s, "upper": %s}`, d.EDTF, string(lower), string(upper))
	return []byte(v), nil
}
