package types

// A Place isn't something directly used by the Navitia.io api
// However, it allows the library user to use idiomatic go when working with the library
// If you want a countainer, see PlaceCountainer
// Place is satisfied by:
// - StopArea
// - POI
// - Address
// - StopPoint
// - AdministrativeRegion
type Place interface {
	// PlaceID returns the ID associated with the Place
	PlaceID() ID

	// PlaceName returns the name of the Place
	PlaceName() string

	// PlaceType returns the name of the type of the Place
	PlaceType() string
}

// A StopArea represents a stop area: a place where a public transportation method may stop for a traveller.
type StopArea struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`

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

// PlaceID returns the ID associated with the StopArea
// Helps satisfy Place
func (sa StopArea) PlaceID() ID {
	return sa.ID
}

// PlaceName returns the name of the StopArea
// Helps satisfy Place
func (sa StopArea) PlaceName() string {
	return sa.Name
}

// PlaceType returns the type of place, in this case "stop_area"
// Helps satisfy Place
func (sa StopArea) PlaceType() string {
	return "stop_area"
}

// A POI is a Point Of Interest. A loosely-defined place.
type POI struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`

	// The name is directly taken from the data whereas the label is something computed by navitia for better traveler information.
	// If you don't know what to display, display the label
	Label string `json:"label"`

	// The type of the POI
	Type POIType `json:"poi_type"`
}

// PlaceID returns the ID associated with the POI.
// Helps satisfy Place
func (poi POI) PlaceID() ID {
	return poi.ID
}

// PlaceName returns the name of the POI
// Helps satisfy Place
func (poi POI) PlaceName() string {
	return poi.Name
}

// PlaceType returns the type of place, in this case "poi"
// Helps satisfy Place
func (poi POI) PlaceType() string {
	return "poi"
}

// A POIType codes for the type of the point of interest
// TODO: A list of usual types ?
type POIType struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

// An Address codes for a real-world address
type Address struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`

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

// PlaceID returns the ID associated with the Address
// Helps satisfy Place
func (add Address) PlaceID() ID {
	return add.ID
}

// PlaceName returns the name of the Address
// Helps satisfy Place
func (add Address) PlaceName() string {
	return add.Name
}

// PlaceType returns the type of place, in this case "address"
// Helps satisfy Place
func (add Address) PlaceType() string {
	return "address"
}

// A StopPoint codes for a stop point in a line
type StopPoint struct {
	ID ID `json:"id"`

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

// PlaceID returns the ID associated with the Stop Point
// Helps satisfy Place
func (sp StopPoint) PlaceID() ID {
	return sp.ID
}

// PlaceName returns the name of the Stop Point
// Helps satisfy Place
func (sp StopPoint) PlaceName() string {
	return sp.Name
}

// PlaceType returns the type of place, in this case "stop_point"
// Helps satisfy Place
func (sp StopPoint) PlaceType() string {
	return "stop_point"
}

// An AdministrativeRegion represents an administrative region: a region under the control/responsibility of a specific organisation.
type AdministrativeRegion struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`

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

// PlaceID returns the ID associated with the AdministrativeRegion
// Helps satisfy Place
func (ar AdministrativeRegion) PlaceID() ID {
	return ar.ID
}

// PlaceName returns the name of the AdministrativeRegion
// Helps satisfy Place
func (ar AdministrativeRegion) PlaceName() string {
	return ar.Name
}

// PlaceType returns the type of place, in this case "administrative_region"
// Helps satisfy Place
func (ar AdministrativeRegion) PlaceType() string {
	return "administrative_region"
}
