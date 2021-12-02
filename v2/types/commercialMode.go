package types

// CommercialMode are close from physical modes,
// but not normalized and can refer to a brand, something that can
// be specific to a network, and known to the traveler.
// Examples: RER in Paris, Busway in Nantes, and also of course Bus, Métro, etc.
//
// Integrators should mainly use that value for text output to the traveler.
type CommercialMode struct {
	ID            string          `json:"id"`             // Identifier of the commercial mode
	Name          string          `json:"name"`           // Name of the commercial mode
	PhysicalModes []*PhysicalMode `json:"physical_modes"` // Physical modes of this commercial mode
}
