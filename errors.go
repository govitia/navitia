package navitia

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// RemoteErrorID is an ID for a remote error
type RemoteErrorID string

// RemoteErrXXX are the known error ids that can be returned by the navitia server
const (
	// 404 Errors

	RemoteErrDateOutOfBounds       RemoteErrorID = "date_out_of_bounds"         // When the given date is out of bounds of the production dates of the region
	RemoteErrNoOrigin              RemoteErrorID = "no_origin"                  // Couldn’t find an origin for the journeys
	RemoteErrNoDestination         RemoteErrorID = "no_destination"             // Couldn’t find an destination for the journeys
	RemoteErrNoOriginNoDestination RemoteErrorID = "nor_origin_nor_destination" // Couldn’t find an origin nor a destination for the journeys
	RemoteErrUnknownObject         RemoteErrorID = "unknown_object"             // Unknown Object

	// 400 Errors

	RemoteErrBadFilter     RemoteErrorID = "bad_filter"      // Bad filter (with custom filter)
	RemoteErrUnableToParse RemoteErrorID = "unable_to_parse" // Unable to parse mal-formed custom filter"
)

// remoteErrorsDescriptions contains human-readable descriptions for a given remote error ID
//
// Can also be used as a list of known error IDs
var remoteErrorsDescriptions = map[RemoteErrorID]string{
	RemoteErrDateOutOfBounds:       "When the given date is out of bounds of the production dates of the region",
	RemoteErrNoOrigin:              "Couldn’t find an origin for the journeys",
	RemoteErrNoDestination:         "Couldn’t find an destination for the journeys",
	RemoteErrNoOriginNoDestination: "Couldn’t find an origin nor a destination for the journeys",
	RemoteErrUnknownObject:         "Unknown Object",
	RemoteErrBadFilter:             "Bad filter (with custom filter)",
	RemoteErrUnableToParse:         "Unable to parse mal-formed custom filter",
}

// A RemoteError represents an error sent by the server
type RemoteError struct {
	StatusCode int
	ID         RemoteErrorID `json:"id"`
	Message    string        `json:"message"`
}

// Error formats the error in a human-readable format
// Also allows it to satisfy the error interface
func (err RemoteError) Error() string {
	var s string

	// If this is a 40x error then use our information about errors
	if err.StatusCode == 404 || err.StatusCode == 400 {
		s = fmt.Sprintf("request error (id: %s):", err.ID)
		if desc, ok := remoteErrorsDescriptions[err.ID]; ok {
			s += fmt.Sprintf(" %s:", desc)
		}
		s += err.Message
	} else {
		s = fmt.Sprintf("remote failure (id: %s): %s", err.ID, err.Message)
	}

	return s
}

// parseRemoteError parses a non 200 OK status-coded response and returns the error
func parseRemoteError(resp *http.Response) error {
	var remoteErr = &RemoteError{StatusCode: resp.StatusCode}

	// Parse it
	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&remoteErr)
	if err != nil {
		return errors.Wrap(err, "parseRemoteError: error while decoding JSON")
	}

	// Return
	return remoteErr
}
