package types

import (
	"encoding/json"

	"github.com/aabizri/navitia/internal/unmarshal"
	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for a Display
func (d *Display) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Pointers to the corresponding real values
		Headsign       *string      `json:"headsign"`
		Network        *string      `json:"network"`
		Direction      *string      `json:"direction"`
		CommercialMode *ID          `json:"commercial_mode"`
		PhysicalMode   *ID          `json:"physical_mode"`
		Label          *string      `json:"label"`
		Code           *string      `json:"code"`
		Description    *string      `json:"description"`
		Equipments     *[]Equipment `json:"equipments"`

		// Values to process
		Color     string `json:"color"`
		TextColor string `json:"text_color"`
	}{
		Headsign:       &d.Headsign,
		Network:        &d.Network,
		Direction:      &d.Direction,
		CommercialMode: &d.CommercialMode,
		PhysicalMode:   &d.PhysicalMode,
		Label:          &d.Label,
		Code:           &d.Code,
		Description:    &d.Description,
		Equipments:     &d.Equipments,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Display.UnmarshalJSON: error while unmarshalling Line")
	}

	// Create the error generator
	gen := unmarshal.NewGenerator("Display", &b)
	defer gen.Close()

	// Now process the values
	// We expect a color string length of 6 because it should be coded in hexadecimal
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return gen.Gen(err, "Color", "color", str, "error in parseColor")
		}
		d.Color = clr
	}
	if str := data.TextColor; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return gen.Gen(err, "TextColor", "text_color", str, "error in parseColor")
		}
		d.TextColor = clr
	}

	return nil
}
