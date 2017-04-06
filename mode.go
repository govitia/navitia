package types

// A Mode represents a non-public transportation mode
type Mode string

const (
	ModeWalking Mode = "walking"
	ModeBike         = "bike"
	ModeCar          = "car"

	// Not used in Section
	ModeBikeShare = "bss"
)

type CommercialModeID string

type CommercialMode struct {
	ID            CommercialModeID `json:"id"`
	Name          string           `json:"name"`
	PhysicalModes []PhysicalMode   `json:"physical_modes"`
}

type PhysicalModeID string

type PhysicalMode struct {
	ID              PhysicalModeID   `json:"id"`
	Name            string           `json:"name"`
	CommercialModes []CommercialMode `json:"commercial_mode"`
}

// PhysicalModes is defined to help a programmer list all possible physical modes
var PhysicalModes = map[string]PhysicalModeID{
	"Air":  PhysicalModeAir,
	"Boat": PhysicalModeBoat,
	"Bus":  PhysicalModeBus,
	"Bus (rapid transit)": PhysicalModeBusRapidTransit,
	"Coach":               PhysicalModeCoach,
	"Ferry":               PhysicalModeFerry,
	"Funicular":           PhysicalModeFunicular,
	"Local train":         PhysicalModeLocalTrain,
	"Long-distance train": PhysicalModeLongDistanceTrain,
	"Metro":               PhysicalModeMetro,
	"Rapid transit":       PhysicalModeRapidTransit,
	"Shuttle":             PhysicalModeShuttle,
	"Taxi":                PhysicalModeTaxi,
	"Train":               PhysicalModeTrain,
	"Tramway":             PhysicalModeTramway,
}

const (
	PhysicalModeAir               PhysicalModeID = "physical_mode:Air"
	PhysicalModeBoat                             = "physical_mode:Boat"
	PhysicalModeBus                              = "physical_mode:Bus"
	PhysicalModeBusRapidTransit                  = "physical_mode:BusRapidTransit"
	PhysicalModeCoach                            = "physical_mode:Coach"
	PhysicalModeFerry                            = "physical_mode:Ferry"
	PhysicalModeFunicular                        = "physical_mode:Funicular"
	PhysicalModeLocalTrain                       = "physical_mode:LocalTrain"
	PhysicalModeLongDistanceTrain                = "physical_mode:LongDistanceTrain"
	PhysicalModeMetro                            = "physical_mode:Metro"
	PhysicalModeRapidTransit                     = "physical_mode:RapidTransit"
	PhysicalModeShuttle                          = "physical_mode:Shuttle"
	PhysicalModeTaxi                             = "physical_mode:Taxi"
	PhysicalModeTrain                            = "physical_mode:Train"
	PhysicalModeTramway                          = "physical_mode:Tramway"
)
