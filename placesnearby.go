package navitia

import (
	"context"
	"net/url"
	"strconv"

	"github.com/aabizri/navitia/types"
)

// PlacesNearbyRequest is the query you need to build before passing it to PlacesNearby
type PlacesNearbyRequest struct {
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

	// Depth can expand the data by making it more verbose.
	// Acceptable values are 0 (light), 1 (regular), 2 (rich), 3 (verbose)
	Depth uint8
}

// toURL formats a Places request to url
func (req PlacesNearbyRequest) toURL() (url.Values, error) {
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

// PlacesNearbyResults doesn't have pagination, as the remote API doesn't support it.
type PlacesNearbyResults struct {
	Places  []types.Container `json:"places_nearby"`
	Logging `json:"-"`
}

func (s *Session) placesNearby(ctx context.Context, url string, req PlacesNearbyRequest) (*PlacesNearbyResults, error) {
	var results = &PlacesNearbyResults{}
	return results, s.request(ctx, url, req, results)
}

const placesNearbyEndpoint = "places_nearby"

// PlacesNearby searches for places near a point.
// It acts similarly to session.Scope(Coords.ID()).PlacesNearby(Coords.ID())
func (s *Session) PlacesNearby(ctx context.Context, coords types.Coordinates, req PlacesNearbyRequest) (*PlacesNearbyResults, error) {
	// Build the url
	url := s.apiURL + "/coord/" + string(coords.ID()) + "/" + placesNearbyEndpoint

	// Call & return
	return s.placesNearby(ctx, url, req)
}

// PlacesNearby searches for places near a point in a specific coverage.
func (scope *Scope) PlacesNearby(ctx context.Context, coords types.Coordinates, req PlacesNearbyRequest) (*PlacesNearbyResults, error) {
	// Build the url
	url := scope.baseURL + "/" + string(coords.ID()) + "/" + placesNearbyEndpoint

	// Call & return
	return scope.session.placesNearby(ctx, url, req)
}

// PlacesNearbyResource searches for places near a resource in a specific coverage.
func (scope *Scope) PlacesNearbyResource(ctx context.Context, resourceType string, resourceID types.ID, req PlacesNearbyRequest) (*PlacesNearbyResults, error) {
	// Build the url
	url := scope.baseURL + "/" + resourceType + "/" + string(resourceID) + "/" + placesNearbyEndpoint

	// Call & return
	return scope.session.placesNearby(ctx, url, req)
}
