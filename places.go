package navitia

import (
	"context"
	"fmt"
	"github.com/aabizri/navitia/types"
	"net/url"
	"strconv"
)

// PlacesResults doesn't have pagination
type PlacesResults struct {
	Places []types.Place

	Logging
	session *Session
}

// String implements Stringer and pretty-prints a PlacesResults
func (res PlacesResults) String() string {
	var msg string
	for i, place := range res.Places {
		if place != nil {
			msg += fmt.Sprintf("Place #%d (%s): %s\n", i, place.PlaceType(), place.String())
		}
	}
	return msg
}

// PlacesRequest is the query you need to build before passing it to Places
type PlacesRequest struct {
	Query string // The search item

	// Types are the type of objects to query
	// It can either be a stop_area, an address, a poi or an administrative_region
	Types []string

	AdminURI []string // If given it will filter the search by specific admin uris

	DisableGeoJSON bool

	Around types.Coordinates // If given, it will prioritize objects around these coordinates

	// Maximum amount of results
	Count uint
}

// toURL formats a Places request to url
func (req PlacesRequest) toURL() (url.Values, error) {
	params := url.Values{
		"q": []string{req.Query},
	}

	if len(req.Types) != 0 {
		params["type[]"] = req.Types
	}

	if len(req.AdminURI) != 0 {
		params["admin_uri[]"] = req.AdminURI
	}

	if req.DisableGeoJSON {
		params["disable_geojson"] = []string{"true"}
	}

	if req.Count != 0 {
		countStr := strconv.FormatUint(uint64(req.Count), 10)
		params["count"] = []string{countStr}
	}
	return params, nil
}

// places is the internal function used by Places functions
func (s *Session) places(ctx context.Context, url string, params PlacesRequest) (*PlacesResults, error) {
	var results = &PlacesResults{session: s}
	err := s.request(ctx, url, params, results)
	return results, err
}

const placesEndpoint = "places"

// Places searches in all geographical objects using their names, returning a list of corresponding places.
//
// It is context aware.
func (s *Session) Places(ctx context.Context, params PlacesRequest) (*PlacesResults, error) {
	// Create the URL
	url := s.APIURL + "/" + placesEndpoint

	// Call
	return s.places(ctx, url, params)
}

// PlacesR searches in all geographical objects within a given region using their names, returning a list of places.
//
// It is context aware.
func (s *Session) PlacesR(ctx context.Context, params PlacesRequest, regionID string) (*PlacesResults, error) {
	// Create the URL
	url := s.APIURL + "/coverage/" + regionID + "/" + placesEndpoint

	// Call
	return s.places(ctx, url, params)
}
