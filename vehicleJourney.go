package navitia

import (
	"net/url"
	"time"

	"github.com/govitia/navitia/types"
	"github.com/govitia/navitia/utils"
)

// JourneyResults contains the results of a Journey request
//
// Warning: types.Journey.From / types.Journey.To aren't guaranteed to be filled.
// Based on very basic inspection, it seems they aren't filled when there are sections...
type VehicleJourneyResults struct {
	VehicleJourneys []types.VehicleJourney `json:"vehicle_journeys"`

	Disruptions []types.Disruption `json:"disruptions"`

	Paging Paging `json:"links"`

	Logging `json:"-"`

	session *Session
}

// Count returns the number of results available in a JourneyResults.
func (jr *VehicleJourneyResults) Count() int {
	return len(jr.VehicleJourneys)
}

// VehicleJourneyRequest contain the parameters needed to make a Journey request.
type VehicleJourneyRequest struct {
	ID types.ID
	// There must be at least one From or To parameter defined
	// When used with just one of them, the resulting Journey won't have a populated Sections field.
	From types.ID
	To   types.ID

	// When do you want to depart ? Or is DateIsArrival when do you want to arrive at your destination.
	Date          time.Time
	DateIsArrival bool

	// The traveller's type
	Traveler types.TravelerType

	// Define the freshness of data to use to compute journeys
	Freshness types.DataFreshness

	// Forbidden public transport objects
	Forbidden []types.ID

	// Allowed public transport objects
	// Note: This counstraint intersects with Forbidden
	Allowed []types.ID

	// Force the first section mode if it isn't a public transport mode
	// Note: The parameter is inclusive, not exclusive. As such if you want to forbid a mode
	// you have to include all modes except that one.
	FirstSectionModes []string

	// Same, but for the last section
	LastSectionModes []string

	// MaxDurationToPT is the maximum allowed duration to reach the public transport.
	// Use this to limit the walking/biking part.
	MaxDurationToPT time.Duration

	// These four following parameters set the speed of each mode (Walking, Bike, BSS & car)
	// In meters per second
	WalkingSpeed   float64
	BikeSpeed      float64
	BikeShareSpeed float64
	CarSpeed       float64

	// Minimum and maximum amounts of journeys suggested
	MinJourneys uint
	MaxJourneys uint

	// Count fixes the amount of journeys to be returned, overriding minimum & maximum amount
	// Note: if Count=0 then it isn't taken into account
	Count uint

	// Maximum number of transfers in each journey
	MaxTransfers uint

	// Maximum duration of a trip
	MaxDuration time.Duration // To seconds

	// Wheelchair restricts the answer to accessible public transports
	Wheelchair bool

	// Headsign If given, add a filter on the vehicle journeys that has the
	// given value as headsign (on vehicle journey itself or at a stop time).
	Headsign string

	// Since If given, filter on a period, optional.
	Since time.Time
	// Until, like Since, filter on a period, optional too.
	Until time.Time
}

// toURL formats a journey request to url
// Should be refactored using a switch statement.
func (req VehicleJourneyRequest) toURL() (url.Values, error) {
	rb := utils.NewRequestBuilder()

	// Encode the from and to
	rb.AddString("from", string(req.From))
	rb.AddString("to", string(req.To))

	if !req.Date.IsZero() {
		rb.AddDateTime("datetime", req.Date)
		if req.DateIsArrival {
			rb.AddString("datetime_represents", "arrival")
		}
	}

	rb.AddString("traveler_type", string(req.Traveler))
	rb.AddString("data_freshness", string(req.Freshness))
	rb.AddIDSlice("forbidden_uris[]", req.Forbidden)
	rb.AddIDSlice("allowed_id[]", req.Allowed)
	rb.AddMode("first_section_mode[]", req.FirstSectionModes)
	rb.AddMode("last_section_mode[]", req.LastSectionModes)

	// max_duration_to_pt
	rb.AddInt("max_duration_to_pt", int(req.MaxDurationToPT/time.Second))

	// walking_speed, bike_speed, bss_speed & car_speed
	rb.AddFloat64("walking_speed", req.WalkingSpeed)
	rb.AddFloat64("bike_speed", req.BikeSpeed)
	rb.AddFloat64("bss_speed", req.BikeShareSpeed)
	rb.AddFloat64("car_speed", req.CarSpeed)

	// If count is defined don't bother with the minimimal and maximum amount of items to return
	if count := req.Count; count != 0 {
		rb.AddUInt("count", count)
	} else {
		rb.AddUInt("min_nb_journeys", req.MinJourneys)
		rb.AddUInt("max_nb_journeys", req.MaxJourneys)
	}

	// max_nb_transfers
	rb.AddUInt("max_nb_transfers", req.MaxTransfers)

	// max_duration
	rb.AddInt("max_duration", int(req.MaxDuration/time.Second))

	// headsign
	rb.AddString("headsign", req.Headsign)

	// wheelchair
	if req.Wheelchair {
		rb.AddString("wheelchair", "true")
	}

	rb.AddDateTime("since", req.Since)
	rb.AddDateTime("until", req.Until)

	return rb.Values(), nil
}
