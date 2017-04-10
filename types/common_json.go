package types

import (
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
