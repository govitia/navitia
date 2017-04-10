package gonavitia

import (
	"./types"
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
	"strconv"
	"time"
)

// JourneyResults countains the results of a Journey request
type JourneyResults struct {
	Journeys []types.Journey

	createdAt   time.Time
	populatedAt time.Time
}

// JourneyRequest countain the parameters needed to make a Journey request
type JourneyRequest struct {
	// There must be at least one From or To parameter defined
	// When used with just one of them, the resulting Journey won't have a populated Sections field.
	From types.PositionID
	To   types.PositionID

	Date          time.Time
	DateIsArrival bool

	// The traveler's type
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
	FirstSectionModes []types.Mode

	// Same, but for the last section
	LastSectionModes []types.Mode

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
				params.Add(key, id.FormatURL())
			}
		}
	}
	addModes := func(key string, modes []types.Mode) {
		if len(modes) != 0 {
			for _, mode := range modes {
				params.Add(key, string(mode))
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
	if from := req.From; from != nil {
		formatted, err := from.FormatURL()
		if err != nil {
			return params, errors.Wrap(err, "error while formatting from field")
		}
		params.Add("from", formatted)
	}
	if to := req.To; to != nil {
		formatted, err := to.FormatURL()
		if err != nil {
			return params, errors.Wrap(err, "error while formatting to field")
		}
		params.Add("to", formatted)
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

const journeysEndpoint string = "journeys"

// Journeys computes a list of journeys according to the parameters given
func (s *Session) Journeys(params JourneyRequest) (*JourneyResults, error) {
	var results = &JourneyResults{createdAt: time.Now()}

	// Get the request
	url := s.APIURL + "/" + journeysEndpoint
	req, err := s.newRequest(url)
	if err != nil {
		return results, errors.Wrap(err, "error while creating request")
	}

	// Encode the parameters
	values, err := params.toURL()
	if err != nil {
		return results, errors.Wrap(err, "error while retrieving url values to be encoded")
	}
	req.URL.RawQuery = values.Encode()

	// Execute the request
	resp, err := s.client.Do(req)

	// Check it
	if err != nil {
		return results, errors.Wrap(err, "errror while executing request")
	}
	if resp.StatusCode != 200 {
		return results, parseRemoteError(resp)
	}

	// Parse it
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(results)
	if err != nil {
		return results, errors.Wrap(err, "JSON decoding failed")
	}
	results.populatedAt = time.Now()

	// Return
	return results, err
}
