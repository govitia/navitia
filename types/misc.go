package types

import (
	"image/color"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// DataFreshness codes for a specific data freshness requirement: realtime or base_schedule
type DataFreshness string

// parseColor, given a hex code #RRGGBB returns a color.NRGBA
func parseColor(str string) (color.NRGBA, error) {
	var clr color.NRGBA

	if len(str) != 6 {
		return clr, errors.Errorf("parseColor: can't parse, invalid length (len=%d instead of 6)", len(str))
	}

	r, err := strconv.ParseUint(str[:2], 16, 8)
	if err != nil {
		return clr, errors.Wrapf(err, "parseColor: red component parsing failed (str: %s)", str[:2])
	}
	g, err := strconv.ParseUint(str[2:4], 16, 8)
	if err != nil {
		return clr, errors.Wrapf(err, "parseColor: green component parsing failed (str: %s)", str[2:4])
	}
	b, err := strconv.ParseUint(str[4:], 16, 8)
	if err != nil {
		return clr, errors.Wrapf(err, "parseColor: blue component parsing failed (str: %s)", str[4:])
	}

	return color.NRGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
	}, nil
}

const (
	// DataFreshnessRealTime means you'll get undisrupted journeys
	DataFreshnessRealTime DataFreshness = "realtime"
	// DataFreshnessBaseSchedule means you can get disrupted journeys in the response.
	DataFreshnessBaseSchedule = "base_schedule"
)

// A PTDateTime (pt stands for “public transport”) is a complex date time object to manage the difference between stop and leaving times at a stop.
// It is used by:
// 	- Row in Schedule
// 	- StopSchedule
// 	- StopDatetime
type PTDateTime struct {
	// Date/Time of departure
	Departure time.Time

	// Date/Time of arrival
	Arrival time.Time
}

// A Code is associated to a dataset
//
// Every object managed by Navitia comes with its own list of ids.
// You will find some source ids, merge ids, etc. in “codes” list in json responses.
// Be careful, these codes may not be unique. The navitia id is the only unique id.
type Code struct {
	Type  string
	Value string
}
