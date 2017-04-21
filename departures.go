package navitia

import (
	"context"
	"net/url"
	"strconv"
	"time"
	"unsafe"

	"github.com/aabizri/navitia/types"
)

// DeparturesResults holds the results of a departures request.
type DeparturesResults struct {
	Departures []struct {
		Display   types.Display
		StopPoint types.StopPoint
		Route     types.Route
		//StopDateTime
	} `json:"departures"`

	Paging Paging `json:"links"`

	Logging `json:"-"`
}

// DeparturesRequest contains the optional parameters for a Departures request.
type DeparturesRequest struct {
	// From what time on do you want to see the results ?
	From time.Time

	// Maximum duration between From and the retrived results.
	//
	// Default value is 24 hours
	Duration time.Duration

	// The maximum amount of results
	//
	// Default value is 10 results
	Count uint

	// ForbiddenURIs
	Forbidden []types.ID

	// Freshness of the data
	Freshness types.DataFreshness

	// Enable GEO information in the results (heavier & slower)
	GEO bool
}

func (req DeparturesRequest) toURL() (url.Values, error) {
	values := url.Values{}

	if datetime := req.From; !datetime.IsZero() {
		str := datetime.Format(types.DateTimeFormat)
		values.Add("datetime", str)
	}

	// If count is defined don't bother with the minimimal and maximum amount of items to return
	if count := req.Count; count != 0 {
		countStr := strconv.FormatUint(uint64(count), 10)
		values.Add("count", countStr)
	}

	// Deal with the forbidden URIs
	if forbidden := req.Forbidden; len(forbidden) != 0 {
		magic := *(*[]string)(unsafe.Pointer(&forbidden))
		values["forbidden_uris[]"] = magic
	}

	// Set the freshness
	if freshness := req.Freshness; freshness != "" {
		values.Add("data_freshness", string(freshness))
	}

	// Add GEO
	if !req.GEO {
		values.Add("disable_geojson", "true")
	}

	return values, nil
}

// departures is the internal function used by Departures functions
func (s *Session) departures(ctx context.Context, url string, req DeparturesRequest) (*DeparturesResults, error) {
	var results = &DeparturesResults{}
	err := s.request(ctx, url, req, results)
	return results, err
}

const departuresEndpoint string = "departures"

// DeparturesSA requests the departures for a given StopArea in a given region.
func (s *Session) DeparturesSA(ctx context.Context, req DeparturesRequest, region types.ID, resource types.ID) (*DeparturesResults, error) {
	// Create the URL
	url := s.APIURL + "/coverage/" + string(region) + "/stop_areas/" + string(resource) + "/" + departuresEndpoint

	return s.departures(ctx, url, req)
}

// DeparturesSP requests the departures for a given StopPoint in a given region.
func (s *Session) DeparturesSP(ctx context.Context, req DeparturesRequest, region types.ID, resource types.ID) (*DeparturesResults, error) {
	// Create the URL
	url := s.APIURL + "/coverage/" + string(region) + "/stop_points/" + string(resource) + "/" + departuresEndpoint

	return s.departures(ctx, url, req)
}

// DeparturesC requests the departures from a point described by coordinates.
func (s *Session) DeparturesC(ctx context.Context, req DeparturesRequest, coords types.Coordinates) (*DeparturesResults, error) {
	// Create the URL
	coordsQ := url.PathEscape(coords.QueryEscape())
	url := s.APIURL + "/coverage/" + coordsQ + "/coords/" + coordsQ + "/" + departuresEndpoint

	return s.departures(ctx, url, req)
}
