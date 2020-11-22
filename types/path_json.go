package types

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for a PathSegment.
func (ps *PathSegment) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Pointers to the corresponding real values
		Length    *uint
		Name      *string
		Direction *int

		// Value to process
		Duration int64
	}{
		Length:    &ps.Length,
		Name:      &ps.Name,
		Direction: &ps.Direction,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Now process the value
	// As the given duration is in second, let's multiply it by one second to have the correct value
	ps.Duration = time.Duration(data.Duration) * time.Second

	return nil
}
