package types

// A trip corresponds to a scheduled vehicle circulation (and all its linked real-time and disrupted routes).
//
// An example : a train, routing a Paris to Lyon itinerary every day at 06h29, is the “Trip” named “6641”.
type Trip struct {
	ID   ID
	Name string
}
