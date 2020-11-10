package types

// ModeXXX are known non-public transportation mode
const (
	ModeWalking = "walking"
	ModeBike    = "bike"
	ModeCar     = "car"

	// Not used in Section
	ModeBikeShare = "bss"
)

// A CommercialMode codes for a commercial method of transportation.
//
// Note that in contrast with physical modes, commercial modes aren't normalised, if you want to query with them, it is best to use a PhysicalMode.
//
// See http://doc.navitia.io/#public-transport-objects
type CommercialMode struct {
	// A CommercialMode ID is in the form of "commercial_mode:something"
	ID ID `json:"id"`

	// Name of the commercial mode
	Name string `json:"name"`

	// Physical modes of this commercial modes
	// Example: []PhysicalMode{PhysicalMode{ID: "physical_mode:Tramway", Name: "Tramway"}}
	PhysicalModes []PhysicalMode `json:"physical_modes"`
}

// A PhysicalMode codes for a physical method of transportation
// For example, air travel, bus, metro and train.
//
// As well, note that physical modes are normalised and fastened, see the list in PhysicalModes
//
// See http://doc.navitia.io/#public-transport-objects
type PhysicalMode struct {
	// Identifier of the physical mode
	// For example: "physical_mode:Tramway"
	ID ID `json:"id"`

	// Name of the physical mode
	// For example: "Tramway"
	Name string `json:"name"`

	// Commercial modes of this physical mode
	CommercialModes []CommercialMode `json:"commercial_mode"`
}

// PhysicalModeXXX are the possible physical modes in ID form
const (
	PhysicalModeAir               ID = "physical_mode:Air"
	PhysicalModeBoat              ID = "physical_mode:Boat"
	PhysicalModeBus               ID = "physical_mode:Bus"
	PhysicalModeBusRapidTransit   ID = "physical_mode:BusRapidTransit"
	PhysicalModeCoach             ID = "physical_mode:Coach"
	PhysicalModeFerry             ID = "physical_mode:Ferry"
	PhysicalModeFunicular         ID = "physical_mode:Funicular"
	PhysicalModeLocalTrain        ID = "physical_mode:LocalTrain"
	PhysicalModeLongDistanceTrain ID = "physical_mode:LongDistanceTrain"
	PhysicalModeMetro             ID = "physical_mode:Metro"
	PhysicalModeRapidTransit      ID = "physical_mode:RapidTransit"
	PhysicalModeShuttle           ID = "physical_mode:Shuttle"
	PhysicalModeTaxi              ID = "physical_mode:Taxi"
	PhysicalModeTrain             ID = "physical_mode:Train"
	PhysicalModeTramway           ID = "physical_mode:Tramway"
)
