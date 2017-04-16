package types

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for a Line
func (l *Line) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Pointers to the corresponding real values
		ID             *ID             `json:"id"`
		Name           *string         `json:"name"`
		Code           *string         `json:"code"`
		OpeningTime    *string         `json:"opening_time"`
		ClosingTime    *string         `json:"closing_time"`
		Routes         *[]Route        `json:"routes"`
		CommercialMode *CommercialMode `json:"commercial_mode"`
		PhysicalModes  *[]PhysicalMode `json:"physical_modes"`

		// Value to process
		Color string `json:"color"`
	}{
		ID:             &l.ID,
		Name:           &l.Name,
		Code:           &l.Code,
		OpeningTime:    &l.OpeningTime,
		ClosingTime:    &l.ClosingTime,
		Routes:         &l.Routes,
		CommercialMode: &l.CommercialMode,
		PhysicalModes:  &l.PhysicalModes,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Line.UnmarshalJSON: error while unmarshalling Line")
	}

	// Now process the value
	// We expect a color string length of 6 because it should be coded in hexadecimal
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return errors.Wrapf(err, "Line.UnmarshalJSON: error while parsing color (given \"color\":\"%s\")", str)
		}
		l.Color = clr
	}

	return nil
}
