package navitia

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

// ExploreResults doesn't have pagination, as the remote API doesn't support it.
type ExploreResults struct {
	// PTObjects are of one of these types, according to the request that was given.
	//	-[]types.CommercialMode
	// 	-[]types.Line
	//	-[]types.Network
	// 	-[]types.Route
	// 	-[]types.StopArea
	// 	-[]types.StopPoint
	// 	-[]types.PhysicalMode
	// 	-[]types.Company
	// 	-[]types.VehicleJourney [NOT IMPLEMENTED YET]
	// 	-[]types.Disruption
	PTObjects interface{}

	Logging `json:"-"`
}

// UnmarshalJSON implements unmarshalling for ExploreResults
func (sasr *ExploreResults) UnmarshalJSON(b []byte) error {
	// first will hold the preliminary values
	first := make(map[string]json.RawMessage)

	// Unmarshal to first
	err := json.Unmarshal(b, &first)
	if err != nil {
		return errors.Wrap(err, "error in first-pass unmarshalling")
	}
	// Create a value
	var (
		recv   interface{}
		second json.RawMessage
	)

	// Switch on it
	for k := range first {
		switch k {
		case CommercialModesSelector:
			recv = []types.CommercialMode{}
		case LinesSelector:
			recv = []types.Line{}
		case NetworksSelector:
			recv = []types.Network{}
		case RoutesSelector:
			recv = []types.Route{}
		case StopAreasSelector:
			recv = []types.StopArea{}
		case StopPointsSelector:
			recv = []types.StopPoint{}
		case PhysicalModesSelector:
			recv = []types.PhysicalMode{}
		case CompaniesSelector:
			recv = []types.Company{}
		/*case VehicleJourneysSelector:
		recv = []types.VehicleJourneys*/
		case DisruptionsSelector:
			recv = []types.Disruption{}
		}

		// If we have found something, let's break
		if recv != nil {
			second = first[k]
			break
		}
	}

	// If we have found nothing, return an error
	if recv == nil {
		return errors.New("error: no known key in response")
	}

	// Else, let's unmarshal
	err = json.Unmarshal(second, &recv)
	if err != nil {
		return errors.Wrap(err, "error while unmarshalling")
	}

	// Now assign it to sasr.PTObjects
	sasr.PTObjects = recv

	return nil
}

// ExploreRequest is the query you need to build before passing it to Explore function
type ExploreRequest struct {
	Query string // The search item

	// Depth can expand the data by making it more verbose.
	// Acceptable values are 0 (light), 1 (regular), 2 (rich), 3 (verbose)
	Depth uint8

	// Enables GeoJSON data in the reply. GeoJSON objects can be VERY large ! >1MB.
	Geo bool

	// Maximum amount of results
	Count uint

	// Request for specific pickup line. It refers to the odt section.
	// Warning: Only works with Line request
	//
	// It can take one of these values:
	// 	"all": no filter, provide all public transport lines, whatever its type
	// 	"scheduled": provide only regular lines
	// 	"with_stops": to get regular, “odt_with_stop_time” and “odt_with_stop_point” lines.
	// 	"zonal"" : to get “odt_with_zone” lines with non-detailed journeys
	ODTLevel string // NOT IMPLEMENTED

	// If coordinates are specified in the filtering, does a proximity search with given radius.
	// Default value: 200meters.
	Radius uint // NOT IMPLEMENTED

	// Since and Until are used for Vehicle Journeys & Disruptions to filter for a period.
	// TODO: Implement them.
	Since time.Time
	Until time.Time
}

// toURL formats a ExploreRequest request to url.Values
func (req ExploreRequest) toURL() (url.Values, error) {
	params := url.Values{
		"depth":           []string{strconv.FormatUint(uint64(req.Depth), 10)},
		"disable_geojson": []string{strconv.FormatBool(!req.Geo)},
	}

	if req.Count != 0 {
		countStr := strconv.FormatUint(uint64(req.Count), 10)
		params["count"] = []string{countStr}
	}

	return params, nil
}

// selectAndSearch is the internal function used by Explore functions
func (s *Session) explore(ctx context.Context, url string, params ExploreRequest) (*ExploreResults, error) {
	var results = &ExploreResults{}
	err := s.request(ctx, url, params, results)

	return results, err
}

// XXXSelector are used in PTObjectsSelect to select a specific category of public transportation objects to be selected
const (
	NetworksSelector        = "networks"
	LinesSelector           = "lines"
	RoutesSelector          = "routes"
	StopPointsSelector      = "stop_points"
	StopAreasSelector       = "stop_areas"
	CommercialModesSelector = "commercial_modes"
	PhysicalModesSelector   = "physical_modes"
	CompaniesSelector       = "companies"
	VehicleJourneysSelector = "vehicle_journeys"
	DisruptionsSelector     = "disruptions"
)

// ExploreRegion searches in all elements of the given selector (lines, networks, etc.) within a coverage, returning a list of ptObjects of the specific type.
func (scope *Scope) ExploreRegion(ctx context.Context, selector string, opts ExploreRequest) (*ExploreResults, error) {
	// Create the URL
	url := scope.session.apiURL + "/coverage/" + string(scope.region) + "/" + selector

	// Call
	return scope.session.explore(ctx, url, opts)
}

// Explore searches in all elements of the given selector (lines, networks, etc.) within the region covering the given coordinates, returning a list of ptObjects of the specific type.
func (s *Session) Explore(ctx context.Context, selector string, lat, lng float64, opts ExploreRequest) (*ExploreResults, error) {
	// Create the URL
	coords := fmt.Sprintf("%3.3f;%3.3f", lng, lat)
	url := s.apiURL + "/" + coords + "/" + selector

	// Call
	return s.explore(ctx, url, opts)
}
