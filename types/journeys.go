package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// A JourneyQualification qualifies a Journey, see const declaration.
type JourneyQualification string

// JourneyXXX qualify journeys
const (
	JourneyBest          JourneyQualification = "best"
	JourneyRapid         JourneyQualification = "rapid"
	JourneyComfort       JourneyQualification = "comfort"
	JourneyCar           JourneyQualification = "car"
	JourneyLessWalk      JourneyQualification = "less_fallback_walk"
	JourneyLessBike      JourneyQualification = "less_fallback_bike"
	JourneyLessBikeShare JourneyQualification = "less_fallback_bss"
	JourneyFastest       JourneyQualification = "fastest"
	JourneyNoPTWalk      JourneyQualification = "non_pt_walk"
	JourneyNoPTBike      JourneyQualification = "non_pt_bike"
	JourneyNoPTBikeShare JourneyQualification = "non_pt_bss"
)

// JourneyQualifications is a user-friendly slice of all journey qualification
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

	// Status from the whole journey taking into acount the most disturbing information retrieved on every object used
	Status Effect
}

// jsonJourney define the JSON implementation of Journey types
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonJourney struct {
	Duration  int64 `json:"duration"`
	Transfers *uint `json:"nb_transfers"`

	Departure string `json:"departure_date_time"`
	Requested string `json:"requested_date_time"`
	Arrival   string `json:"arrival_date_time"`

	Sections *[]Section `json:"sections"`

	From *Container `json:"from"`
	To   *Container `json:"to"`

	Type *JourneyQualification `json:"type"`

	Fare *Fare `json:"fare"`

	Status *Effect `json:"status"`
}

// CO2Emissions holds how much CO2 is emitted.
type CO2Emissions struct {
	Unit  string
	Value float64
}

// jsonCO2Emissions define the JSON implementation of CO2Emissions types
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonCO2Emissions struct {
	Unit  *string `json:"unit"`
	Value string  `json:"value"`
}

// TravelerType is a Traveler's type
// Defines speeds & accessibility values for different types of people
type TravelerType string

// The defined types of the api
const (
	// A standard Traveler
	TravelerStandard TravelerType = "standard"

	// A slow walker
	TravelerSlowWalker TravelerType = "slow_walker"

	// A fast walker
	TravelerFastWalker TravelerType = "fast_walker"

	// A Traveler with luggage
	TravelerWithLuggage TravelerType = "luggage"

	// A Traveler in a wheelchair
	TravelerInWheelchair TravelerType = "wheelchair"
)

// UnmarshalJSON implements json.Unmarshaller for a Journey.
// Behaviour:
//	- If "from" is empty, then don't populate the From field.
//	- Same for "to"
func (j *Journey) UnmarshalJSON(b []byte) error {
	data := &jsonJourney{
		Transfers: &j.Transfers,
		Sections:  &j.Sections,
		From:      &j.From,
		To:        &j.To,
		Type:      &j.Type,
		Fare:      &j.Fare,
		Status:    &j.Status,
	}

	// Now unmarshall the raw data into the analogous structure
	if err := json.Unmarshal(b, data); err != nil {
		return fmt.Errorf("error while unmarshalling Journey: %w", err)
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Journey", b}

	// As the given duration is in second, let's multiply it by one second to have the correct value
	j.Duration = time.Duration(data.Duration) * time.Second

	var err error
	// For departure, requested and arrival, we use parseDateTime
	j.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return gen.err(err, "Departure", "departure_date_time", data.Departure, "parseDateTime failed")
	}
	j.Requested, err = parseDateTime(data.Requested)
	if err != nil {
		return gen.err(err, "Requested", "requested_date_time", data.Requested, "parseDateTime failed")
	}
	j.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return gen.err(err, "Arrival", "arrival_date_time", data.Arrival, "parseDateTime failed")
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaller for CO2Emissions
func (c *CO2Emissions) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &jsonCO2Emissions{
		Unit: &c.Unit,
	}

	// Now unmarshall the raw data into the analogous structure
	if err := json.Unmarshal(b, data); err != nil {
		return fmt.Errorf("error while unmarshalling CO2Emissions: %w", err)
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"CO2Emissions", b}

	// Now parse the value
	f, err := strconv.ParseFloat(data.Value, 64)
	if err != nil {
		return gen.err(err, "Value", "value", data.Value, "error in strconv.ParseFloat")
	}
	c.Value = f

	return nil
}
