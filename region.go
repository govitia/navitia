package navitia

import (
	"net/url"

	"github.com/govitia/navitia/types"
	"github.com/govitia/navitia/utils"
)

const regionEndpoint string = "coverage"

// A RegionResults holds results for a coverage query
// This Results doesn't support paging :(.
type RegionResults struct {
	// The list of regions retrieved
	Regions []types.Region `json:"Regions"`

	// Timing information
	Logging

	// Held session
	session *Session
}

// RegionRequest contains the parameters needed to make a Coverage request.
type RegionRequest struct {
	// Count is the number of items to return, if count=0, then it will return the default number
	// BUG: Count doesn't work, server-side.
	Count uint

	// Enables Geo data (in MKT format) in the reply. Geo objects can be large and slower to parse.
	Geo bool
}

func (req RegionRequest) toURL() (url.Values, error) {
	rb := utils.NewRequestBuilder()

	if req.Count != 0 {
		rb.AddUInt("count", req.Count)
	}

	if !req.Geo {
		rb.AddString("disable_geojson", "true")
	}

	return rb.Values(), nil
}
