package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// Effect codes for known journey status information
// For example, reduced service, detours or moved stops.
//
// See https://developers.google.com/transit/gtfs-realtime/reference/Effect for more information
type Effect string

// JourneyStatusXXX are known JourneyStatuse
const (
	// Service suspended.
	EffectNoService Effect = "NO_SERVICE"

	// Service running at lowered capacity.
	JourneyStatusReducedService = "REDUCED_SERVICE"

	// Service running but with substantial delays expected.
	JourneyStatusSignificantDelay = "SIGNIFICANT_DELAY"

	// Service running on alternative routes to avoid problem.
	JourneyStatusDetour = "DETOUR"

	// Service above normal capacity.
	JourneyStatusAdditionalService = "ADDITIONAL_SERVICE"

	// Service different from normal capacity.
	JourneyStatusModifiedService = "MODIFIED_SERVICE"

	// Miscellaneous, undefined Effect.
	JourneyStatusOtherEffect = "OTHER_EFFECT"

	// Default setting: Undetermined or Effect not known.
	JourneyStatusUnknownEffect = "UNKNOWN_EFFECT"

	// Stop not at previous location or stop no longer on route.
	JourneyStatusStopMoved = "STOP_MOVED"
)

// A Disruption reports the specifics of a Disruption
type Disruption struct {
	ID ID `json:"id"` // ID of the Disruption

	// State of the disruption.
	// The state is computed using the application_periods of the disruption and the current time of the query.
	// It can be either "Past", "Active" or "Future"
	Status string `json:"status"`

	InputDisruptionID ID               // For traceability, ID of original input disruption
	InputImpactID     ID               // For traceability: Id of original input impact
	Severity          Severity         `json:"severity"` // Severity gives some categorization element
	Periods           []Period         // Dates where the current disruption is active
	Messages          []Message        // Text to provide to the traveller
	LastUpdated       time.Time        // Last Update of that disruption
	Impacted          []ImpactedObject `json:"impacted_stops"` // Objects impacted
	Cause             string           // The cause of that disruption
	Category          string           // The category of the disruption, optional.
	DisruptionID      string           `json:"disruption_id"`
}

// jsonDisruption define the JSON implementation of Disruption struct
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonDisruption struct {
	// The references
	ID                *ID               `json:"id"`
	Status            *string           `json:"status"`
	InputDisruptionID *ID               `json:"disruption_id"`
	InputImpactID     *ID               `json:"impact_id"`
	Severity          *Severity         `json:"severity"`
	Periods           *[]Period         `json:"application_periods"`
	Messages          *[]Message        `json:"messages"`
	Impacted          *[]ImpactedObject `json:"impacted_objects"`
	Cause             *string           `json:"cause"`
	Category          *string           `json:"category"`

	// Those we will process
	LastUpdated string `json:"updated_at"`
}

type PTObjectDisruptions struct {
	EmbeddedType string `json:"embedded_type"`
	ID           string `json:"id"`
	Quality      int    `json:"quality"`
	Name         string `json:"name"`
	Trip         Trip   `json:"trip"`
}

// UnmarshalJSON implements json.Unmarshaller for a Disruption
func (d *Disruption) UnmarshalJSON(b []byte) error {
	data := &jsonDisruption{
		ID:                &d.ID,
		Status:            &d.Status,
		InputDisruptionID: &d.InputDisruptionID,
		InputImpactID:     &d.InputImpactID,
		Severity:          &d.Severity,
		Periods:           &d.Periods,
		Messages:          &d.Messages,
		Impacted:          &d.Impacted,
		Cause:             &d.Cause,
		Category:          &d.Category,
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Disruption", b}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return fmt.Errorf("error while unmarshalling Disruption: %w", err)
	}

	// Now we process the Update time
	d.LastUpdated, err = parseDateTime(data.LastUpdated)
	if err != nil {
		return gen.err(err, "LastUpdated", "updated_at", data.LastUpdated, "parseDateTime failed")
	}

	// Finished !
	return nil
}
