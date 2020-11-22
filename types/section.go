package types

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/geojson"
)

// A Section holds information about a specific section
type Section struct {
	Type       SectionType
	ID         ID
	Mode       string
	From       Container
	To         Container
	Departure  time.Time        // Departure time
	Arrival    time.Time        // Arrival time
	Duration   time.Duration    // Duration of travel
	Path       []PathSegment    // The path taken by this section
	Geo        *geom.LineString // The path in geojson format
	StopTimes  []StopTime       // List of the stop times of this section
	Display    Display          // Information to display
	Additional []PTMethod       // Additional informations, from what I can see this is always a PTMethod
}

// jsonSection define the JSON implementation of Section struct
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonSection struct {
	// Pointers to the corresponding real values
	Type       *SectionType   `json:"type"`
	ID         *ID            `json:"id"`
	From       *Container     `json:"from"`
	To         *Container     `json:"to"`
	Mode       *string        `json:"mode"`
	StopTimes  *[]StopTime    `json:"stop_date_times"`
	Display    *Display       `json:"display_informations"`
	Additional *[]PTMethod    `json:"additional_informations"`
	Path       *[]PathSegment `json:"path"`

	// Values to process
	Departure string            `json:"departure_date_time"`
	Arrival   string            `json:"arrival_date_time"`
	Duration  int64             `json:"duration"`
	Geo       *geojson.Geometry `json:"geojson"`
}

// A SectionType codifies the type of section that can be encountered
type SectionType string

// These are the types of sections that can be returned from the API
const (
	// Public transport section
	SectionPublicTransport SectionType = "public_transport"

	// Street section
	SectionStreetNetwork SectionType = "street_network"

	// Waiting section between transport
	SectionWaiting SectionType = "waiting"

	// This “stay in the vehicle” section occurs when the traveller has to stay in the vehicle when the bus change its routing.
	SectionStayIn SectionType = "stay_in"

	// Transfer section
	SectionTransfer SectionType = "transfer"

	// Teleportation section. Used when starting or arriving to a city or a stoparea (“potato shaped” objects) Useful to make navitia idempotent.
	// Warning: Be careful: no Path nor Geo items in this case
	SectionCrowFly SectionType = "crow_fly"

	// Vehicle may not drive along: traveler will have to call agency to confirm journey
	// Also sometimes called ODT
	SectionOnDemandTransport SectionType = "on_demand_transport"

	// Taking a bike from a bike sharing system (bss)
	SectionBikeShareRent SectionType = "bss_rent"

	// Putting back a bike from a bike sharing system (bss)
	SectionBikeSharePutBack SectionType = "bss_put_back"

	// Boarding on plane
	SectionBoarding SectionType = "boarding"

	// Landing off the plane
	SectionLanding SectionType = "landing"
)

// SectionTypes is the type of a section
var SectionTypes = map[SectionType]string{
	SectionPublicTransport:   "Public transport section",
	SectionStreetNetwork:     "Street section",
	SectionWaiting:           "Waiting section between transport",
	SectionStayIn:            "This “stay in the vehicle” section occurs when the traveller has to stay in the vehicle when the bus change its routing.",
	SectionTransfer:          "Transfer section",
	SectionCrowFly:           "Teleportation section. Used when starting or arriving to a city or a stoparea (“potato shaped” objects) Useful to make navitia idempotent",
	SectionOnDemandTransport: "Vehicle may not drive along: traveler will have to call agency to confirm journey",
	SectionBikeShareRent:     "Taking a bike from a bike sharing system (bss)",
	SectionBikeSharePutBack:  "Putting back a bike from a bike sharing system (bss)",
	SectionBoarding:          "Boarding on plane",
	SectionLanding:           "Landing off the plane",
}

// A StopTime stores info about a stop in a route: when the vehicle comes in, when it comes out, and what stop it is.
type StopTime struct {
	// The PTDateTime of the stop, this stores the info about the arrival & departure
	PTDateTime       PTDateTime
	StopPoint        StopPoint `json:"stop_point"` // The stop point in question
	DropOffAllowed   bool      `json:"drop_off_allowed"`
	UTCDepartureTime string    `json:"utc_departure_time"`
	Headsign         string    `json:"headsign"`
	UTCArrivalTime   string    `json:"utc_arrival_time"`
	PickupAllowed    bool      `json:"pickup_allowed"`
	DepartureTime    string    `json:"departure_time"`
}

// A PTMethod is a Public Transportation method: it can be regular, estimated times or ODT (on-demand transport)
type PTMethod string

// PTMethodXXX codes for known PTMethod
const (
	// PTMethodRegular: No on-demand transport. Line does not contain any estimated stop times, nor zonal stop point location. No need to call too.
	PTMethodRegular PTMethod = "regular"

	// PTMethodDateTimeEstimated: No on-demand transport. However, line has at least one estimated date time.
	PTMethodDateTimeEstimated PTMethod = "had_date_time_estimated"

	// PTMethodODTStopTime: Line does not contain any estimated stop times, nor zonal stop point location. But you will have to call to take it.
	PTMethodODTStopTime PTMethod = "odt_with_stop_time"

	// PTMethodODTStopPoint: Line can contain some estimated stop times, but no zonal stop point location. And you will have to call to take it.
	PTMethodODTStopPoint PTMethod = "odt_with_stop_point"

	// PTMethodODTZone: Line can contain some estimated stop times, and zonal stop point location. And you will have to call to take it. Well, not really a public transport line, more a cab…
	PTMethodODTZone PTMethod = "odt_with_zone"
)

/*
UnmarshalJSON implements json.Unmarshaller for a Section

Behaviour:
	- If "from" is empty, then don't populate the From field.
	- Same for "to"
*/
func (s *Section) UnmarshalJSON(b []byte) error {
	data := &jsonSection{
		Type:       &s.Type,
		ID:         &s.ID,
		From:       &s.From,
		To:         &s.To,
		Mode:       &s.Mode,
		Display:    &s.Display,
		Additional: &s.Additional,
		StopTimes:  &s.StopTimes,
		Path:       &s.Path,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return fmt.Errorf("error while unmarshalling Section: %w", err)
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"Section", b}

	// For departure and arrival, we use parseDateTime
	s.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return gen.err(err, "Departure", "departure_date_time", data.Departure, "parseDateTime failed")
	}
	s.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return gen.err(err, "Arrival", "arrival_date_time", data.Arrival, "parseDateTime failed")
	}

	// As the given duration is in second, let's multiply it by one second to have the correct value
	s.Duration = time.Duration(data.Duration) * time.Second

	// Now let's deal with the geom
	if data.Geo != nil {
		// Catch an error !
		if data.Geo.Coordinates == nil {
			return gen.err(nil, "Geo", "geojson", data.Geo, "Geo.Coordinates is nil, can't continue as that will cause a panic")
		}

		// Let's decode it
		geot, err := data.Geo.Decode()
		if err != nil {
			return gen.err(err, "Geo", "geojson", data.Geo, "Geo.Decode() failed")
		}
		// And let's assert the type
		geo, ok := geot.(*geom.LineString)
		if !ok {
			return gen.err(err, "Geo", "geojson", data.Geo, "Geo type assertion failed!")
		}
		// Now let's assign it
		s.Geo = geo
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaller for a PTDateTime
func (ptdt *PTDateTime) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	data := &struct {
		Departure string `json:"departure_date_time"`
		Arrival   string `json:"arrival_date_time"`
	}{}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return fmt.Errorf("error while unmarshalling PTDateTime: %w", err)
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"PTDateTime", b}

	// Now we use parseDateTime
	ptdt.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return gen.err(err, "Departure", "departure_date_time", data.Departure, "parseDateTime failed")
	}
	ptdt.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return gen.err(err, "Arrival", "arrival_date_time", data.Arrival, "parseDateTime failed")
	}

	return nil
}
