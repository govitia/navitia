package types

import "image/color"

// A DisplayInformations hold informations useful to display
// Used by Section ["section"], RouteSchedule ["route_schedule"], StopSchedule ["stop_schedule"], Departure ["departure"], Arrival ["arrival"].
type DisplayInformations struct {
	// The headsign associated with the object
	Headsign string

	// The name of the belonging network
	Network string

	// A direction to take
	Direction string

	// The commercial mode in ID Form
	CommercialMode ID

	// The physical mode in ID Form
	PhysicalMode ID

	// The label of the object
	Label string

	// Hexadecimal color of the line
	Color color.Color

	// The text color for this section
	TextColor color.Color

	// The code of the line
	Code string

	// Description
	Description string

	// Equipments on this object
	Equipments []Equipment
}
