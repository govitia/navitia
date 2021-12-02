package types

import (
	"encoding/json"
	"fmt"
	"image/color"
	"strconv"

	"github.com/pkg/errors"
)

// A Line codes for a public transit line.
// Warning: a Line isn't a route, it has no direction information, and can have several embranchments.
// See http://doc.navitia.io/#public-transport-objects.
type Line struct {
	ID    ID          `json:"id"`    // ID is the navitia identifier of the line, eg: "line:RAT:M6"
	Name  string      `json:"name"`  // Name of the line eg: "Nation - Charles de Gaule Etoile"
	Code  string      `json:"code"`  // Code is the codename of the line
	Color color.Color `json:"color"` // Color of the Line, eg "FFFFFF"

	// OpeningTime is the opening time of the line
	OpeningTime struct {
		Hours   uint8 `json:"hours"`
		Minutes uint8 `json:"minutes"`
		Seconds uint8 `json:"seconds"`
	} `json:"opening_time"`

	// ClosingTime is the closing time of the line
	ClosingTime struct {
		Hours   uint8 `json:"hours"`
		Minutes uint8 `json:"minutes"`
		Seconds uint8 `json:"seconds"`
	} `json:"closing_time"`

	Routes         []Route        `json:"routes"`          // Routes contains the routes of the line
	CommercialMode CommercialMode `json:"commercial_mode"` // CommercialMode of the line
	PhysicalModes  []PhysicalMode `json:"physical_modes"`  // PhysicalModes of the line
}

// jsonLine define the JSON implementation of Line types.
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonLine struct {
	ID             *ID             `json:"id"`              // ID is the navitia identifier of the line, eg: "line:RAT:M6"
	Name           *string         `json:"name"`            // Name of the line eg: "Nation - Charles de Gaule Etoile"
	Code           *string         `json:"code"`            // Code is the codename of the line
	Routes         *[]Route        `json:"routes"`          // Routes contains the routes of the line
	CommercialMode *CommercialMode `json:"commercial_mode"` // CommercialMode of the line
	PhysicalModes  *[]PhysicalMode `json:"physical_modes"`  // PhysicalModes of the line

	// Value to process
	Color       string `json:"color"`        // Color of the Line, eg "FFFFFF"
	OpeningTime string `json:"opening_time"` // OpeningTime is the opening time of the line
	ClosingTime string `json:"closing_time"` // ClosingTime is the closing time of the line
}

// UnmarshalJSON implements json.Unmarshaller for a Line
func (l *Line) UnmarshalJSON(b []byte) error {
	data := jsonLine{
		ID:             &l.ID,
		Name:           &l.Name,
		Code:           &l.Code,
		Routes:         &l.Routes,
		CommercialMode: &l.CommercialMode,
		PhysicalModes:  &l.PhysicalModes,
	}

	if err := json.Unmarshal(b, &data); err != nil {
		return fmt.Errorf("error while unmarshalling Line types : %w", err)
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"Line", b}

	// Now process the values

	// For Color: we expect a color string length of 6 because it should be coded in hexadecimal
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return gen.err(err, "Color", "color", str, "error in parseColor")
		}
		l.Color = clr
	}

	// For OpeningTime and ClosingTime: we define a function to help us
	parseTime := func(str string) (h, m, s uint8, err error) {
		if len(str) != 6 {
			err = errors.Errorf("time string not to standard: len=%d instead of 6", len(str))
			return
		}
		h64, err := strconv.ParseUint(str[:2], 10, 8)
		if err != nil {
			err = errors.Wrap(err, "error while parsing hours")
			return
		}
		m64, err := strconv.ParseUint(str[2:4], 10, 8)
		if err != nil {
			err = errors.Wrap(err, "error while parsing minutes")
			return
		}
		s64, err := strconv.ParseUint(str[4:], 10, 8)
		if err != nil {
			err = errors.Wrap(err, "error while parsing seconds")
			return
		}
		return uint8(h64), uint8(m64), uint8(s64), nil
	}
	// We expect as well a 6-character long value
	if str := data.OpeningTime; len(str) == 6 {
		t := &l.OpeningTime
		var err error
		t.Hours, t.Minutes, t.Seconds, err = parseTime(str)
		if err != nil {
			return gen.err(err, "OpeningTime", "opening_time", str, "error in parseTime")
		}
	}
	if str := data.ClosingTime; len(str) == 6 {
		t := &l.ClosingTime
		var err error
		t.Hours, t.Minutes, t.Seconds, err = parseTime(str)
		if err != nil {
			return gen.err(err, "ClosingTime", "closing_time", str, "error in parseTime")
		}
	}

	return nil
}
