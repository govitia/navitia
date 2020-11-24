package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// A PathSegment (called Path item in the Navitia API) is a part of a path
type PathSegment struct {
	Length   uint          `json:"length"`   // The Length of the segment
	Name     string        `json:"name"`     // The Name of the way corresponding to the segment
	Duration time.Duration `json:"duration"` // The duration in seconds of the segment

	// The angle in degree between the previous segment and this segment
	// = 0 Means going straight
	// < 0 Means turning left
	// > 0 Means turning right
	Direction int `json:"direction"`
}

// jsonPathSegment define the JSON implementation of PathSegment struct
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonPathSegment struct {
	// Pointers to the corresponding real values
	Length    *uint
	Name      *string
	Direction *int

	// Value to process
	Duration int64
}

// UnmarshalJSON implements json.Unmarshaller for a PathSegment
func (ps *PathSegment) UnmarshalJSON(b []byte) error {
	data := &jsonPathSegment{
		Length:    &ps.Length,
		Name:      &ps.Name,
		Direction: &ps.Direction,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return fmt.Errorf("error while unmarshalling PathSegment: %w", err)
	}

	// Now process the value
	// As the given duration is in second, let's multiply it by one second to have the correct value
	ps.Duration = time.Duration(data.Duration) * time.Second

	return nil
}
