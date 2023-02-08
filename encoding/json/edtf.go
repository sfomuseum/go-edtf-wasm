package json

import (
	"fmt"

	"github.com/sfomuseum/go-edtf"
)

func MarshalEDTFDate(d *edtf.EDTFDate) ([]byte, error) {

	start, err := MarshalDateRange(d.Start)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal starting date range, %w", err)
	}
	
	end, err := MarshalDateRange(d.End)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal ending date range, %w", err)
	}
	
	v := fmt.Sprintf(`{"edtf": "%s", "end": %s, "feature": "%s", "level": %d, "start": %s}`, d.EDTF, string(end), d.Feature, d.Level, string(start))
	return []byte(v), nil
}
