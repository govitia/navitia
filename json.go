// json.go provides types & functions for json unmarshalling

package types

import (
	"fmt"
	"github.com/pkg/errors"
	"time"
)

// parseDateTime parses a time formatted under iso-date-time as indicated in the Navitia api.
// This is simply parsing a date formatted under the standard ISO 8601.
func parseDateTime(datetime string) (time.Time, error) {
	res, err := time.Parse(DateTimeFormat, datetime)
	if err != nil {
		err = errors.Wrap(err, "parseDateTime: error while parsing datetime")
	}
	return res, err
}

// UnmarshalError is returned when unmarshalling fails
// It implements both error and github.com/pkg/errors's causer
type UnmarshalError struct {
	// JSON Key where failure occured
	Key string

	// Name of the key in package
	Name string

	// Value associated with the key
	Value interface{}

	// Message of the error
	Message string

	// Underlying error
	Underlying error
}

// Cause implements github.com/pkg/error's causer
func (err UnmarshalError) Cause() error {
	return err.Underlying
}

// Error implements error
func (err UnmarshalError) Error() string {
	msg := fmt.Sprintf("Unmarshalling %s (json: \"%s\") with value \"%v\" failed", err.Name, err.Key, err.Value)
	if err.Message != "" {
		msg += ": " + err.Message
	}
	if err.Underlying != nil {
		msg += " [" + err.Cause().Error() + "]"
	}
	return msg
}

// unmarshalErr creates a new UnmarshalError
func unmarshalErr(underlyingErr error, name string, key string, value interface{}, message string) error {
	return UnmarshalError{
		Key:        key,
		Name:       name,
		Value:      value,
		Message:    message,
		Underlying: underlyingErr,
	}
}
