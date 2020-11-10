package navitia

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/pkg/errors"

	"github.com/govitia/navitia/types"
	"github.com/govitia/navitia/utils"
)

// A Connection is either a Departure or an Arrival
type Connection struct {
	Display   types.Display
	StopPoint types.StopPoint
	Route     types.Route
	// StopDateTime
}

// ConnectionsResults holds the results of a departures or arrivals request.
type ConnectionsResults struct {
	Connections []Connection
	Paging      Paging `json:"links"`
	Logging     `json:"-"`
}

// UnmarshalJSON implements unmarshalling for ConnectionsResults.
func (cr *ConnectionsResults) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Pointers to the corresponding real values
		Paging *Paging `json:"links"`

		// Value to process
		Departures *[]Connection `json:"departures"`
		Arrivals   *[]Connection `json:"arrivals"`
	}{
		Paging: &cr.Paging,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "ConnectionsResults.UnmarshalJSON: error while unmarshalling Line")
	}

	// Now process the values
	switch {
	case data.Departures != nil:
		cr.Connections = *data.Departures
	case data.Arrivals != nil:
		cr.Connections = *data.Arrivals
	}
	// else there's nor Departures nor Arrivals found

	return nil
}

// ConnectionsRequest contains the optional parameters for a Departures request.
type ConnectionsRequest struct {
	// From what time on do you want to see the results ?
	From time.Time

	// Maximum duration between From and the retrieved results (default 24h)
	Duration time.Duration

	// The maximum amount of results (default 10)
	Count uint

	// ForbiddenURIs
	Forbidden []types.ID

	// Freshness of the data
	Freshness types.DataFreshness

	// Enables GeoJSON data in the reply. GeoJSON objects can be VERY large ! >1MB.
	Geo bool
}

func (req ConnectionsRequest) toURL() (url.Values, error) {
	rb := utils.NewRequestBuilder()

	rb.AddDateTime("datetime", req.From)

	// If count is defined don't bother with the minimimal and maximum amount of items to return
	if req.Count != 0 {
		rb.AddUInt("count", req.Count)
	}

	// Deal with the forbidden URIs
	rb.AddIDSlice("forbidden_uris[]", req.Forbidden)

	// Set the freshness
	rb.AddString("data_freshness", string(req.Freshness))

	// Add GEO
	if !req.Geo {
		rb.AddString("disable_geojson", "true")
	}

	return rb.Values(), nil
}

const (
	departuresEndpoint string = "departures"
	arrivalsEndpoint          = "arrivals"
)
