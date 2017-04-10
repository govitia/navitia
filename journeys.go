package types

import (
	"golang.org/x/text/currency"
	"time"
)

// A JourneyQualification qualifies a Journey, see const declaration.
type JourneyQualification string

// JourneySomething qualify journeys
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

// JourneyQualifications is a user-friendly map of all journey qualification
var JourneyQualifications = map[string]JourneyQualification{
	"Best":                                   JourneyBest,
	"Rapid":                                  JourneyRapid,
	"Comfort":                                JourneyComfort,
	"Car":                                    JourneyCar,
	"Less walking":                           JourneyLessWalk,
	"Less biking":                            JourneyLessBike,
	"Less bike sharing":                      JourneyLessBikeShare,
	"Fastest":                                JourneyFastest,
	"No public transit, prefer walking":      JourneyNoPTWalk,
	"No public transit, prefer biking":       JourneyNoPTBike,
	"No public transit, prefer bike-sharing": JourneyNoPTBikeShare,
}

// DateTimeFormat is the format used by the Navitia Api for use with time pkg.
// Few external use-cases but still there are some
const DateTimeFormat string = "20060102T150405" // YYYYMMDDThhmmss

// A Journey holds information about a possible journey
type Journey struct {
	Duration  time.Duration
	Transfers uint

	Departure time.Time
	Requested time.Time
	Arrival   time.Time

	Sections []Section

	From Place
	To   Place

	Type JourneyQualification

	Fare Fare

	//Status from the whole journey taking into acount the most disturbing information retrieved on every object used
	Status JourneyStatus
}

// JourneyStatus codes for known journey status information
// For example, reduced service, detours or moved stops.
type JourneyStatus string

// JourneyStatusXXX are known JourneyStatuse
const (
	JourneyStatusNoService         JourneyStatus = "NO_SERVICE"
	JourneyStatusReducedService                  = "REDUCED_SERVICE"
	JourneyStatusSignificantDelay                = "SIGNIFICANT_DELAY"
	JourneyStatusDetour                          = "DETOUR"
	JourneyStatusAdditionalService               = "ADDITIONAL_SERVICE"
	JourneyStatusModifiedService                 = "MODIFIED_SERVICE"
	JourneyStatusOtherEffect                     = "OTHER_EFFECT"
	JourneyStatusUnknownEffect                   = "UNKNOWN_EFFECT"
	JourneyStatusStopMoved                       = "STOP_MOVED"
)

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
