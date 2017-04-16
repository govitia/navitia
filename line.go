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
	ID ID

	// Name is the name of the line
	// For example: "Nation - Charles de Gaule Etoile"
	Name string

	// Code is the codename of the line.
	// For example: "6"
	Code string

	// Color is the color given to the line
	// For example: "79BB92" in Hex
	Color color.Color

	// OpeningTime is the opening time of the line in HHMMSS format
	OpeningTime string

	// ClosingTime is the closing time of the line in HHMMSS format
	ClosingTime string

	// Routes countains the routes of the line
	Routes []Route

	// CommercialMode of the line
	CommercialMode CommercialMode

	// PhysicalModes of the line
	PhysicalModes []PhysicalMode
}
