package types

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for a Region
func (r *Region) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	data := &struct {
		ID     *ID     `json:"id"`
		Name   *string `json:"name"`
		Status *string `json:"status"`

		DatasetCreation string `json:"dataset_created_at"`
		LastLoaded      string `json:"last_load_at"`

		ProductionStart string `json:"start_production_date"`
		ProductionEnd   string `json:"end_production_date"`

		Error *string `json:"error"`
	}{
		ID:     &r.ID,
		Name:   &r.Name,
		Status: &r.Status,
		Error:  &r.Error,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Now we use parseDateTime
	r.DatasetCreation, err = parseDateTime(data.DatasetCreation)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}
	r.LastLoaded, err = parseDateTime(data.LastLoaded)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}
	r.ProductionStart, err = parseDateTime(data.ProductionStart)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}
	r.ProductionEnd, err = parseDateTime(data.ProductionEnd)
	if err != nil {
		return errors.Wrap(err, "Error while parsing datetime")
	}

	return nil

}
