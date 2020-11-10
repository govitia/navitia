package types

import (
	"image/color"
)

// A Line codes for a public transit line.
//
// Warning: a Line isn't a route, it has no direction information, and can have several embranchments.
//
// See http://doc.navitia.io/#public-transport-objects
type Line struct {
	// ID is the navitia identifier of the line
	// For example: "line:RAT:M6"
	ID ID `json:"id"`

	// Name is the name of the line
	// For example: "Nation - Charles de Gaule Etoile"
	Name string `json:"name"`

	// Code is the codename of the line.
	// For example: "6"
	Code string `json:"code"`

	// Color is the color given to the line
	// For example: "79BB92" in Hex
	Color color.Color `json:"color"`

	// OpeningTime is the opening time of the line
	OpeningTime struct {
		Hours   uint8 `json:"hours"`
		Minutes uint8 `json:"minutes"`
		Seconds uint8 `json:"seconds"`
	} `json:"opening_time"`

	// ClosingTime is the closing time of the line
	ClosingTime struct {
		Hours   uint8 `json:"hours"`
		Minutes uint8 `json:"minutes"`
		Seconds uint8 `json:"seconds"`
	} `json:"closing_time"`

	// Routes contains the routes of the line
	Routes []Route `json:"routes"`

	// CommercialMode of the line
	CommercialMode CommercialMode `json:"commercial_mode"`

	// PhysicalModes of the line
	PhysicalModes []PhysicalMode `json:"physical_modes"`
}
