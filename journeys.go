package navitia

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/aabizri/navitia/types"
)

// JourneyResults contains the results of a Journey request
//
// Warning: types.Journey.From / types.Journey.To aren't guaranteed to be filled.
// Based on very basic inspection, it seems they aren't filled when there are sections...
type JourneyResults struct {
	Journeys []types.Journey `json:"journeys"`

	Paging Paging `json:"links"`

	Logging `json:"-"`

	session *Session
}

// Len returns the number of results available in a JourneyResults
func (jr *JourneyResults) Len() int {
	return len(jr.Journeys)
}

// JourneyRequest contain the parameters needed to make a Journey request
type JourneyRequest struct {
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
	// Note: The parameter is inclusive, not exclusive. As such if you want to forbid a mode you have to include all modes except that one.
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
}

// toURL formats a journey request to url
// Should be refactored using a switch statement
func (req JourneyRequest) toURL() (url.Values, error) {
	params := url.Values{}

	// Define a few useful functions
	addUint := func(key string, amount uint64) {
		if amount != 0 {
			str := strconv.FormatUint(amount, 10)
			params.Add(key, str)
		}
	}
	addInt := func(key string, amount int64) {
		if amount != 0 {
			str := strconv.FormatInt(amount, 10)
			params.Add(key, str)
		}
	}
	addString := func(key string, str string) {
		if str != "" {
			params.Add(key, str)
		}
	}
	addIDSlice := func(key string, ids []types.ID) {
		if len(ids) != 0 {
			for _, id := range ids {
				params.Add(key, string(id))
			}
		}
	}
	addModes := func(key string, modes []string) {
		if len(modes) != 0 {
			for _, mode := range modes {
				params.Add(key, mode)
			}
		}
	}
	addFloat := func(key string, amount float64) {
		if amount != 0 {
			speedStr := strconv.FormatFloat(amount, 'f', 3, 64)
			params.Add(key, speedStr)
		}
	}

	// Encode the from and to
	if from := req.From; from != "" {
		params.Add("from", string(from))
	}
	if to := req.To; to != "" {
		params.Add("to", string(to))
	}

	if datetime := req.Date; !datetime.IsZero() {
		str := datetime.Format(types.DateTimeFormat)
		params.Add("datetime", str)
		if req.DateIsArrival {
			params.Add("datetime_represents", "arrival")
		}
	}

	addString("traveler_type", string(req.Traveler))

	addString("data_freshness", string(req.Freshness))

	addIDSlice("forbidden_uris[]", req.Forbidden)

	addIDSlice("allowed_id[]", req.Allowed)

	addModes("first_section_mode[]", req.FirstSectionModes)

	addModes("last_section_mode[]", req.LastSectionModes)

	// max_duration_to_pt
	addInt("max_duration_to_pt", int64(req.MaxDurationToPT/time.Second))

	// walking_speed, bike_speed, bss_speed & car_speed
	addFloat("walking_speed", req.WalkingSpeed)
	addFloat("bike_speed", req.BikeSpeed)
	addFloat("bss_speed", req.BikeShareSpeed)
	addFloat("car_speed", req.CarSpeed)

	// If count is defined don't bother with the minimimal and maximum amount of items to return
	if count := req.Count; count != 0 {
		addUint("count", uint64(count))
	} else {
		addUint("min_nb_journeys", uint64(req.MinJourneys))
		addUint("max_nb_journeys", uint64(req.MaxJourneys))
	}

	// max_nb_transfers
	addUint("max_nb_transfers", uint64(req.MaxTransfers))

	// max_duration
	addInt("max_duration", int64(req.MaxDuration/time.Second))

	// wheelchair
	if req.Wheelchair {
		params.Add("wheelchair", "true")
	}

	return params, nil
}

// journeys is the internal function used by Journeys functions
func (s *Session) journeys(ctx context.Context, url string, req JourneyRequest) (*JourneyResults, error) {
	var results = &JourneyResults{session: s}
	err := s.request(ctx, url, req, results)
	return results, err
}

const journeysEndpoint string = "journeys"

// Journeys computes a list of journeys according to the parameters given
func (s *Session) Journeys(ctx context.Context, req JourneyRequest) (*JourneyResults, error) {
	// Create the URL
	url := s.apiURL + "/" + journeysEndpoint

	// Call
	return s.journeys(ctx, url, req)
}

// Journeys computes a list of journeys according to the parameters given in a specific scope
func (scope *Scope) Journeys(ctx context.Context, req JourneyRequest) (*JourneyResults, error) {
	// Create the URL
	url := scope.session.apiURL + "/coverage/" + string(scope.region) + "/" + journeysEndpoint

	// Call
	return scope.session.journeys(ctx, url, req)
}
