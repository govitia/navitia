package navitia

import (
	"golang.org/x/net/context"

	"github.com/govitia/navitia/types"
)

// A Scope is a coverage-scoped question, allowing you to query information about a specific region.
//
// It is needed for every non-global request you wish to make, and helps have better results with some global request too!
type Scope struct {
	region  types.ID
	session *Session
}

// ArrivalsSA requests the arrivals for a given StopArea in a given region.
func (scope *Scope) ArrivalsSA(ctx context.Context, req ConnectionsRequest, resource types.ID) (*ConnectionsResults, error) {
	// Create the URL
	scopeURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/stop_areas/" + string(resource) + "/" + arrivalsEndpoint

	return scope.session.connections(ctx, scopeURL, req)
}

// ArrivalsSP requests the arrivals for a given StopPoint in a given region.
func (scope *Scope) ArrivalsSP(ctx context.Context, req ConnectionsRequest, resource types.ID) (*ConnectionsResults, error) {
	// Create the URL
	scopeURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/stop_points/" + string(resource) + "/" + arrivalsEndpoint

	return scope.session.connections(ctx, scopeURL, req)
}

// ArrivalsC requests the arrivals from a point described by coordinates.
func (s *Session) ArrivalsC(ctx context.Context, req ConnectionsRequest, coords types.Coord) (*ConnectionsResults, error) {
	// Create the URL
	coordsQ := string(coords.ID())
	scopeURL := s.APIURL + "/coverage/" + coordsQ + "/coords/" + coordsQ + "/" + arrivalsEndpoint

	return s.connections(ctx, scopeURL, req)
}

// Departures computes a list of Departures according to the parameters given in a specific scope
func (scope *Scope) Departures(ctx context.Context, req DeparturesRequest) (*DeparturesResults, error) {
	// there is a special case for departures stop areas, it needs to be added before any parameters
	filterByVJ := ""
	if req.StopArea != "" {
		filterByVJ = "stop_areas/" + req.StopArea
	}

	// Create the URL
	reqURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/" + filterByVJ + "/" + departuresEndpoint

	return scope.session.departures(ctx, reqURL, req)
}

// DeparturesSA requests the departures for a given StopArea
func (scope *Scope) DeparturesSA(ctx context.Context, req ConnectionsRequest, resource types.ID) (*ConnectionsResults, error) {
	// Create the URL
	scopeURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/stop_areas/" + string(resource) + "/" + departuresEndpoint

	return scope.session.connections(ctx, scopeURL, req)
}

// DeparturesSP requests the departures for a given StopPoint
func (scope *Scope) DeparturesSP(ctx context.Context, req ConnectionsRequest, resource types.ID) (*ConnectionsResults, error) {
	// Create the URL
	scopeURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/stop_points/" + string(resource) + "/" + departuresEndpoint

	return scope.session.connections(ctx, scopeURL, req)
}

// Journeys computes a list of journeys according to the parameters given in a specific scope
func (scope *Scope) Journeys(ctx context.Context, req JourneyRequest) (*JourneyResults, error) {
	// Create the URL
	reqURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/" + journeysEndpoint

	// Call
	return scope.session.journeys(ctx, reqURL, req)
}

// Places searches in all geographical objects within a coverage using their names, returning a list of places.
// It is context aware.
func (scope *Scope) Places(ctx context.Context, params PlacesRequest) (*PlacesResults, error) {
	// Create the URL
	reqURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/" + placesEndpoint

	// Call
	return scope.session.places(ctx, reqURL, params)
}

// VehicleJourneys computes a list of VehicleJourneys according to the parameters given in a specific scope
func (scope *Scope) VehicleJourneys(ctx context.Context, req VehicleJourneyRequest) (*VehicleJourneyResults, error) {
	// there is a special case for vehicle journey ID, it needs to be added before any parameters
	filterByVJ := ""
	if req.ID != "" {
		filterByVJ = "/" + string(req.ID)
	}

	// Create the URL
	reqURL := scope.session.APIURL + "/coverage/" + string(scope.region) + "/" + vehicleJourneysEndpoint + filterByVJ

	return scope.session.vehicleJourneys(ctx, reqURL, req)
}
