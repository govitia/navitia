package navitia

import (
	"context"
	"net/url"
	"sort"
	"strconv"

	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

// PlacesResults doesn't have pagination, as the remote API doesn't support it.
//
// PlacesResults can be sorted, it implements sort.Interface.
type PlacesResults struct {
	Places []types.Container `json:"places"`

	Logging `json:"-"`
}

// Len is the number of Places in the results.
func (pr *PlacesResults) Len() int {
	return len(pr.Places)
}

// Less reports if the quality of the Place with the index i is less than that of the Place with the index j
//
// Note: In most use cases, that's the opposite of the desired behaviour, so simply use sort.Reverse and ta-da !
func (pr *PlacesResults) Less(i, j int) bool {
	return pr.Places[i].Quality < pr.Places[j].Quality
}

// Swap swaps the Place of index i and the Place of index j
func (pr *PlacesResults) Swap(i, j int) {
	tmp := pr.Places[i]
	pr.Places[i] = pr.Places[j]
	pr.Places[j] = tmp
}

// PlacesRequest is the query you need to build before passing it to Places
type PlacesRequest struct {
	query string // The search item

	// Types are the type of objects to query
	// It can either be a stop_area, an address, a poi or an administrative_region
	Types []string

	// If given it will filter the search by specific admin uris
	AdminURI []string

	// Enables GeoJSON data in the reply. GeoJSON objects can be VERY large ! >1MB.
	Geo bool

	// If given, it will prioritise objects around these coordinates
	Around types.Coordinates

	// Maximum amount of results
	Count uint

	// Depth can expand the data by making it more verbose.
	// Acceptable values are 0 (light), 1 (regular), 2 (rich), 3 (verbose)
	Depth uint8
}

// toURL formats a Places request to url
func (req PlacesRequest) toURL() (url.Values, error) {
	params := url.Values{
		"q":               []string{req.query},
		"disable_geojson": []string{strconv.FormatBool(!req.Geo)},
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

// places is the internal function used by Places functions
func (s *Session) places(ctx context.Context, url string, query string, opt PlacesRequest) (*PlacesResults, error) {
	var results = &PlacesResults{}

	// Put the query in the PlacesRequest
	opt.query = query

	// Request
	err := s.request(ctx, url, opt, results)
	if err != nil {
		return results, errors.Wrapf(err, "places: error in request for query \"%s\"", query)
	}

	// Sort the places if quality is defined on the results, no need to expand some call
	// Justification for the if condition: If at least of of the results quality is 0, then all of them are 0.
	if results.Len() != 0 && results.Places[0].Quality != 0 {
		sort.Sort(sort.Reverse(results))
	}
	return results, nil
}

const placesEndpoint = "places"

// Places searches in all geographical objects using their names, returning a list of corresponding places.
//
// It is context aware.
func (s *Session) Places(ctx context.Context, query string, opt PlacesRequest) (*PlacesResults, error) {
	// Create the URL
	url := s.apiURL + "/" + placesEndpoint

	// Call
	return s.places(ctx, url, query, opt)
}

// Places searches in all geographical objects within a coverage using their names, returning a list of places.
//
// It is context aware.
func (scope *Scope) Places(ctx context.Context, query string, opt PlacesRequest) (*PlacesResults, error) {
	// Create the URL
	url := scope.baseURL + "/" + placesEndpoint

	// Call
	return scope.session.places(ctx, url, query, opt)
}
