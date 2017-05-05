package types

// A Connection is either a Departure or an Arrival
type Connection struct {
	Display   Display
	StopPoint StopPoint
	Route     Route
	//StopDateTime [TO BE IMPLEMENTED]
}
