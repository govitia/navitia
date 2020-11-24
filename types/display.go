package types

import (
	"encoding/json"
	"fmt"
	"image/color"
)

// A Display holds informations useful to display.
type Display struct {
	Headsign       string      `json:"headsign"`        // The headsign associated with the object
	Network        string      `json:"network"`         // The name of the belonging network
	Direction      string      `json:"direction"`       // A direction to take
	CommercialMode ID          `json:"commercial_mode"` // The commercial mode in ID Form
	PhysicalMode   ID          `json:"physical_mode"`   // The physical mode in ID Form
	Label          string      `json:"label"`           // The label of the object
	Color          color.Color `json:"color"`           // Hexadecimal color of the line
	TextColor      color.Color `json:"text_color"`      // The text color for this section
	Code           string      `json:"code"`            // The code of the line
	Description    string      `json:"description"`     // Description
	Equipments     []Equipment `json:"equipments"`      // Equipments on this object
	Name           string      `json:"name"`            // Name of object
	TripShortName  string      `json:"trip_short_name"` // TripShoerName short name of the current trip
}

// jsonDisplay define the JSON implementation of Display struct
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonDisplay struct {
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
}

// UnmarshalJSON implements json.Unmarshaller for a Display
func (d *Display) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &jsonDisplay{
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
		return fmt.Errorf("error while unmarshalling Display: %w", err)
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"Display", b}

	// Now process the values
	// We expect a color string length of 6 because it should be coded in hexadecimal
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return gen.err(err, "Color", "color", str, "error in parseColor")
		}
		d.Color = clr
	}
	if str := data.TextColor; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return gen.err(err, "TextColor", "text_color", str, "error in parseColor")
		}
		d.TextColor = clr
	}

	return nil
}
