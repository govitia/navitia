package navitia

import (
	"context"
	"net/url"
	"strconv"

	"github.com/aabizri/navitia/types"
)

const regionEndpoint string = "coverage"

// A RegionResults holds results for a coverage query
//
// This Results doesn't support paging :(
type RegionResults struct {
	// The list of regions retrieved
	Regions []types.Region `json:"Regions"`

	// Timing information
	Logging

	// Held session
	session *Session
}

// RegionRequest contains the parameters needed to make a Coverage request
type RegionRequest struct {
	// Count is the number of items to return, if count=0, then it will return the default number
	// BUG: Count doesn't work, server-side.
	Count uint

	// Enables Geo data (in MKT format) in the reply. Geo objects can be large and slower to parse.
	Geo bool
}

func (req RegionRequest) toURL() (url.Values, error) {
	params := url.Values{}

	if count := req.Count; count != 0 {
		countStr := strconv.FormatUint(uint64(count), 10)
		params["count"] = []string{countStr}
	}

	if !req.Geo {
		params["disable_geojson"] = []string{"true"}
	}

	return params, nil
}

func (s *Session) region(ctx context.Context, url string, params RegionRequest) (*RegionResults, error) {
	var results = &RegionResults{session: s}
	err := s.request(ctx, url, params, results)
	return results, err
}

// Regions lists the areas covered by the Navitia API.
// i.e it returns the coverage of the API.
//
// It is context aware.
func (s *Session) Regions(ctx context.Context, req RegionRequest) (*RegionResults, error) {
	// Create the URL
	url := s.apiURL + "/" + regionEndpoint

	// Call and return
	return s.region(ctx, url, req)
}

// RegionByID provides information about a specific region.
//
// If the ID provided isn't an ID of a region, this WILL fail.
//
// It is context aware.
func (s *Session) RegionByID(ctx context.Context, req RegionRequest, id types.ID) (*RegionResults, error) {
	// Build the URL
	url := s.apiURL + "/" + regionEndpoint + "/" + string(id)

	// Call and return
	return s.region(ctx, url, req)
}

// RegionByPos provides information about the region englobing the specific position.
//
// It is context aware.
func (s *Session) RegionByPos(ctx context.Context, req RegionRequest, coords types.Coordinates) (*RegionResults, error) {
	// Build the URL
	coordsQ := string(coords.ID())
	url := s.apiURL + "/" + regionEndpoint + "/" + coordsQ

	// Call and return
	return s.region(ctx, url, req)
}
