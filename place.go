package types

import (
	"fmt"
)

// A Place isn't something directly used by the Navitia.io api.
//
// However, it allows the library user to use idiomatic go when working with the library.
// If you want a countainer, see Container
//
// Place is held by these types:
// 	- StopArea
// 	- POI
// 	- Address
// 	- StopPoint
// 	- Admin
type Place interface{}

// A StopArea represents a stop area: a nameable zone, where there are some stop points.
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
	Admins []Admin `json:"administrative_regions"`

	// Stop points countained in this stop area
	StopPoints []StopPoint `json:"stop_points"`
}

// String pretty-prints the StopArea.
//
// Satisfies Stringer and helps satisfy Place
func (sa StopArea) String() string {
	var label string
	if sa.Label == "" {
		label = sa.Name
	} else {
		label = sa.Label
	}

	format := "%s (id: %s)"
	return fmt.Sprintf(format, label, sa.ID)
}

// A POIType codes for the type of the point of interest
type POIType struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
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

// String pretty-prints the POI.
//
// Satisfies Stringer and helps satisfy Place.
func (poi POI) String() string {
	var label string
	if poi.Label == "" {
		label = poi.Name
	} else {
		label = poi.Label
	}

	format := "%s (id: %s)"
	return fmt.Sprintf(format, label, poi.ID)
}

// An Address codes for a real-world address: a point located in a street.
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
	Admins []Admin `json:"administrative_regions"`
}

// String pretty-prints the Address.
//
// Satisfies Stringer and helps satisfy Place.
func (add Address) String() string {
	var label string
	if add.Label == "" {
		label = add.Name
	} else {
		label = add.Label
	}

	format := "%s (id: %s)"
	return fmt.Sprintf(format, label, add.ID)
}

// A StopPoint codes for a stop point in a line: a location where vehicles can pickup or drop off passengers.
type StopPoint struct {
	ID ID `json:"id"`

	// Name of the stop point
	Name string `json:"name"`

	// Coordinates of the stop point
	Coord Coordinates `json:"coord"`

	// Administrative regions of the stop point
	Admins []Admin `json:"administrative_regions"`

	// List of equipments of the stop point
	Equipments []Equipment `json:"equipment"`

	// Stop Area countaining the stop point
	StopArea *StopArea `json:"stop_area"`
}

// String pretty-prints the StopPoint.
//
// Satisfies Stringer and helps satisfy Place.
func (sp StopPoint) String() string {
	format := "%s (id: %s)"
	return fmt.Sprintf(format, sp.Name, sp.ID)
}

// An Admin represents an administrative region: a region under the control/responsibility of a specific organisation.
// It can be a city, a district, a neightborhood, etc.
type Admin struct {
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
	ZipCode string `json:"zip_code"`
}

// String pretty-prints the Admin.
//
// Satisfies Stringer and helps satisfy Place.
func (ar Admin) String() string {
	var label string
	if ar.Label == "" {
		label = ar.Name
	} else {
		label = ar.Label
	}

	format := "%s (id: %s)"
	return fmt.Sprintf(format, label, ar.ID)
}
