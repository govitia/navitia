package types

// Network represents a specific network
// They are fed by the agencies
type Network struct {
	// ID is the identifier of the network
	ID string `json:"id"`

	// Name is the name of the network
	Name string `json:"name"`
}
