package types

// A Company is a provider of transport
// Example: the RATP in Paris
type Company struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
