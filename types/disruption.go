package types

import (
	"image/color"
	"time"
)

// Effect codes for known journey status information
// For example, reduced service, detours or moved stops.
//
// See https://developers.google.com/transit/gtfs-realtime/reference/Effect for more information.
type Effect string

// JourneyStatusXXX are known JourneyStatuse.
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

// A Disruption reports the specifics of a Disruption.
type Disruption struct {
	// ID of the Disruption
	ID ID `json:"id"`

	// State of the disruption.
	// The state is computed using the application_periods of the disruption and the current time of the query.
	//
	// It can be either "Past", "Active" or "Future"
	Status string `json:"status"`

	// For traceability, ID of original input disruption
	InputDisruptionID ID

	// For traceability: Id of original input impact
	InputImpactID ID

	// Severity gives some categorization element
	Severity Severity `json:"severity"`

	// Dates where the current disruption is active
	Periods []Period

	// Text to provide to the traveller
	Messages []Message

	// Last Update of that disruption
	LastUpdated time.Time

	// Objects impacted
	Impacted []ImpactedObject `json:"impacted_stops"`

	// The cause of that disruption
	Cause string

	// The category of the disruption.
	// Optional.
	Category string

	DisruptionID string `json:"disruption_id"`
}

// A Message contains the text to be provided to the traveler.
type Message struct {
	// The message to bring to the traveler
	Text string `json:"text"`

	// The destination media for this Message.
	Channel *Channel `json:"channel"`
}

// A Channel is a destination media for a message.
type Channel struct {
	// ID of the address
	ID ID `json:"id"`

	// Content Type (text/html etc.) RFC1341.4
	ContentType string `json:"content_type"`

	// Name of the channel
	Name string `json:"name"`

	// Types ?
	Types []string `json:"types,omitempty"`
}

// An ImpactedObject describes a PTObject impacted by a Disruption with some additional info.
type ImpactedObject struct {
	// The impacted public transport object
	Object Container `json:"pt_object"`

	// Only for line section impact, the impacted section
	ImpactedSection ImpactedSection `json:"impacted_section"`

	// Only for Trip delay, the list of delays, stop by stop
	ImpactedStops []ImpactedStop `json:"impacted_stops"`
}

// An ImpactedSection records the impact to a section.
type ImpactedSection struct {
	// The start of the disruption, spatially
	From Container `json:"from"`
	// Until this point
	To Container `json:"to"`
}

// An ImpactedStop records the impact to a stop.
type ImpactedStop struct {
	// The impacted stop point of the trip
	Point StopPoint `json:"stop_point"`

	// New departure hour (format HHMMSS) of the trip on this stop point
	NewDeparture string

	// New arrival hour (format HHMMSS) of the trip on this stop point
	NewArrival string

	// Base departure hour (format HHMMSS) of the trip on this stop point
	BaseDeparture string

	// Base arrival hour (format HHMMSS) of the trip on this stop point
	BaseArrival string

	// Cause of the modification
	Cause string `json:"cause"`

	// Effect on that StopPoint
	// Can be "added", "deleted", "delayed"
	Effect string

	AmendedArrivalTime string `json:"amended_arrival_time"`

	StopTimeEffect string `json:"stop_time_effect"`

	DepartureStatus string `json:"departure_status"`

	IsDetour bool `json:"is_detour"`

	AmendedDepartureTime string `json:"amended_departure_time"`

	BaseArrivalTime string `json:"base_arrival_time"`

	BaseDepartureTime string `json:"base_departure_time"`

	ArrivalStatus string `json:"arrival_status"`
}

type PTObjectDisruptions struct {
	EmbeddedType string `json:"embedded_type"`
	ID           string `json:"id"`
	Quality      int    `json:"quality"`
	Name         string `json:"name"`
	Trip         Trip   `json:"trip"`
}

// Period of effect.
type Period struct {
	Begin time.Time `json:"begin"`
	End   time.Time `json:"end"`
}

// Severity object can be used to make visual grouping.
type Severity struct {
	// Name of severity
	Name string `json:"name"`

	// Priority of the severity. Given by the agency. 0 is the strongest priority, a nil Priority means its undefined (duh).
	Priority *int `json:"priority"`

	// HTML color for classification
	Color color.Color `json:"color"`

	// Effect: Normalized value of the effect on the public transport object
	Effect Effect `json:"effect"`
}
