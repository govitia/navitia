package types

// Network represents a specific network.
// They are fed by the agencies in GTFS format.
// See http://doc.navitia.io/#public-transport-objects.
type Network struct {
	ID   string `json:"id"`   // ID is the identifier of the network
	Name string `json:"name"` // Name is the name of the network
}
