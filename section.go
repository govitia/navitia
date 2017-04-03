package gonavitia

type SectionType string

const (
	// Public transport section
	SectionPublicTransport SectionType = "public_transport"

	// Street section
	SectionStreetNetwork = "street_network"

	// Waiting section between transport
	SectionWaiting = "waiting"

	// This “stay in the vehicle” section occurs when the traveller has to stay in the vehicle when the bus change its routing. Here is an exemple for a journey from A to B: (lollipop line)
	SectionStayIn = "stay_in"

	// Transfer section
	SectionTransfer = "transfer"

	// Teleportation section. Used when starting or arriving to a city or a stoparea (“potato shaped” objects) Useful to make navitia idempotent. Be careful: no “path” nor “geojson” items in this case
	SectionCrowFly = "crow_fly"

	// Vehicle may not drive along: traveler will have to call agency to confirm journey
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

type SectionID string

// A Section holds information about a specific section
type Section struct {
	Type SectionType `json:"type"`
	ID   SectionID   `json:"id"`
	Mode Mode        `json:"mode"`

	From Place `json:"from"`
	To   Place `json:"to"`

	// Information to display
	Display DisplayInformations
}

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
