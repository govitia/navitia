package types

// A Line codes for a public transit line.
// Warning: a Line isn't a route, it has no direction information, and can have several embranchments.
type Line struct {
	// ID is the navitia identifier of the line
	// For example: "line:RAT:M6"
	ID string `json:"id"`

	// Name is the name of the line
	// For example: "Nation - Charles de Gaule Etoile"
	Name string `json:"name"`

	// Code is the codename of the line.
	// For example: "6"
	Code string `json:"code"`

	// Color is the color given to the line
	// For example: "79BB92"
	Color Color `json:"color"`

	// OpeningTime is the opening time of the line in HHMMSS format
	OpeningTime string `json:"opening_time"`

	// ClosingTime is the closing time of the line in HHMMSS format
	ClosingTime string `json:"closing_time"`

	// Routes countains the routes of the line
	Routes []Route `json:"routes"`

	// CommercialMode of the line
	ComercialMode CommercialMode `json:"commercial_mode"`

	// PhysicalModes of the line
	PhysicalModes []PhysicalMode `json:"physical_mode"`
}

// Color is an RGB representation of a color
// TODO: Transition to color.NRGBA
type Color string
