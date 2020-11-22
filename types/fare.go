package types

import (
	"encoding/json"

	"github.com/pkg/errors"
	"golang.org/x/text/currency"
)

// Fare is the fare of some thing
type Fare struct {
	Total currency.Amount
	Found bool
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
	if err := json.Unmarshal(b, data); err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Fare", b}

	// Let's convert the cost now
	// If we have no defined fare, let's skip that part
	if data.Cost.Value == "" || data.Cost.Currency == "" {
		return nil
	}

	// First get the currency unit
	unit, err := currency.ParseISO(data.Cost.Currency)
	if err != nil {
		return gen.err(err, "Total", "cost.currency", data.Cost.Currency, "error while retrieving currency unit via currency.ParseISO")
	}

	// Now let's create the correct amount
	f.Total = unit.Amount(data.Cost.Value)

	return nil
}
