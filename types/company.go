package types

// A Company is a provider of transport
//
// Example: the RATP in Paris
//
// See http://doc.navitia.io/#public-transport-objects
type Company struct {
	// Identifier of the company
	ID string `json:"id"`

	// Name of the company
	Name string `json:"name"`
}
