/* Package unmarshal provides types & functions for json unmarshalling from the navitia API*/
package unmarshal

import (
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	// DateTimeFormat is the format used by the Navitia Api for use with time pkg.
	DateTimeFormat string = "20060102T150405" // YYYYMMDDThhmmss
	// DateFormat is when there is no time info
	DateFormat string = "20060102"
)

// ParseDateTime parses a time formatted under iso-date-time as indicated in the Navitia api.
// This is simply parsing a date formatted under the standard ISO 8601.
// If the given string is empty (i.e ""), then the zero value of time.Time will be returned
func ParseDateTime(datetime string) (time.Time, error) {
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
		err = errors.Wrap(err, "ParseDateTime: error while parsing datetime")
	}
	return res, err
}
