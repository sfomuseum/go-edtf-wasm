package tinygo

import (
	"fmt"

	"github.com/sfomuseum/go-edtf"
)

func MarshalDate(d *edtf.Date) ([]byte, error) {

	approximate, err := MarshalPrecision(d.Approximate)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal approximate value, %w", err)
	}

	unspecified, err := MarshalPrecision(d.Unspecified)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal unspecified value, %w", err)
	}

	uncertain, err := MarshalPrecision(d.Uncertain)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal uncertain value, %w", err)
	}

	precision, err := MarshalPrecision(d.Precision)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal precision value, %w", err)
	}

	inclusivity, err := MarshalPrecision(d.Inclusivity)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal inclusivity value, %w", err)
	}

	timestamp, err := MarshalTimestamp(d.Timestamp)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal timestamp value, %w", err)
	}

	ymd, err := MarshalYMD(d.YMD)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal ymd value, %w", err)
	}

	v := fmt.Sprintf(`{"approximate": %t, "datetime": %s, "inclusivity": %s, "open": %t, "precision": %s", "timestamp": %s, "uncertain": %s, "unknown": %t, "unspecified": %s, "ymd": %s}`, string(approximate), d.DateTime, string(inclusivity), d.Open, string(precision), string(timestamp), string(uncertain), d.Unknown, string(unspecified), string(ymd))

	return []byte(v), nil
}
