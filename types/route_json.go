package types

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for Route
func (r *Route) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Pointers to the corresponding real values
		ID        *ID        `json:"id"`
		Name      *string    `json:"name"`
		Line      *Line      `json:"line"`
		Direction *Container `json:"direction"`

		// Value to process
		Frequence string `json:"is_frequence"`
	}{
		ID:        &r.ID,
		Name:      &r.Name,
		Line:      &r.Line,
		Direction: &r.Direction,
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"Route", b}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Route.UnmarshalJSON: error while unmarshalling Line")
	}

	// Now process the value
	switch {
	case data.Frequence == "true" || data.Frequence == "True":
		r.Frequence = true
	case data.Frequence == "false" || data.Frequence == "False":
		r.Frequence = false
	default:
		return gen.err(nil, "Frequence", "is_frequency", data.Frequence, `String is neither True, true, False or false`)
	}

	return nil
}
