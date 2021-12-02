package types

type EquipmentAvailability struct {
	Status string `json:"status"`
	Cause  *Label `json:"cause,omitempty"`  // If status is unavailable, gives you the cause in a label
	Effect *Label `json:"effect,omitempty"` // If status is unavailable, gives you the effect in a label
	// If status is unavailable, gives the affected period (with a begin & end datetime attributes)
	Periods []*Period `json:"periods"`
}
