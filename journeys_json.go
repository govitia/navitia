package types

import (
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/text/currency"
	"strconv"
	"time"
)

/*
UnmarshalJSON implements json.Unmarshaller for a Journey

Behaviour:
	- If "from" is empty, then don't populate the From field.
	- Same for "to"
*/
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

		From PlaceContainer `json:"from"`
		To   PlaceContainer `json:"to"`

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

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Journey"}

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
		return gen.err(err, "Departure", "departure_date_time", data.Departure, "parseDateTime failed")
	}
	j.Requested, err = parseDateTime(data.Requested)
	if err != nil {
		return gen.err(err, "Requested", "requested_date_time", data.Requested, "parseDateTime failed")
	}
	j.Arrival, err = parseDateTime(data.Arrival)
	if err != nil {
		return gen.err(err, "Arrival", "arrival_date_time", data.Arrival, "parseDateTime failed")
	}

	// For the places, we directly use the embedded type !
	// Warning: it is possible for a countainer to be empty, so we don't fill up the Place in those cases
	if !data.From.IsEmpty() {
		j.From, err = data.From.Place()
		if err != nil {
			return gen.err(err, "From", "from", data.From, " .Place() failed")
		}
	}
	if !data.To.IsEmpty() {
		j.To, err = data.To.Place()
		if err != nil {
			return gen.err(err, "To", "to", data.To, " .Place() failed")
		}
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
		return unmarshalErr(err, "Fare", "N/A", "cost.currency", data.Cost.Currency, "error while retrieving currency unit via currency.ParseISO")
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
		return errors.Wrap(err, "Error while unmarshalling CO2Emissions")
	}

	// Now parse the value
	f, err := strconv.ParseFloat(data.Value, 64)
	if err != nil {
		return unmarshalErr(err, "CO2Emissions", "Value", "value", data.Value, "error in strconv.ParseFloat")
	}
	c.Value = f

	return nil
}
