package types

// A PTObject is a Public Transport object: StopArea, Trip, Line, Route, Network, etc.
type PTObject interface{}

// A Trip corresponds to a scheduled vehicle circulation (and all its linked real-time and disrupted routes).
//
// An example : a train, routing a Paris to Lyon itinerary every day at 06h29, is the “Trip” named “6641”.
type Trip struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}
