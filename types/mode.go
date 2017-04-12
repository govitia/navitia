package types

// A Mode represents a non-public transportation mode
type Mode string

// ModeXXX are known non-public transportation mode
const (
	ModeWalking string = "walking"
	ModeBike           = "bike"
	ModeCar            = "car"

	// Not used in Section
	ModeBikeShare = "bss"
)

// A CommercialMode codes for a commercial method of transportation
type CommercialMode struct {
	// A CommercialMode ID is in the form of "commercial_mode:something"
	ID            ID             `json:"id"`
	Name          string         `json:"name"`
	PhysicalModes []PhysicalMode `json:"physical_modes"`
}

// A PhysicalMode codes for a physical method of transportation
// For example, air travel, bus, metro and train.
type PhysicalMode struct {
	ID              ID               `json:"id"`
	Name            string           `json:"name"`
	CommercialModes []CommercialMode `json:"commercial_mode"`
}

// PhysicalModes is defined to help the user list all possible physical modes in ID form
var PhysicalModes = map[ID]string{
	PhysicalModeAir:               "Air",
	PhysicalModeBoat:              "Boat",
	PhysicalModeBus:               "Bus",
	PhysicalModeBusRapidTransit:   "Bus (Rapid Transit)",
	PhysicalModeCoach:             "Coach",
	PhysicalModeFerry:             "Ferry",
	PhysicalModeFunicular:         "Funicular",
	PhysicalModeLocalTrain:        "Local Train",
	PhysicalModeLongDistanceTrain: "Long-distance Train",
	PhysicalModeMetro:             "Metro",
	PhysicalModeRapidTransit:      "Rapid Transit",
	PhysicalModeShuttle:           "Shuttle",
	PhysicalModeTaxi:              "Taxi",
	PhysicalModeTrain:             "Train",
	PhysicalModeTramway:           "Tramway",
}

// PhysicalModeXXX are the possible physical modes in ID form
const (
	PhysicalModeAir               ID = "physical_mode:Air"
	PhysicalModeBoat                 = "physical_mode:Boat"
	PhysicalModeBus                  = "physical_mode:Bus"
	PhysicalModeBusRapidTransit      = "physical_mode:BusRapidTransit"
	PhysicalModeCoach                = "physical_mode:Coach"
	PhysicalModeFerry                = "physical_mode:Ferry"
	PhysicalModeFunicular            = "physical_mode:Funicular"
	PhysicalModeLocalTrain           = "physical_mode:LocalTrain"
	PhysicalModeLongDistanceTrain    = "physical_mode:LongDistanceTrain"
	PhysicalModeMetro                = "physical_mode:Metro"
	PhysicalModeRapidTransit         = "physical_mode:RapidTransit"
	PhysicalModeShuttle              = "physical_mode:Shuttle"
	PhysicalModeTaxi                 = "physical_mode:Taxi"
	PhysicalModeTrain                = "physical_mode:Train"
	PhysicalModeTramway              = "physical_mode:Tramway"
)
