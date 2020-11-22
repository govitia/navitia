package types

// A Company is a provider of transport
// Example: the RATP in Paris
// See http://doc.navitia.io/#public-transport-objects
type Company struct {
	ID   string `json:"id"`   // Identifier of the company
	Name string `json:"name"` // Name of the company
}
