package types

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

// sectionTypes, as it isn't useful to the library user, stays unexported
var sectionTypes = map[SectionType]string{
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

type SectionID string

// A Section holds information about a specific section
type Section struct {
	Type SectionType `json:"type"`
	ID   SectionID   `json:"id"`
	Mode Mode        `json:"mode"`

	From Place `json:"from"`
	To   Place `json:"to"`

	// Information to display
	Display DisplayInformations `json:"display_informations"`

	// Additional informations, from what I can see this is always a PTMethod
	Additional PTMethod `json:"additional_informations"`
}

// A PTMethod is a Public Transportation method: it can be regular, estimated times or ODT (on-demand transport)
type PTMethod string

const (
	// No on-demand transport. Line does not contain any estimated stop times, nor zonal stop point location. No need to call too.
	PTMethodRegular PTMethod = "regular"
	// No on-demand transport. However, line has at least one estimated date time.
	PTMethodDateTimeEstimated = "had_date_time_estimated"
	// Line does not contain any estimated stop times, nor zonal stop point location. But you will have to call to take it.
	PTMethodODTStopTime = "odt_with_stop_time"
	// Line can contain some estimated stop times, but no zonal stop point location. And you will have to call to take it.
	PTMethodODTStopPoint = "odt_with_stop_point"
	// Line can contain some estimated stop times, and zonal stop point location. And you will have to call to take it. Well, not really a public transport line, more a cab…
	PTMethodODTZone = "odt_with_zone"
)

// A DisplayInformations hold informations useful to display
type DisplayInformations struct {
	// The name of the belonging network
	Network string `json:"network"`

	// A direction to take
	Direction string `json:"direction"`

	// The commercial mode
	CommercialMode CommercialModeID `json:"commercial_mode"`

	// The physical mode
	PhysicalMode PhysicalModeID `json:"physical_mode"`

	// The label of the object
	Label string `json:"label"`

	// Hexadecimal color of the line
	Color string `json:"color"`

	// The code of the line
	Code string `json:"code"`

	// Description
	Description string `json:"description"`

	// Equipments on this object
	Equipments []Equipment `json:"equipments"`
}
