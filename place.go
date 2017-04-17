package types

import (
	"fmt"
	"github.com/pkg/errors"
)

// A Place isn't something directly used by the Navitia.io api.
//
// However, it allows the library user to use idiomatic go when working with the library.
// If you want a countainer, see PlaceContainer
//
// Place is satisfied by:
// 	- StopArea
// 	- POI
// 	- Address
// 	- StopPoint
// 	- AdministrativeRegion
type Place interface {
	// PlaceID returns the ID associated with the Place
	PlaceID() ID

	// PlaceName returns the name of the Place
	PlaceName() string

	// PlaceType returns the name of the type of the Place
	PlaceType() string

	// String() for string printing
	String() string
}

// PlaceContainer is the ugly countainer sent by Navitia to make us all cry :(
//
// However, this can be useful. May be removed from the public API in gonavitia v1.
type PlaceContainer struct {
	ID           ID     `json:"id"`
	Name         string `json:"name"`
	Quality      uint   `json:"quality,omitempty"`
	EmbeddedType string `json:"embedded_type"`

	// Four possibilitiess
	StopArea             *StopArea             `json:"stop_area,omitempty"`
	POI                  *POI                  `json:"POI,omitempty"`
	Address              *Address              `json:"address,omitempty"`
	StopPoint            *StopPoint            `json:"stop_point,omitempty"`
	AdministrativeRegion *AdministrativeRegion `json:"administrative_region,omitempty"`
}

// ErrInvalidPlaceContainer is returned after a check on a PlaceContainer
type ErrInvalidPlaceContainer struct {
	// If the PlaceContainer has a zero ID.
	NoID bool

	// If the PlaceContainer has a zero EmbeddedType.
	NoEmbeddedType bool

	// If the PlaceContainer has an unknown EmbeddedType
	UnknownEmbeddedType bool

	// If the PlaceContainer has a known EmbeddedType but the corresponding concrete value is nil.
	NilConcretePlaceValue bool

	// TODO:
	// If the PlaceContainer has more than one non-nil concrete place values.
	// MultipleNonNilConcretePlaceValues bool
}

// Error satisfies the error interface
func (err ErrInvalidPlaceContainer) Error() string {
	// Count the number of anomalies
	var anomalies uint

	msg := "Error: Invalid non-empty PlaceContainer (%d anomalies):"

	if err.NoID {
		msg += "\n\tNo ID specified"
		anomalies++
	}
	if err.NoEmbeddedType {
		msg += "\n\tEmpty EmbeddedType"
		anomalies++
	}
	if err.UnknownEmbeddedType {
		msg += "\n\tUnknown EmbeddedType"
		anomalies++
	}
	if err.NilConcretePlaceValue {
		msg += "\n\tKnown EmbeddedType but nil corresponding concrete value"
		anomalies++
	}

	return fmt.Sprintf(msg, anomalies)
}

// IsEmpty returns whether or not pc is empty
func (pc PlaceContainer) IsEmpty() bool {
	empty := PlaceContainer{}
	return pc == empty
}

var placeTypes = []string{
	embeddedStopArea,
	embeddedPOI,
	embeddedAddress,
	embeddedStopPoint,
	embeddedAddress,
}

const (
	embeddedStopArea  string = "stop_area"
	embeddedPOI              = "poi"
	embeddedAddress          = "address"
	embeddedStopPoint        = "stop_point"
	embeddedAdmin            = "administrative_region"
)

// Check checks the validity of the PlaceContainer. Returns an ErrInvalidPlaceContainer.
//
// An empty PlaceContainer is valid. But those cases aren't:
// 	- If the PlaceContainer has a zero ID.
// 	- If the PlaceContainer has a zero EmbeddedType.
// 	- If the PlaceContainer has an unknown EmbeddedType.
// 	- If the PlaceContainer has a known EmbeddedType but the corresponding concrete value is nil.
func (pc PlaceContainer) Check() error {
	if pc.IsEmpty() {
		return nil
	}
	err := ErrInvalidPlaceContainer{}

	// Check for zero ID
	err.NoID = (pc.ID == "")

	// Check if known EmbeddedType but the corresponding concrete value is nil.
	// Also check for empty and unknown embedded type
	absent := &err.NilConcretePlaceValue
	switch pc.EmbeddedType {
	case embeddedStopArea:
		*absent = (pc.StopArea == nil)
	case embeddedPOI:
		*absent = (pc.POI == nil)
	case embeddedAddress:
		*absent = (pc.Address == nil)
	case embeddedStopPoint:
		*absent = (pc.StopPoint == nil)
	case embeddedAdmin:
		*absent = (pc.AdministrativeRegion == nil)
	default:
		// Check for an empty embedded type
		err.NoEmbeddedType = (pc.EmbeddedType == "")
		// In case its not empty yet doesn't match with a known type
		err.UnknownEmbeddedType = !err.NoEmbeddedType
	}

	emptyErr := ErrInvalidPlaceContainer{}
	if err != emptyErr {
		return err
	}

	return nil
}

// Place returns the Place countained in the PlaceContainer
// If PlaceContainer is empty, Place returns an error.
// Check() is run on the PlaceContainer.
func (pc PlaceContainer) Place() (Place, error) {
	// If PlaceContainer is empty, return an error
	if pc.IsEmpty() {
		return nil, errors.Errorf("this place countainer is empty, can't extract a Place from it")
	}

	// Check validity
	err := pc.Check()
	if err != nil {
		return nil, err
	}

	// Check for each type
	switch pc.EmbeddedType {
	case embeddedStopArea:
		return pc.StopArea, nil
	case embeddedPOI:
		return pc.POI, nil
	case embeddedAddress:
		return pc.Address, nil
	case embeddedStopPoint:
		return pc.StopPoint, nil
	case embeddedAdmin:
		return pc.AdministrativeRegion, nil
	default:
		return nil, errors.Errorf("no known embedded type indicated (we have \"%s\"), can't return a place !", pc.EmbeddedType) // THIS IS VERY SERIOUS AS WE ALREADY CHECKED THE STRUCTURE
	}
}

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
	return embeddedStopArea
}

// String pretty-prints the StopArea.
// Satisifes Stringer and helps satisfy Place
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

// PlaceName returns the name of the POI.
// Helps satisfy Place
func (poi POI) PlaceName() string {
	return poi.Name
}

// PlaceType returns the type of place, in this case "poi".
// Helps satisfy Place
func (poi POI) PlaceType() string {
	return embeddedPOI
}

// String pretty-prints the POI.
// Satisifes Stringer and helps satisfy Place
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

// A POIType codes for the type of the point of interest
// TODO: A list of usual types ?
type POIType struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
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
	return embeddedAddress
}

// String pretty-prints the Address.
// Satisifes Stringer and helps satisfy Place
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
	return embeddedStopPoint
}

// String pretty-prints the StopPoint.
// Satisifes Stringer and helps satisfy Place
func (sp StopPoint) String() string {
	format := "%s (id: %s)"
	return fmt.Sprintf(format, sp.Name, sp.ID)
}

// An AdministrativeRegion represents an administrative region: a region under the control/responsibility of a specific organisation.
// It can be a city, a district, a neightborhood, etc.
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
	return embeddedAdmin
}

// String pretty-prints the AdministrativeRegion.
// Satisifes Stringer and helps satisfy Place
func (ar AdministrativeRegion) String() string {
	var label string
	if ar.Label == "" {
		label = ar.Name
	} else {
		label = ar.Label
	}

	format := "%s (id: %s)"
	return fmt.Sprintf(format, label, ar.ID)
}
