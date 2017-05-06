package types

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/aabizri/navitia/internal/unmarshal"
	"github.com/pkg/errors"
	"golang.org/x/text/currency"
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

		From *Container `json:"from"`
		To   *Container `json:"to"`

		Type *JourneyQualification `json:"type"`

		Fare *Fare `json:"fare"`

		Status *Effect `json:"status"`
	}{
		Transfers: &j.Transfers,
		Sections:  &j.Sections,
		From:      &j.From,
		To:        &j.To,
		Type:      &j.Type,
		Fare:      &j.Fare,
		Status:    &j.Status,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Let's create the error generator
	gen := unmarshal.NewGenerator("Journey", &b)
	defer gen.Close()

	// As the given duration is in second, let's multiply it by one second to have the correct value
	j.Duration = time.Duration(data.Duration) * time.Second

	// For departure, requested and arrival, we use unmarshal.ParseDateTime
	j.Departure, err = unmarshal.ParseDateTime(data.Departure)
	if err != nil {
		return gen.Gen(err, "Departure", "departure_date_time", data.Departure, "unmarshal.ParseDateTime failed")
	}
	j.Requested, err = unmarshal.ParseDateTime(data.Requested)
	if err != nil {
		return gen.Gen(err, "Requested", "requested_date_time", data.Requested, "unmarshal.ParseDateTime failed")
	}
	j.Arrival, err = unmarshal.ParseDateTime(data.Arrival)
	if err != nil {
		return gen.Gen(err, "Arrival", "arrival_date_time", data.Arrival, "unmarshal.ParseDateTime failed")
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

	// Let's create the error generator
	gen := unmarshal.NewGenerator("Fare", &b)
	defer gen.Close()

	// Let's convert the cost now
	// If we have no defined fare, let's skip that part
	if data.Cost.Value == "" || data.Cost.Currency == "" {
		return nil
	}

	// First get the currency unit
	unit, err := currency.ParseISO(data.Cost.Currency)
	if err != nil {
		return gen.Gen(err, "Total", "cost.currency", data.Cost.Currency, "error while retrieving currency unit via currency.ParseISO")
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

	// Let's create the error generator
	gen := unmarshal.NewGenerator("CO2Emissions", &b)
	defer gen.Close()

	// Now parse the value
	f, err := strconv.ParseFloat(data.Value, 64)
	if err != nil {
		return gen.Gen(err, "Value", "value", data.Value, "error in strconv.ParseFloat")
	}
	c.Value = f

	return nil
}
