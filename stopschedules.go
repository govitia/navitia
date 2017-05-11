package navitia

import (
	"context"
	"net/url"
	"strconv"
	"time"
	"unsafe"

	"github.com/aabizri/navitia/internal/unmarshal"
	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

type StopSchedulesResults struct {
	StopSchedules []types.StopSchedule `json:"stop_schedules"`

	Notes []types.Note `json:"notes"`

	Disruptions []types.Disruption `json:"disruptions"`

	Paging Paging `json:"links"`

	Logging
}

type StopSchedulesRequest struct {
	// The date/time from which the results will be returned
	// Default value: Now
	From time.Time

	// The maximum duration in seconds of values being returned
	// Default value: 86400 seconds = 24 hours.
	Duration time.Duration

	// Lines, modes & networks to avoid in the results
	Forbidden []types.ID

	// Items per stop schedule
	ItemsPerSchedule uint

	// Data Freshness
	Freshness types.DataFreshness

	// Enable GEOJson in the results
	Geo bool
}

// toURL formats a StopSchedulesRequest to url
func (req StopSchedulesRequest) toURL() (url.Values, error) {
	params := url.Values{
		"disable_geojson": []string{strconv.FormatBool(!req.Geo)},
	}

	if !req.From.IsZero() {
		params.Add("from_datetime", req.From.Format(unmarshal.DateTimeFormat))
	}

	if req.Duration != 0 {
		durationInSeconds := req.Duration / time.Second
		params.Add("duration", strconv.FormatInt(int64(durationInSeconds), 10))
	}

	if len(req.Forbidden) != 0 {
		magic := *(*[]string)(unsafe.Pointer(&req.Forbidden))
		params["forbidden_uris[]"] = magic
	}

	if req.ItemsPerSchedule != 0 {
		ipsStr := strconv.FormatUint(uint64(req.ItemsPerSchedule), 10)
		params["items_per_schedule"] = []string{ipsStr}
	}

	return params, nil
}

const stopSchedulesEndpoint string = "stop_schedules"

// stopSchedules is the internal function used by StopSchedules functions
func (s *Session) stopSchedules(ctx context.Context, url string, opts StopSchedulesRequest) (*StopSchedulesResults, error) {
	var results = &StopSchedulesResults{}
	err := s.request(ctx, url, opts, results)

	return results, errors.Wrapf(err, "error in call to stopSchedules for url %s", url)
}

// StopSchedules returns the stop schedules for a specific resource. For the coordinates-enabled method, see StopSchedulesCoords.
func (scope *Scope) StopSchedules(ctx context.Context, resID types.ID, opts StopSchedulesRequest) (*StopSchedulesResults, error) {
	// Build the URL
	resType, err := resID.Type()
	if err != nil {
		return nil, errors.Wrapf(err, "StopSchedules: couldn't extract resource type from resource id \"%s\"", resID)
	}
	resSelector := resourceTypeToSelector[resType]
	url := scope.baseURL + "/" + resSelector + "/" + string(resID) + "/" + stopSchedulesEndpoint

	// Call and return
	return scope.session.stopSchedules(ctx, url, opts)
}

// StopSchedulesExplicit returns the stop schedules for a specific resource given an explcit resource type.
func (scope *Scope) StopSchedulesExplicit(ctx context.Context, resType string, resID types.ID, opts StopSchedulesRequest) (*StopSchedulesResults, error) {
	// Build the URL
	resSelector := resourceTypeToSelector[resType]
	url := scope.baseURL + "/" + resSelector + "/" + string(resID) + "/" + stopSchedulesEndpoint

	// Call and return
	return scope.session.stopSchedules(ctx, url, opts)
}

// StopSchedulesCoords returns the stop schedules for a specific resource associated with some coordinates.
func (scope *Scope) StopSchedulesCoords(ctx context.Context, coords types.Coordinates, opts StopSchedulesRequest) (*StopSchedulesResults, error) {
	// Build the URL
	url := scope.baseURL + "/coords/" + string(coords.ID()) + "/" + stopSchedulesEndpoint

	// Call and return
	return scope.session.stopSchedules(ctx, url, opts)
}
