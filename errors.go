package gonavitia

import "fmt"

type RemoteErrorID string

const (
	// 404 Errors

	RemoteErrDateOutOfBounds       RemoteErrorID = "date_out_of_bounds"
	RemoteErrNoOrigin                            = "no_origin"
	RemoteErrNoDestination                       = "no_destination"
	RemoteErrNoOriginNoDestination               = "nor_origin_nor_destination"
	RemoteErrUnknownObject                       = "unknown_object"

	// 400 Errors

	RemoteErrBadFilter     = "bad_filter"
	RemoteErrUnableToParse = "unable_to_parse"
)

// Human-readable descriptions for a given remote error ID
// Can also be used as a list of known error IDs
var RemoteErrorsDescriptions = map[RemoteErrorID]string{
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
	ID         RemoteErrorID `json:"error"`
	Message    string        `json:"message"`
}

// String formats the error in a human-readable format
func (err RemoteError) String() string {

	var s string

	// If this is a 40x error then use our information about errors
	if err.StatusCode == 404 || err.StatusCode == 400 {
		s = fmt.Sprintf("request error (id: %s):", err.ID)
		if desc, ok := RemoteErrorsDescriptions[err.ID]; ok {
			s += fmt.Sprintf(" %s:", desc)
		}
		s += err.Message
	} else {
		s = fmt.Sprintf("remote failure (id: %s): %s", err.ID, err.Message)
	}

	return s
}
