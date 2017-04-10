package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

// UnmarshalJSON implements json.Unmarshaller for a Section
func (s *Section) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		Type *SectionType `json:"type"`
		ID   *ID          `json:"id"`
		Mode *Mode        `json:"mode"`

		Departure string `json:"departure_date_time"`
		Arrival   string `json:"arrival_date_time"`

		From PlaceCountainer `json:"from"`
		To   PlaceCountainer `json:"to"`

		StopTimes *[]StopTime `json:"stop_date_times"`

		Duration int64

		// Information to display
		Display *DisplayInformations `json:"display_informations"`

		// Additional informations, from what I can see this is always a PTMethod
		Additional *[]PTMethod `json:"additional_informations"`
	}{
		Type:       &s.Type,
		ID:         &s.ID,
		Mode:       &s.Mode,
		Display:    &s.Display,
		Additional: &s.Additional,
		StopTimes:  &s.StopTimes,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Now process the two PlaceCountainer
	s.From, err = data.From.Place()
	if err != nil {
		return errors.Wrap(err, "Error while parsing places")
	}
	s.To, err = data.To.Place()
	if err != nil {
		return errors.Wrap(err, "Error while parsing places")
	}

	// For departure and arrival, we use parseDateTime
	s.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}
	s.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}

	// As the given duration is in second, let's multiply it by one second to have the correct value
	s.Duration = time.Duration(data.Duration) * time.Second

	return nil
}

// UnmarshalJSON implements json.Unmarshaller for a PTDateTime
func (ptdt *PTDateTime) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	data := &struct {
		Departure string `json:"departure_date_time"`
		Arrival   string `json:"arrival_date_time"`
	}{}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Now we use parseDateTime
	ptdt.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}
	ptdt.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}

	return nil
}
