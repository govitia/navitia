package types

import (
	"encoding/json"
	"fmt"
	"image/color"
)

// Severity object can be used to make visual grouping.
type Severity struct {
	// Name of severity
	Name string `json:"name"`

	// Priority of the severity. Given by the agency. 0 is the strongest priority, a nil Priority means its undefined (duh).
	Priority *int `json:"priority"`

	// HTML color for classification
	Color color.Color `json:"color"`

	// Effect: Normalized value of the effect on the public transport object
	Effect Effect `json:"effect"`
}

// jsonSeverity define the JSON implementation of Severity struct
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonSeverity struct {
	// The references
	Name     *string `json:"name"`
	Priority *int    `json:"priority,omitempty"` // As priority can be null, and 0 is the highest priority.
	Effect   *Effect `json:"effect"`

	// Those we will process
	Color string `json:"color"`
}

// UnmarshalJSON implements json.Unmarshaller for a Severity
func (s *Severity) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &jsonSeverity{
		Name:     &s.Name,
		Priority: s.Priority,
		Effect:   &s.Effect,
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Severity", b}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return fmt.Errorf("error while unmarshalling Severity: %w", err)
	}

	// Process the color
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return gen.err(err, "Color", "color", str, "error in parseColor")
		}
		s.Color = clr
	}

	return nil
}
