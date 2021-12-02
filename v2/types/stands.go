package types

type StandsStatus string

const (
	SSUnavailable StandsStatus = "unavailable" // Navitia is not able to obtain information about the station
	SSOpen        StandsStatus = "open"        // The station is open
	SSClosed      StandsStatus = "closed"      // The station is closed
)

// Stands id A description of the number of stands/places and vehicles available at a bike sharing station.
type Stands struct {
	AvailablePlaces int `json:"available_places"` // Number of places where one can park
	AvailableBikes  int `json:"available_bikes"`  // Number of bikes available
	// Total number of stands (occupied or not, with or without special equipment)
	TotalStands int          `json:"total_stands"`
	Status      StandsStatus `json:"status"` // Information about the station itself:
}
