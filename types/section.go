package types

import "time"

// A SectionType codifies the type of section that can be encountered
type SectionType string

// These are the types of sections that can be returned from the API
const (
	// Public transport section
	SectionPublicTransport SectionType = "public_transport"

	// Street section
	SectionStreetNetwork = "street_network"

	// Waiting section between transport
	SectionWaiting = "waiting"

	// This “stay in the vehicle” section occurs when the traveller has to stay in the vehicle when the bus change its routing.
	SectionStayIn = "stay_in"

	// Transfer section
	SectionTransfer = "transfer"

	// Teleportation section. Used when starting or arriving to a city or a stoparea (“potato shaped” objects) Useful to make navitia idempotent.
	// Warning: Be careful: no Path nor Geo items in this case
	SectionCrowFly = "crow_fly"

	// Vehicle may not drive along: traveler will have to call agency to confirm journey
	// Also sometimes called ODT
	SectionOnDemandTransport = "on_demand_transport"

	// Taking a bike from a bike sharing system (bss)
	SectionBikeShareRent = "bss_rent"

	// Putting back a bike from a bike sharing system (bss)
	SectionBikeSharePutBack = "bss_put_back"

	// Boarding on plane
	SectionBoarding = "boarding"

	// Landing off the plane
	SectionLanding = "landing"
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

// A Section holds information about a specific section
type Section struct {
	Type SectionType
	ID   ID
	Mode Mode

	// Arrival time & departure time
	Departure time.Time
	Arrival   time.Time

	// Duration of travel
	Duration time.Duration

	// From & To
	From Place
	To   Place

	// The path taken by this section
	Path []PathSegment

	// List of the stop times of this section
	StopTimes []StopTime

	// Information to display
	Display DisplayInformations

	// Additional informations, from what I can see this is always a PTMethod
	Additional []PTMethod
}

// A StopTime stores info about a stop in a route: when the vehicle comes in, when it comes out, and what stop it is.
type StopTime struct {
	// The PTDateTime of the stop, this stores the info about the arrival & departure
	PTDateTime PTDateTime

	// The stop point in question
	StopPoint StopPoint
}

// A PTDateTime (pt stands for “public transport”) is a complex date time object to manage the difference between stop and leaving times at a stop.
type PTDateTime struct {
	// Date/Time of departure
	Departure time.Time

	// Date/Time of arrival
	Arrival time.Time
}

// A PTMethod is a Public Transportation method: it can be regular, estimated times or ODT (on-demand transport)
type PTMethod string

// PTMethodXXX codes for known PTMethod
const (
	// PTMethodRegular: No on-demand transport. Line does not contain any estimated stop times, nor zonal stop point location. No need to call too.
	PTMethodRegular PTMethod = "regular"

	// PTMethodDateTimeEstimated: No on-demand transport. However, line has at least one estimated date time.
	PTMethodDateTimeEstimated = "had_date_time_estimated"

	// PTMethodODTStopTime: Line does not contain any estimated stop times, nor zonal stop point location. But you will have to call to take it.
	PTMethodODTStopTime = "odt_with_stop_time"

	// PTMethodODTStopPoint: Line can contain some estimated stop times, but no zonal stop point location. And you will have to call to take it.
	PTMethodODTStopPoint = "odt_with_stop_point"

	// PTMethodODTZone: Line can contain some estimated stop times, and zonal stop point location. And you will have to call to take it. Well, not really a public transport line, more a cab…
	PTMethodODTZone = "odt_with_zone"
)

// A DisplayInformations hold informations useful to display
type DisplayInformations struct {
	// The headsign associated with the object
	Headsign string `json:"headsign"`

	// The name of the belonging network
	Network string `json:"network"`

	// A direction to take
	Direction string `json:"direction"`

	// The commercial mode in ID Form
	CommercialMode ID `json:"commercial_mode"`

	// The physical mode in ID Form
	PhysicalMode ID `json:"physical_mode"`

	// The label of the object
	Label string `json:"label"`

	// Hexadecimal color of the line
	Color string `json:"color"`

	// The text color for this section
	TextColor Color `json:"text_color"`

	// The code of the line
	Code string `json:"code"`

	// Description
	Description string `json:"description"`

	// Equipments on this object
	Equipments []Equipment `json:"equipments"`
}
