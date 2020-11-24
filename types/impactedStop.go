package types

// An ImpactedStop records the impact to a stop
type ImpactedStop struct {
	// The impacted stop point of the trip
	Point StopPoint `json:"stop_point"`

	// New departure hour (format HHMMSS) of the trip on this stop point
	NewDeparture string

	// New arrival hour (format HHMMSS) of the trip on this stop point
	NewArrival string

	// Base departure hour (format HHMMSS) of the trip on this stop point
	BaseDeparture string

	// Base arrival hour (format HHMMSS) of the trip on this stop point
	BaseArrival string

	// Cause of the modification
	Cause string `json:"cause"`

	// Effect on that StopPoint
	// Can be "added", "deleted", "delayed"
	Effect string

	AmendedArrivalTime string `json:"amended_arrival_time"`

	StopTimeEffect string `json:"stop_time_effect"`

	DepartureStatus string `json:"departure_status"`

	IsDetour bool `json:"is_detour"`

	AmendedDepartureTime string `json:"amended_departure_time"`

	BaseArrivalTime string `json:"base_arrival_time"`

	BaseDepartureTime string `json:"base_departure_time"`

	ArrivalStatus string `json:"arrival_status"`
}
