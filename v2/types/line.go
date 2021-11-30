package types

// A Line codes for a public transit line.
// Warning: a Line isn't a route, it has no direction information, and can have several branches.
// See http://doc.navitia.io/#public-transport-objects.
type Line struct {
	ID    string   `json:"id"`    // Identifier of the line
	Name  string   `json:"name"`  // Name of the line
	Code  string   `json:"code"`  // Code name of the line
	Color HexColor `json:"color"` // Color of the line
}
