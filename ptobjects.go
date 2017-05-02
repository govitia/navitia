package navitia

import (
	"context"
	"net/url"
	"sort"
	"strconv"

	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

// PTObjectsResults doesn't have pagination, as the remote API doesn't support it.
//
// PTObjectsResults can be sorted, it implements sort.Interface.
type PTObjectsResults struct {
	PTObjects []types.Container `json:"pt_objects"`

	Logging `json:"-"`
}

// Len is the number of PTObjects in the results.
func (pr *PTObjectsResults) Len() int {
	return len(pr.PTObjects)
}

// Less reports if the quality of the Place with the index i is less than that of the Place with the index j
//
// Note: In most use cases, that's the opposite of the desired behaviour, so simply use sort.Reverse and ta-da !
func (pr *PTObjectsResults) Less(i, j int) bool {
	return pr.PTObjects[i].Quality < pr.PTObjects[j].Quality
}

// Swap swaps the Place of index i and the Place of index j
func (pr *PTObjectsResults) Swap(i, j int) {
	tmp := pr.PTObjects[i]
	pr.PTObjects[i] = pr.PTObjects[j]
	pr.PTObjects[j] = tmp
}

// PTObjectsRequest is the query you need to build before passing it to PTObjects
type PTObjectsRequest struct {
	Query string // The search item

	// Types are the type of objects to query
	// It can either be a stop_area, an address, a poi or an administrative_region
	Types []string

	// Enables GeoJSON data in the reply. GeoJSON objects can be VERY large ! >1MB.
	Geo bool

	// Maximum amount of results
	Count uint
}

// toURL formats a PTObjects request to url
func (req PTObjectsRequest) toURL() (url.Values, error) {
	params := url.Values{
		"q": []string{req.Query},
	}

	if len(req.Types) != 0 {
		params["type[]"] = req.Types
	}

	if !req.Geo {
		params["disable_geojson"] = []string{"true"}
	}

	if req.Count != 0 {
		countStr := strconv.FormatUint(uint64(req.Count), 10)
		params["count"] = []string{countStr}
	}

	return params, nil
}

const ptObjectsEndpoint = "pt_objects"

// PTObjects searches in all public transportation objects within a coverage using their names, returning a list of ptObjects.
//
// Different types can be returned:
// - types.Network
// - types.CommercialMode
// - types.Line
// - types.Route
// - types.StopArea
func (scope *Scope) PTObjects(ctx context.Context, params PTObjectsRequest) (*PTObjectsResults, error) {
	// Create the URL
	url := scope.session.apiURL + "/coverage/" + string(scope.region) + "/" + ptObjectsEndpoint

	// Create the results
	var results = &PTObjectsResults{}
	err := scope.session.request(ctx, url, params, results)
	if err != nil {
		return nil, errors.Wrap(err, "PTObjects: error while executing request")
	}

	// Sort the ptObjects if quality is defined on the results, no need to expand some call
	// Justification for the if condition: If at least of of the results quality is 0, then all of them are 0.
	if results.Len() != 0 && results.PTObjects[0].Quality != 0 {
		sort.Sort(sort.Reverse(results))
	}

	// Call
	return results, nil
}
