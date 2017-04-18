package types

import (
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// An ID is used throughout the library, it is something used by the navitia API and useful to communicate with it.
type ID string

// Check for ID validity
func (id ID) Check() error {
	if len(id) == 0 {
		return errors.Errorf("ID invalid: an empty string \"\" is not a valid ID")
	}
	return nil
}

// QueryEscape formats the given ID so that it can be safely used in a URL query
func (id ID) QueryEscape() string {
	if strings.Contains(string(id), ";") {
		return string(id)
	}
	return url.QueryEscape(string(id))
}

// typeNames stores navitia-side name of types that may appear in IDs
var typeNames = map[string]bool{
	"network":         true,
	"line":            true,
	"route":           true,
	"stop_area":       true,
	"commercial_mode": true,
	"physical_mode":   true,
	"company":         true,
	"admin":           true,
	"stop_point":      true,
}

// Type gets the type of object this ID refers to
// Possible types: network, line, route, stop_area, commercial_mode, physical_mode, company, admin, stop_point.
// Note that this doesn't always work. WIP
// If no type is found, type returns an empty string
func (id ID) Type() string {
	splitted := strings.Split(string(id), ":")
	if len(splitted) == 0 {
		return ""
	}

	possible := splitted[0]
	if typeNames[possible] {
		return possible
	}

	return ""
}
