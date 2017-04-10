package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/text/currency"
	"strconv"
	"time"
)

// UnmarshalJSON implements json.Unmarshaller for a Journey
func (j *Journey) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		Duration  int64 `json:"duration"`
		Transfers *uint `json:"nb_transfers"`

		Departure string `json:"departure_date_time"`
		Requested string `json:"requested_date_time"`
		Arrival   string `json:"arrival_date_time"`

		Sections *[]Section `json:"sections"`

		From PlaceCountainer `json:"from"`
		To   PlaceCountainer `json:"to"`

		Type *JourneyQualification `json:"type"`

		Fare *Fare `json:"fare"`

		Status *JourneyStatus `json:"status"`
	}{
		Transfers: &j.Transfers,
		Sections:  &j.Sections,
		Type:      &j.Type,
		Fare:      &j.Fare,
		Status:    &j.Status,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// As the given duration is in second, let's multiply it by one second to have the correct value
	j.Duration = time.Duration(data.Duration) * time.Second

	// For departure, requested and arrival, we use parseDateTime
	j.Departure, err = parseDateTime(data.Departure)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}
	j.Requested, err = parseDateTime(data.Requested)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}
	j.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}

	// For the places, we directly use the embedded type !
	j.From, err = data.From.Place()
	if err != nil {
		return errors.Wrap(err, "Error while parsing places")
	}
	j.To, err = data.To.Place()
	if err != nil {
		return errors.Wrap(err, "Error while parsing places")
	}

	return nil
}

// UnmarshalJSON implements json.Unmarshaller for a Fare
func (f *Fare) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		Found *bool `json:"found"`
		Cost  struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"cost"`
	}{
		Found: &f.Found,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Let's convert the cost now
	// If we have no defined fare, let's skip that part
	if data.Cost.Value == "" || data.Cost.Currency == "" {
		return nil
	}
	// First get the currency unit
	unit, err := currency.ParseISO(data.Cost.Currency)
	if err != nil {
		return errors.Wrap(err, "Error while dealing with currency unit !")
	}
	// Now let's create the correct amount
	f.Total = unit.Amount(data.Cost.Value)

	return nil
}

// UnmarshalJSON implements json.Unmarshaller for CO2Emissions
func (c *CO2Emissions) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		Unit  *string `json:"unit"`
		Value string  `json:"value"`
	}{
		Unit: &c.Unit,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Now parse the value
	f, err := strconv.ParseFloat(data.Value, 64)
	if err != nil {
		return errors.Wrap(err, "Error while parsing CO2 emissions value")
	}
	c.Value = f

	return nil
}
