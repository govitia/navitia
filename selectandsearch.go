package navitia

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

// SelectAndSearchResults doesn't have pagination, as the remote API doesn't support it.
//
// SelectAndSearchResults can be sorted, it implements sort.Interface.
type SelectAndSearchResults struct {
	PTObjects interface{}

	Logging `json:"-"`
}

// UnmarshalJSON implements unmarshalling for SelectAndSearchResults
func (sasr *SelectAndSearchResults) UnmarshalJSON(b []byte) error {
	first := make(map[string]json.RawMessage)

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
		return errors.New("error: no known receive types known")
	}

	// Else, let's unmarshal
	err := json.Unmarshal(second, recv)
	if err != nil {
		return errors.Wrap(err, "error while unmarshalling")
	}

	// Now assign it to sasr.PTObjects
	sasr.PTObjects = recv
	return nil
}

// SelectAndSearchRequest is the query you need to build before passing it to SelectAndSearch function
type SelectAndSearchRequest struct {
	Query string // The search item

	// Depth can expand the data by making it more verbose.
	// Acceptable values are 0 (light), 1 (regular), 2 (rich), 3 (verbose)
	Depth uint8

	// Enables GeoJSON data in the reply. GeoJSON objects can be VERY large ! >1MB.
	Geo bool

	// Maximum amount of results
	Count uint
}

// toURL formats a SelectAndSearchRequest request to url.Values
func (req SelectAndSearchRequest) toURL() (url.Values, error) {
	params := url.Values{
		"q":               []string{req.Query},
		"depth":           []string{strconv.FormatUint(uint64(req.Depth), 10)},
		"disable_geojson": []string{strconv.FormatBool(!req.Geo)},
	}

	if req.Count != 0 {
		countStr := strconv.FormatUint(uint64(req.Count), 10)
		params["count"] = []string{countStr}
	}

	return params, nil
}

// selectAndSearch is the internal function used by SelectAndSearch functions
func (s *Session) selectAndSearch(ctx context.Context, url string, params PTObjectsRequest) (*SelectAndSearchResults, error) {
	var results = &SelectAndSearchResults{}
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

// SelectAndSearch searches in all elements of the given selector (lines, networks, etc.) within a coverage using their names, returning a list of ptObjects of the specific type.
func (scope *Scope) SelectAndSearch(ctx context.Context, params PTObjectsRequest, selector string) (*SelectAndSearchResults, error) {
	// Create the URL
	url := scope.session.apiURL + "/coverage/" + string(scope.region) + "/" + selector

	// Call
	return scope.session.selectAndSearch(ctx, url, params)
}

// SelectAndSearchC searches in all elements of the given selector (lines, networks, etc.) within the region covering the given coordinates using their names, returning a list of ptObjects of the specific type.
func (s *Session) SelectAndSearchC(ctx context.Context, params PTObjectsRequest, selector string, lat, lng float64) (*SelectAndSearchResults, error) {
	// Create the URL
	coords := fmt.Sprintf("%3.3f;%3.3f", lng, lat)
	url := s.apiURL + "/" + coords + "/" + selector

	// Call
	return s.selectAndSearch(ctx, url, params)
}
