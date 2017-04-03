package gonavitia

import "time"

type JourneyResults struct {
	Journeys []Journey
}

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

// A Journey holds information about a possible journey
type Journey struct {
	Duration  time.Duration
	Transfers uint

	Departure time.Time
	Requested time.Time
	Arrival   time.Time

	Sections []Section

	From Place `json:"from"`
	To   Place `json:"to"`

	Type JourneyQualification `json:"type"`

	Fare Fare `json:"fare"`
}

// Fare is the fare of some thing
type Fare struct {
	Total Cost `json:"total"`
	Found bool `json:"found"`
}

// Cost is the cost of something
// I know value should NOT be float, but that's what the api gives us
type Cost struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"`
}
