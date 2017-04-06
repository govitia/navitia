package types

// A PlaceID codes for a place
type PlaceID string

// A Place is satisfied by:
// - StopArea
// - POI
// - Address
// - StopPoint
// - AdministrativeRegion
type Place interface {
	PlaceID() PlaceID
	PlaceName() string
	PlaceType() string
}

// A StopArea
type StopArea struct {
	ID   PlaceID `json:"id"`
	Name string  `json:"name"`

	// Label of the stop area.
	// The name is directly taken from the data whereas the label is something computed by navitia for better traveler information.
	// If you don't know what to display, display the label
	Label string `json:"label"`

	// Coordinates of the stop area
	Coord Coordinates `json:"coord"`

	// Administrative regions of the stop area in which is placed the stop area
	AdministrativeRegions []AdministrativeRegion `json:"administrative_regions"`

	// Stop points countained in this stop area
	StopPoints []StopPoint `json:"stop_points"`
}

func (sa StopArea) PlaceID() PlaceID {
	return sa.ID
}

func (sa StopArea) PlaceName() string {
	return sa.Name
}

func (sa StopArea) PlaceType() string {
	return "stop_area"
}

type POI struct {
	ID   PlaceID `json:"id"`
	Name string  `json:"name"`

	// The name is directly taken from the data whereas the label is something computed by navitia for better traveler information.
	// If you don't know what to display, display the label
	Label string `json:"label"`

	// The type of the POI
	Type POIType `json:"poi_type"`
}

func (poi POI) PlaceID() PlaceID {
	return poi.ID
}

func (poi POI) PlaceName() string {
	return poi.Name
}

func (poi POI) PlaceType() string {
	return "poi"
}

// A POIType codes for the type of the point of interest
// TODO: A list of usual types ?
type POIType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// An Address codes for a real-world address
type Address struct {
	ID   PlaceID `json:"id"`
	Name string  `json:"name"`

	// Label of the address
	// The name is directly taken from the data whereas the label is something computed by navitia for better traveler information.
	// If you don't know what to display, display the label
	Label string `json:"label"`

	// Coordinates of the address
	Coord Coordinates `json:"coord"`

	// House number of the address
	HouseNumber uint `json:"house_number"`

	// Administrative regions of the stop area in which is placed the stop area
	AdministrativeRegions []AdministrativeRegion `json:"administrative_regions"`
}

// A StopPoint codes for a stop point in a line
type StopPoint struct {
	ID PlaceID `json:"id"`

	// Name of the stop point
	Name string `json:"name"`

	// Coordinates of the stop point
	Coord Coordinates `json:"coord"`

	// Administrative regions of the stop point
	AdministrativeRegions []AdministrativeRegion `json:"administrative_regions"`

	// List of equipments of the stop point
	Equipments []Equipment `json:"equipment"`

	// Stop Area countaining the stop point
	StopArea *StopArea `json:"stop_area"`
}

type AdministrativeRegion struct {
	ID   PlaceID `json:"id"`
	Name string  `json:"name"`

	// Label of the address
	// The name is directly taken from the data whereas the label is something computed by navitia for better traveler information.
	// If you don't know what to display, display the label
	Label string `json:"label"`

	// Coordinates of the administrative region
	Coord Coordinates `json:"coord"`

	// Level of the administrative region
	Level int `json:"level"`

	// Zip code of the administrative region
	ZipCode string
}
