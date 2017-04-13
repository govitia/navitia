package navitia

import (
	"github.com/aabizri/navitia/types"
	"net/url"
	"strconv"
)

const coverageEndpoint string = "coverage"

// A CoverageResults holds results for a coverage query
type CoverageResults struct {
	// The list of regions retrieved
	Regions []types.Region `json:"Regions"`

	// Paging information
	Paging Paging `json:"Pagination"`

	// Timing information
	Logging

	// Held session
	session *Session
}

type coverageRequest struct {
	Count uint
}

func (req coverageRequest) toURL() (url.Values, error) {
	params := url.Values{}

	if count := req.Count; count != 0 {
		countStr := strconv.FormatUint(uint64(count), 10)
		params["count"] = []string{countStr}
	}

	return params, nil
}

func (s *Session) coverage(url string, params coverageRequest) (*CoverageResults, error) {
	var results = &CoverageResults{session: s}
	err := s.request(url, params, results)
	return results, err
}

// Coverage lists the areas covered by the Navitia API
// count is the number of items to return, if count=0, then it will return the default number
// BUG: Count doesn't work, server-side
func (s *Session) Coverage(count uint) (*CoverageResults, error) {
	// Create the URL
	url := s.APIURL + "/" + coverageEndpoint

	// Create the query
	params := coverageRequest{Count: count}

	// Call and return
	return s.coverage(url, params)
}

// RegionByID provides information about a specific region
// If the ID provided isn't an ID of a region, this WILL fail.
func (s *Session) RegionByID(id types.ID) (*CoverageResults, error) {
	// Build the URL
	url := s.APIURL + "/" + coverageEndpoint + "/" + string(id)

	// Create the query
	params := coverageRequest{}

	// Call and return
	return s.coverage(url, params)
}

// RegionByPos provides information about the region englobing the specific position
func (s *Session) RegionByPos(coords types.Coordinates) (*CoverageResults, error) {
	// Build the URL
	coordsFormatted := coords.QueryEscape()
	url := s.APIURL + "/" + coverageEndpoint + "/" + coordsFormatted

	// Create the query
	params := coverageRequest{}

	// Call and return
	return s.coverage(url, params)
}
