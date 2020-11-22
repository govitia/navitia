package types

// An ImpactedObject describes a PTObject impacted by a Disruption with some additional info.
type ImpactedObject struct {
	// The impacted public transport object
	Object Container `json:"pt_object"`

	// Only for line section impact, the impacted section
	ImpactedSection ImpactedSection `json:"impacted_section"`

	// Only for Trip delay, the list of delays, stop by stop
	ImpactedStops []ImpactedStop `json:"impacted_stops"`
}
