package types

// An ImpactedSection records the impact to a section
type ImpactedSection struct {
	// The start of the disruption, spatially
	From Container `json:"from"`
	// Until this point
	To Container `json:"to"`
}
