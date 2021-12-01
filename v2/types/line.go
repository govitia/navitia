package types

// A Line codes for a public transit line.
// Warning: a Line isn't a route, it has no direction information, and can have several branches.
// See http://doc.navitia.io/#public-transport-objects.
type Line struct {
	ID             string          `json:"id"`              // Identifier of the line
	Name           string          `json:"name"`            // Name of the line
	Code           string          `json:"code"`            // Code name of the line
	Color          HexColor        `json:"color"`           // Color of the line
	OpeningTime    string          `json:"opening_time"`    // Opening time of the line
	ClosingTime    string          `json:"closing_time"`    // Closing time of the line
	Routes         []*Route        `json:"routes"`          // Routes of the line
	CommercialMode CommercialMode  `json:"commercial_mode"` // Commercial mode of the line
	PhysicalModes  []*PhysicalMode `json:"physical_modes"`  // Physical modes of the line
}
