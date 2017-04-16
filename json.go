// json.go provides types & functions for json unmarshalling

package types

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const (
	// DateTimeFormat is the format used by the Navitia Api for use with time pkg.
	DateTimeFormat string = "20060102T150405" // YYYYMMDDThhmmss
	// DateFormat is when there is no time info
	DateFormat string = "20060102"
)

// parseDateTime parses a time formatted under iso-date-time as indicated in the Navitia api.
// This is simply parsing a date formatted under the standard ISO 8601.
// If the given string is empty (i.e ""), then the zero value of time.Time will be returned
func parseDateTime(datetime string) (time.Time, error) {
	// If there's no datetime given, just return the zero value
	if datetime == "" || datetime == "not-a-date-time" {
		return time.Time{}, nil
	}

	// If the datetime doesn't countain a "T", then it does not have time info
	var format string
	if strings.Contains(datetime, "T") {
		format = DateTimeFormat
	} else {
		format = DateFormat
	}

	// Parse it
	res, err := time.Parse(format, datetime)
	if err != nil {
		err = errors.Wrap(err, "parseDateTime: error while parsing datetime")
	}
	return res, err
}

// UnmarshalError is returned when unmarshalling fails
// It implements both error and github.com/pkg/errors's causer
type UnmarshalError struct {
	// Type on which the unmarshaller where the error occured works
	Type string

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
	msg := fmt.Sprintf("(*%s).UnmarshalJSON: Unmarshalling %s (json: \"%s\") with value \"%v\" failed", err.Type, err.Name, err.Key, err.Value)
	if err.Message != "" {
		msg += ": " + err.Message
	}
	if err.Underlying != nil {
		msg += " [" + err.Cause().Error() + "]"
	}
	return msg
}

// unmarshalErrorer allows us to make better error messages
type unmarshalErrorMaker struct {
	Type string
}

// err creates a new UnmarshalError
func (gen unmarshalErrorMaker) err(underlyingErr error, name string, key string, value interface{}, message string) error {
	return UnmarshalError{
		Type:       gen.Type,
		Key:        key,
		Name:       name,
		Value:      value,
		Message:    message,
		Underlying: underlyingErr,
	}
}

// unmarshalErr creates a new UnmarshalError
func unmarshalErr(underlyingErr error, type_ string, name string, key string, value interface{}, message string) error {
	return UnmarshalError{
		Type:       type_,
		Key:        key,
		Name:       name,
		Value:      value,
		Message:    message,
		Underlying: underlyingErr,
	}
}
