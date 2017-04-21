package types

import (
	"time"

	"golang.org/x/text/currency"
)

// A JourneyQualification qualifies a Journey, see const declaration.
type JourneyQualification string

// JourneyXXX qualify journeys
const (
	JourneyBest          JourneyQualification = "best"
	JourneyRapid                              = "rapid"
	JourneyComfort                            = "comfort"
	JourneyCar                                = "car"
	JourneyLessWalk                           = "less_fallback_walk"
	JourneyLessBike                           = "less_fallback_bike"
	JourneyLessBikeShare                      = "less_fallback_bss"
	JourneyFastest                            = "fastest"
	JourneyNoPTWalk                           = "non_pt_walk"
	JourneyNoPTBike                           = "non_pt_bike"
	JourneyNoPTBikeShare                      = "non_pt_bss"
)

// JourneyQualifications is a user-friendly slice of all journey qualification
//
// As it might be used in requests, this is exported
var JourneyQualifications = []JourneyQualification{
	JourneyBest,
	JourneyRapid,
	JourneyComfort,
	JourneyCar,
	JourneyLessWalk,
	JourneyLessBike,
	JourneyLessBikeShare,
	JourneyFastest,
	JourneyNoPTWalk,
	JourneyNoPTBike,
	JourneyNoPTBikeShare,
}

// A Journey holds information about a possible journey
type Journey struct {
	Duration  time.Duration
	Transfers uint

	Departure time.Time
	Requested time.Time
	Arrival   time.Time

	CO2Emissions CO2Emissions

	Sections []Section

	From Container
	To   Container

	Type JourneyQualification

	Fare Fare

	//Status from the whole journey taking into acount the most disturbing information retrieved on every object used
	Status Effect
}

// CO2Emissions contains the
type CO2Emissions struct {
	Unit  string
	Value float64
}

// Fare is the fare of some thing
type Fare struct {
	Total currency.Amount
	Found bool
}

// TravelerType is a Traveler's type
// Defines speeds & accessibility values for different types of people
type TravelerType string

// The defined types of the api
const (
	// A standard Traveler
	TravelerStandard TravelerType = "standard"

	// A slow walker
	TravelerSlowWalker = "slow_walker"

	// A fast walker
	TravelerFastWalker = "fast_walker"

	// A Traveler with luggage
	TravelerWithLuggage = "luggage"

	// A Traveler in a wheelchair
	TravelerInWheelchair = "wheelchair"
)
