package types

// Calendar is returned on vehicle journey message and indicates periodicity informations
// about transport schedules.
type Calendar struct {
	ActivePeriods []ActivePeriod `json:"active_periods"`
	WeekPattern   WeekPattern    `json:"week_pattern"`
	Exceptions    []Exception    `json:"exceptions"`
}
