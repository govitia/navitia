package types

// Trip corresponds to a scheduled vehicle circulation (and all its linked real-time and disrupted routes).
// Example: a train, routing a Paris to Lyon itinerary every day at 06h29, is the "Trip" named "6641".
//
// It encapsulates many instances of vehicle_journey.
type Trip struct {
	ID   string `json:"id"`   // The id of the trip
	Name string `json:"name"` // The name of the trip
}
