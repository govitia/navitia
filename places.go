package gonavitia

import (
	"github.com/aabizri/gonavitia/types"
	"net/url"
)

// PlacesResults doesn't have pagination
type PlacesResults struct {
	Places []types.Place

	Logging
	session *Session
}

type PlacesRequest struct {
	Query string // The search item

	// Types are the type of objects to query
	// It can either be a stop_area, an address, a poi or an administrative_region
	Types []string

	AdminURI []string // If given it will filter the search by specific admin uris

	DisableGeoJson bool

	Around types.Coordinates // If given, it will prioritize objects around these coordinates
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

	if req.DisableGeoJson {
		params["disable_geojson"] = []string{"true"}
	}

	return params, nil
}

// places is the internal function used by Places functions
func (s *Session) places(url string, params PlacesRequest) (*PlacesResults, error) {
	var results = &PlacesResults{session: s}
	err := s.request(url, params, results)
	return results, err
}

const placesEndpoint = "places"

// Places search in all geographical objects using their names, returning a list of corresponding places.
func (s *Session) Places(params PlacesRequest) (*PlacesResults, error) {
	// Create the URL
	url := s.APIURL + "/" + placesEndpoint

	// Call
	return s.places(url, params)
}

// Places search in all geographical objects within a given region using their names, returning a list of places.
func (s *Session) PlacesR(params PlacesRequest, regionID string) (*PlacesResults, error) {
	// Create the URL
	url := s.APIURL + "/coverage/" + regionID + "/" + placesEndpoint

	// Call
	return s.places(url, params)
}
