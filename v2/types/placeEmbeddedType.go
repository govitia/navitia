package types

// PlaceEmbeddedType is an Enum used to identify what kind of objects /places and /places_nearby services are managing.
// It's also used inside different responses (journeys, ...).
type PlaceEmbeddedType string

const (
	PETAdministrativeRegion PlaceEmbeddedType = "administrative_region" // a city, a district, a neighborhood
	PETStopArea             PlaceEmbeddedType = "stop_area"             // a nameable zone, where there are some stop points
	PETStopPoint            PlaceEmbeddedType = "stop_point"
	PETAddress              PlaceEmbeddedType = "address"
	PETPoi                  PlaceEmbeddedType = "poi"
)
