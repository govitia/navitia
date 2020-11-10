package types

import "image/color"

// A Display holds informations useful to display.
type Display struct {
	// The headsign associated with the object
	Headsign string `json:"headsign"`

	// The name of the belonging network
	Network string `json:"network"`

	// A direction to take
	Direction string `json:"direction"`

	// The commercial mode in ID Form
	CommercialMode ID `json:"commercial_mode"`

	// The physical mode in ID Form
	PhysicalMode ID `json:"physical_mode"`

	// The label of the object
	Label string `json:"label"`

	// Hexadecimal color of the line
	Color color.Color `json:"color"`

	// The text color for this section
	TextColor color.Color `json:"text_color"`

	// The code of the line
	Code string `json:"code"`

	// Description
	Description string `json:"description"`

	// Equipments on this object
	Equipments []Equipment `json:"equipments"`

	// Name of object
	Name string `json:"name"`

	// TripShoerName short name of the current trip
	TripShortName string `json:"trip_short_name"`
}
