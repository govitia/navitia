package navitia

import (
	"context"

	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

type coordsResults struct {
	Regions []types.ID    `json:"regions"`
	Address types.Address `json:"address"`
	Logging `json:"-"`
}

const coordsEndpoint = "coords"

// Coords , given some coordinates, answers you
// 	- Your detailed postal address
// 	- The right coverage, that is the region ID that can be used to scope future requests
func (s *Session) Coords(ctx context.Context, coords types.Coordinates) (address *types.Address, regionID types.ID, err error) {
	// Build the URL
	url := s.apiURL + "/" + coordsEndpoint + "/" + string(coords.ID())

	// Create the result value
	res := &coordsResults{}

	// Launch the request
	err = s.requestURL(ctx, url, res)
	if err != nil {
		return nil, "", errors.Wrap(err, "Coords: error while requesting")
	}

	// Validate the response
	if res == nil || len(res.Regions) == 0 {
		return nil, "", errors.Wrap(err, "Coords: invalid response")
	}

	// Return the correct values
	return &res.Address, res.Regions[0], nil
}
