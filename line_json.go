package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

// UnmarshalJSON implements json.Unmarshaller for a Line
func (l *Line) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Pointers to the corresponding real values
		ID             *ID             `json:"id"`
		Name           *string         `json:"name"`
		Code           *string         `json:"code"`
		Routes         *[]Route        `json:"routes"`
		CommercialMode *CommercialMode `json:"commercial_mode"`
		PhysicalModes  *[]PhysicalMode `json:"physical_modes"`

		// Value to process
		Color       string `json:"color"`
		OpeningTime string `json:"opening_time"`
		ClosingTime string `json:"closing_time"`
	}{
		ID:             &l.ID,
		Name:           &l.Name,
		Code:           &l.Code,
		Routes:         &l.Routes,
		CommercialMode: &l.CommercialMode,
		PhysicalModes:  &l.PhysicalModes,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Line.UnmarshalJSON: error while unmarshalling Line")
	}

	// Now process the values

	// For Color: we expect a color string length of 6 because it should be coded in hexadecimal
	if str := data.Color; len(str) == 6 {
		clr, err := parseColor(str)
		if err != nil {
			return errors.Wrapf(err, "Line.UnmarshalJSON: error while parsing color (given \"color\":\"%s\")", str)
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
		t.Hours, t.Minutes, t.Seconds, err = parseTime(str)
		if err != nil {
			return errors.Wrap(err, "Line.UnmarshalJSON: error while parsing OpeningTime")
		}
	}
	if str := data.ClosingTime; len(str) == 6 {
		t := &l.ClosingTime
		t.Hours, t.Minutes, t.Seconds, err = parseTime(str)
		if err != nil {
			return errors.Wrap(err, "Line.UnmarshalJSON: error while parsing OpeningTime")
		}
	}

	return nil
}
