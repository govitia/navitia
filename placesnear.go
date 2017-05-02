package navitia

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/aabizri/navitia/types"
)

// PlacesNearRequest is the query you need to build before passing it to PlacesNear
type PlacesNearRequest struct {
	// Distance from the coordinates or resource from which to look in meters
	// Default value: 500
	Distance uint

	// Type of objects to be queried
	Types []string

	// If given it will filter the search by specific admin uris
	AdminURI []string

	// Enables GeoJSON data in the reply. GeoJSON objects can be VERY large ! >1MB.
	Geo bool

	// Filter the search.
	// Ex: "places_type.id=theater"
	Filter string

	// Maximum amount of results
	Count uint
}

// toURL formats a Places request to url
func (req PlacesNearRequest) toURL() (url.Values, error) {
	params := url.Values{
		"disable_geojson": []string{strconv.FormatBool(!req.Geo)},
	}

	if dist := req.Distance; dist != 0 {
		params.Add("distance", strconv.FormatUint(uint64(dist), 10))
	}

	if len(req.Types) != 0 {
		params["type[]"] = req.Types
	}

	if len(req.AdminURI) != 0 {
		params["admin_uri[]"] = req.AdminURI
	}

	if req.Count != 0 {
		countStr := strconv.FormatUint(uint64(req.Count), 10)
		params["count"] = []string{countStr}
	}
	return params, nil
}

// PlacesNearResults doesn't have pagination, as the remote API doesn't support it.
type PlacesNearResults struct {
	Places  []types.Container `json:"places_nearby"`
	Logging `json:"-"`
}

func (s *Session) placesNear(ctx context.Context, url string, req PlacesNearRequest) (*PlacesNearResults, error) {
	var results = &PlacesNearResults{}
	return results, s.request(ctx, url, req, results)
}

const placesNearEndpoint = "places_nearby"

// PlacesNear searches for places near a point.
func (s *Session) PlacesNear(ctx context.Context, lat, lng float64, req PlacesNearRequest) (*PlacesNearResults, error) {
	// Build the url
	coords := fmt.Sprintf("%3.3f;%3.3f", lat, lng)
	url := s.apiURL + "/coord/" + coords + "/" + placesNearEndpoint

	// Call & return
	return s.placesNear(ctx, url, req)
}

// PlacesNear searches for places near a point in a specific coverage.
func (scope *Scope) PlacesNear(ctx context.Context, lat, lng float64, req PlacesNearRequest) (*PlacesNearResults, error) {
	// Build the url
	coords := fmt.Sprintf("%3.3f;%3.3f", lat, lng)
	url := scope.session.apiURL + "/coverage/" + string(scope.region) + "/" + coords + "/" + placesNearEndpoint

	// Call & return
	return scope.session.placesNear(ctx, url, req)
}

// PlacesNearResource searches for places near a resource in a specific coverage.
func (scope *Scope) PlacesNearResource(ctx context.Context, resourceType string, resourceID types.ID, req PlacesNearRequest) (*PlacesNearResults, error) {
	// Build the url
	url := scope.session.apiURL + "/coverage/" + string(scope.region) + "/" + resourceType + "/" + string(resourceID) + "/" + placesNearEndpoint

	// Call & return
	return scope.session.placesNear(ctx, url, req)
}
