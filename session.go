package navitia

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"
	"sort"
	"time"

	"github.com/pkg/errors"

	"github.com/govitia/navitia/types"
)

const (
	// default Navitia REST service
	defaultAPIURL = "https://api.navitia.io/v1"

	// Maximum size of response in bytes
	// 10 megabytes
	maxSize int64 = 10e6
)

var defaultClient = &http.Client{}

// Session holds a current session, it is thread-safe
type Session struct {
	APIKey string
	APIURL string

	client  *http.Client
	created time.Time
}

// New creates a new session given an API Key.
// It acts as a convenience wrapper to NewCustom.
//
// Warning: No Timeout is indicated in the default http client, and as such, it is strongly advised to use NewCustom with a custom *http.Client !
func New(key string) (*Session, error) {
	return NewCustom(key, path.Clean(defaultAPIURL), defaultClient)
}

// NewCustom creates a custom new session given an API key, URL to api base & http client
func NewCustom(key, url string, client *http.Client) (*Session, error) {
	return &Session{
		APIKey:  key,
		APIURL:  url,
		created: time.Now(),
		client:  client,
	}, nil
}

// departures is the internal function used by Departures & Arrivals functions
func (s *Session) connections(ctx context.Context, url string, req ConnectionsRequest) (*ConnectionsResults, error) {
	results := &ConnectionsResults{}
	err := s.request(ctx, url, req, results)
	return results, err
}

// departures is the internal function used by Journeys functions
func (s *Session) departures(ctx context.Context, url string, req DeparturesRequest) (*DeparturesResults, error) {
	results := &DeparturesResults{session: s}
	err := s.request(ctx, url, req, results)
	return results, err
}

// Departures computes a list of Departures according to the parameters given
func (s *Session) Departures(ctx context.Context, req DeparturesRequest) (*DeparturesResults, error) {
	// Create the URL
	reqURL := s.APIURL + "/" + departuresEndpoint

	return s.departures(ctx, reqURL, req)
}

// DeparturesC requests the departures from a point described by coordinates.
func (s *Session) DeparturesC(ctx context.Context, req ConnectionsRequest, coords types.Coord) (*ConnectionsResults, error) {
	// Create the URL
	coordsQ := string(coords.ID())
	scopeURL := s.APIURL + "/coverage/" + coordsQ + "/coords/" + coordsQ + "/" + departuresEndpoint

	return s.connections(ctx, scopeURL, req)
}

// journeys is the internal function used by Journeys functions
func (s *Session) journeys(ctx context.Context, url string, req JourneyRequest) (*JourneyResults, error) {
	results := &JourneyResults{session: s}
	err := s.request(ctx, url, req, results)
	return results, err
}

// Journeys computes a list of journeys according to the parameters given
func (s *Session) Journeys(ctx context.Context, req JourneyRequest) (*JourneyResults, error) {
	// Create the URL
	reqURL := s.APIURL + "/" + journeysEndpoint

	// Call
	return s.journeys(ctx, reqURL, req)
}

// places is the internal function used by Places functions
func (s *Session) places(ctx context.Context, url string, params PlacesRequest) (*PlacesResults, error) {
	results := &PlacesResults{session: s}
	err := s.request(ctx, url, params, results)

	// Sort the places if quality is defined on the results, no need to expand some call
	// Justification for the if condition: If at least of of the results quality is 0, then all of them are 0.
	if results.Len() != 0 && results.Places[0].Quality != 0 {
		sort.Sort(sort.Reverse(results))
	}
	return results, err
}

// Places searches in all geographical objects using their names, returning a list of corresponding places.
// It is context aware.
func (s *Session) Places(ctx context.Context, params PlacesRequest) (*PlacesResults, error) {
	// Create the URL
	reqURL := s.APIURL + "/" + placesEndpoint

	// Call
	return s.places(ctx, reqURL, params)
}

func (s *Session) region(ctx context.Context, url string, params RegionRequest) (*RegionResults, error) {
	results := &RegionResults{session: s}
	err := s.request(ctx, url, params, results)
	return results, err
}

// Regions lists the areas covered by the Navitia API.
// i.e it returns the coverage of the API.
// It is context aware.
func (s *Session) Regions(ctx context.Context, req RegionRequest) (*RegionResults, error) {
	// Create the URL
	reqURL := s.APIURL + "/" + regionEndpoint

	// Call and return
	return s.region(ctx, reqURL, req)
}

// RegionByID provides information about a specific region.
// If the ID provided isn't an ID of a region, this WILL fail.
// It is context aware.
func (s *Session) RegionByID(ctx context.Context, req RegionRequest, id types.ID) (*RegionResults, error) {
	// Build the URL
	reqURL := s.APIURL + "/" + regionEndpoint + "/" + string(id)

	// Call and return
	return s.region(ctx, reqURL, req)
}

// RegionByPos provides information about the region englobing the specific position.
// It is context aware.
func (s *Session) RegionByPos(ctx context.Context, req RegionRequest, coords types.Coord) (*RegionResults, error) {
	// Build the URL
	coordsQ := string(coords.ID())
	reqURL := s.APIURL + "/" + regionEndpoint + "/" + coordsQ

	// Call and return
	return s.region(ctx, reqURL, req)
}

// requestURL requests a url, with the query already encoded in, and decodes the result in res.
func (s *Session) requestURL(ctx context.Context, url string, res results) error {
	// Store creation time
	res.creating()

	// Create the request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrapf(err, "couldn't create new request (for %s)", url)
	}

	// Add basic auth
	req.SetBasicAuth(s.APIKey, "")

	// Execute the request
	resp, err := s.client.Do(req)
	res.sending()

	// Check the response
	if err != nil {
		return errors.Wrap(err, "error while executing request")
	}
	if resp.StatusCode != http.StatusOK {
		return parseRemoteError(resp)
	}

	// Defer the close
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	// Check for cancellation
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Limit the reader
	reader := io.LimitReader(resp.Body, maxSize)

	// Parse the now limited body
	dec := json.NewDecoder(reader)
	err = dec.Decode(res)
	if err != nil {
		return errors.Wrap(err, "JSON decoding failed")
	}
	res.parsing()

	return err
}

// request does a request given a url, query and results to populate
func (s *Session) request(ctx context.Context, baseURL string, query query, res results) error {
	// Encode the parameters
	values, err := query.toURL()
	if err != nil {
		return errors.Wrap(err, "error while retrieving url values to be encoded")
	}
	reqURL := baseURL + "?" + values.Encode()

	// Call requestURL
	return s.requestURL(ctx, reqURL, res)
}

// Scope creates a coverage-scoped session given a region ID.
func (s *Session) Scope(region types.ID) *Scope {
	return &Scope{region: region, session: s}
}

// vehicleJourneys is the internal function used by VehicleJourneys functions
func (s *Session) vehicleJourneys(ctx context.Context, url string, req VehicleJourneyRequest) (*VehicleJourneyResults, error) {
	results := &VehicleJourneyResults{session: s}
	err := s.request(ctx, url, req, results)
	return results, err
}

const vehicleJourneysEndpoint string = "vehicle_journeys"

// VehicleJourneys computes a list of VehicleJourneys according to the parameters given
func (s *Session) VehicleJourneys(ctx context.Context, req VehicleJourneyRequest) (*VehicleJourneyResults, error) {
	// Create the URL
	reqURL := s.APIURL + "/" + vehicleJourneysEndpoint

	return s.vehicleJourneys(ctx, reqURL, req)
}
