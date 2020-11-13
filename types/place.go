package types

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

	Timezone string `json:"timezone"`
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

// A StopPoint codes for a stop point in a line: a location where vehicles can pickup or drop off passengers.
type StopPoint struct {
	ID ID `json:"id"`

	// Name of the stop point
	Name string `json:"name"`

	Label string `json:"label"`

	// Coordinates of the stop point
	Coord Coordinates `json:"coord"`

	// Administrative regions of the stop point
	Admins []Admin `json:"administrative_regions"`

	// List of equipments of the stop point
	Equipments []Equipment `json:"equipment"`

	// Stop Area countaining the stop point
	StopArea *StopArea `json:"stop_area"`

	CommercialModes []CommercialMode `json:"commercial_modes"`

	Links []Link `json:"links"`

	PhysicalModes []PhysicalMode `json:"physical_modes"`

	FareZone FareZone `json:"fare_zone"`
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

	Insee string `json:"insee"`
}
