package navitia

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/aabizri/navitia/internal/unmarshal"
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
	Objects interface{}

	Paging Paging

	Logging
}

// ExploreRequest is the query you need to build before passing it to Explore function
type ExploreRequest struct {
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
	// 	"all": no filter, provide all public transport lines, whatever its type (Default)
	// 	"scheduled": provide only regular lines
	// 	"with_stops": to get regular, “odt_with_stop_time” and “odt_with_stop_point” lines.
	// 	"zonal"" : to get “odt_with_zone” lines with non-detailed journeys
	ODTLevel string

	// If coordinates are specified in the filtering, does a proximity search with given radius.
	// Default value: 200meters.
	Radius uint

	// Since and Until are used for Vehicle Journeys & Disruptions to filter for a period.
	// 	-On vehicle_journey this filter is applied using only the first stop time.
	// 	-On disruption this filter must intersect with one application period. “since” is included and “until” is excluded.
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

	if ol := req.ODTLevel; ol != "" {
		if ol != "all" && ol != "scheduled" && ol != "with_stops" && ol != "zonal" {
			return params, errors.Errorf("toURL: ODTLevel unknown (%s)", ol)
		}
		params["odt_level"] = []string{ol}
	}

	if req.Radius != 0 {
		radiusStr := strconv.FormatUint(uint64(req.Radius), 10)
		params["distance"] = []string{radiusStr}
	}

	// Deal with since and/or until, checking that the interval is correct
	if since := req.Since; !since.IsZero() {
		sinceStr := since.Format(unmarshal.DateTimeFormat)
		params["since"] = []string{sinceStr}
	}
	if until := req.Until; !until.IsZero() {
		if until.After(req.Since) {
			return params, errors.New("toURL: until date is after before date")
		}
		untilStr := until.Format(unmarshal.DateTimeFormat)
		params["until"] = []string{untilStr}
	}
	return params, nil
}

// explore is the internal function used by Explore functions
func (s *Session) explore(ctx context.Context, url string, params ExploreRequest) (*ExploreResults, error) {
	var results = &ExploreResults{}
	err := s.request(ctx, url, params, results)

	return results, errors.Wrapf(err, "error in call to explore for url %s", url)
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

// resourceTypeToSelector allows us to convert from a resource type to a selector
var resourceTypeToSelector = map[string]string{
	types.EmbeddedNetwork:        NetworksSelector,
	types.EmbeddedLine:           LinesSelector,
	types.EmbeddedRoute:          RoutesSelector,
	types.EmbeddedStopPoint:      StopPointsSelector,
	types.EmbeddedStopArea:       StopAreasSelector,
	types.EmbeddedCommercialMode: CommercialModesSelector,
	types.EmbeddedPhysicalMode:   PhysicalModesSelector,
	types.EmbeddedCompany:        CompaniesSelector,
	types.EmbeddedVehicleJourney: VehicleJourneysSelector,
	types.EmbeddedDisruption:     DisruptionsSelector,
}

// Explore searches in all elements of the given selector (lines, networks, etc.) within a scope, returning a list of ptObjects of the specific type.
func (scope *Scope) Explore(ctx context.Context, selector string, opts ExploreRequest) (*ExploreResults, error) {
	// Create the URL
	url := scope.baseURL + "/" + selector

	// Call
	return scope.session.explore(ctx, url, opts)
}

// ExploreResource searches in all elements of the given selector (lines, networks, etc.) linked to a resource inside a scope, returning a list of ptObjects
// of the specific type.
func (scope *Scope) ExploreResource(ctx context.Context, resID types.ID, selector string, opts ExploreRequest) (*ExploreResults, error) {
	// Extract the type
	resType, err := resID.Type()
	if err != nil {
		return nil, errors.Wrapf(err, "ExploreResource: couldn't extract type from resource ID \"%s\"", resID)
	}
	// Get the selector equivalent
	typeSelector, ok := resourceTypeToSelector[resType]
	if !ok {
		return nil, errors.Errorf("ExploreResource: couldn't find the typeSelector equivalent to the resource type identified (\"%s\")", resType)
	}

	// Create the URL
	url := scope.baseURL + "/" + typeSelector + "/" + string(resID) + "/" + selector

	// Call
	return scope.session.explore(ctx, url, opts)
}
