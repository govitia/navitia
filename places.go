package navitia

import (
	"net/url"

	"github.com/govitia/navitia/types"
	"github.com/govitia/navitia/utils"
)

const placesEndpoint = "places"

// PlacesResults doesn't have pagination, as the remote API doesn't support it.
// PlacesResults can be sorted, it implements sort.Interface.
type PlacesResults struct {
	Places []types.Container `json:"places"`

	Logging `json:"-"`

	session *Session
}

// Len is the number of Places in the results.
func (pr *PlacesResults) Len() int {
	return len(pr.Places)
}

// Less reports if the quality of the Place with the index i is less than that of the Place with the index j
// Note: In most use cases, that's the opposite of the desired behaviour, so simply use sort.Reverse and ta-da !
func (pr *PlacesResults) Less(i, j int) bool {
	return pr.Places[i].Quality < pr.Places[j].Quality
}

// Swap swaps the Place of index i and the Place of index j.
func (pr *PlacesResults) Swap(i, j int) {
	pr.Places[i], pr.Places[j] = pr.Places[j], pr.Places[i]
}

// PlacesRequest is the query you need to build before passing it to Places.
type PlacesRequest struct {
	Query string // The search item

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
}

// toURL formats a Places request to url.
func (req PlacesRequest) toURL() (url.Values, error) {
	rb := utils.NewRequestBuilder()

	rb.AddString("q", req.Query)
	rb.AddStringSlice("type[]", req.Types)
	rb.AddStringSlice("admin_uri[]", req.AdminURI)

	if !req.Geo {
		rb.AddString("disable_geojson", "true")
	}

	if req.Count != 0 {
		rb.AddUInt("count", req.Count)
	}
	return rb.Values(), nil
}
