package types

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for a DisplayInformations
func (di *DisplayInformations) UnmarshalJSON(b []byte) error {
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
		Headsign:       &di.Headsign,
		Network:        &di.Network,
		Direction:      &di.Direction,
		CommercialMode: &di.CommercialMode,
		PhysicalMode:   &di.PhysicalMode,
		Label:          &di.Label,
		Code:           &di.Code,
		Description:    &di.Description,
		Equipments:     &di.Equipments,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "DisplayInformations.UnmarshalJSON: error while unmarshalling Line")
	}

	// Now process the value
	// We expect a color string length of 6 because it should be coded in hexadecimal
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return errors.Wrapf(err, "DisplayInformations.UnmarshalJSON: error while parsing color (given \"color\":\"%s\")", str)
		}
		di.Color = clr
	}
	if str := data.TextColor; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return errors.Wrapf(err, "DisplayInformations.UnmarshalJSON: error while parsing text color (given \"color\":\"%s\")", str)
		}
		di.TextColor = clr
	}

	return nil
}
