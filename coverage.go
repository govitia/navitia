package gonavitia

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

const coverageEndpoint string = "coverage"

// A CoverageResults holds results for a coverage query
type CoverageResults struct {
	// The list of regions retrieved
	Regions []Region `json:"Regions"`

	// Paging information
	Paging Paging `json:"Pagination"`

	// Timing information
	createdAt   time.Time
	populatedAt time.Time

	// Held session
	session *Session
}

// Coverage lists the areas covered by the Navitia API
// count is the number of items to return, if count=0, then it will return the default number
// BUG: Count doesn't work, server-side
func (s *Session) Coverage(count uint) (*CoverageResults, error) {
	var results = &CoverageResults{session: s, createdAt: time.Now()}

	// Get the request
	url := s.APIURL + "/" + coverageEndpoint
	req, err := s.newRequest(url)
	if err != nil {
		return results, err
	}

	// Add the parameters if count != 0
	if count != 0 {
		params := req.URL.Query()
		countStr := strconv.FormatUint(uint64(count), 10)
		params.Add("count", countStr)
		req.URL.RawQuery = params.Encode()
	}

	// Execute the request
	resp, err := s.client.Do(req)

	// Check it
	if err != nil {
		err = errors.Wrap(err, "Error in request")
		return results, err
	}
	if resp.StatusCode != 200 {
		return results, errors.Errorf("Didn't get right code: %d instead of 200", resp.StatusCode)
	}

	// Parse it
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(results)
	if err != nil {
		return results, errors.Wrap(err, "JSON decoding failed")
	}
	results.populatedAt = time.Now()

	// Return
	return results, nil
}

// RegionByID provides information about a specific region
func (s *Session) RegionByID(id RegionID) (*CoverageResults, error) {
	var results = &CoverageResults{session: s, createdAt: time.Now()}

	// Build the URL
	url := s.APIURL + "/" + coverageEndpoint + "/" + string(id)

	// Get the request
	req, err := s.newRequest(url)
	if err != nil {
		return results, err
	}

	// Execute the request
	resp, err := s.client.Do(req)

	// Check it
	if err != nil {
		err = errors.Wrap(err, "Error in request")
		return results, err
	}
	if resp.StatusCode != 200 {
		return results, errors.Errorf("Didn't get right code: %d instead of 200", resp.StatusCode)
	}

	// Parse it
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(results)
	if err != nil {
		return results, errors.Wrap(err, "JSON decoding failed")
	}
	results.populatedAt = time.Now()

	// Return
	return results, nil
}

// RegionByPos provides information about the region englobing the specific position
func (s *Session) RegionByPos(coords Coordinates) (*CoverageResults, error) {
	var results = &CoverageResults{session: s, createdAt: time.Now()}

	// Build the URL
	url := s.APIURL + "/" + coverageEndpoint + "/" + coords.formatURL()

	// Get the request
	req, err := s.newRequest(url)
	if err != nil {
		return results, err
	}

	// Execute the request
	resp, err := s.client.Do(req)

	// Check it
	if err != nil {
		err = errors.Wrap(err, "Error in request")
		return results, err
	}
	if resp.StatusCode != 200 {
		return results, errors.Errorf("Didn't get right code: %d instead of 200", resp.StatusCode)
	}

	// Parse it
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(results)
	if err != nil {
		return results, errors.Wrap(err, "JSON decoding failed")
	}
	results.populatedAt = time.Now()

	// Return
	return results, nil
}
