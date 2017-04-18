package types

import (
	"encoding/json"
	"strconv"

	"github.com/pkg/errors"
)

// UnmarshalJSON implements json.Unmarshaller for a Coordinates
func (c *Coordinates) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	data := &struct {
		Latitude  string `json:"lat"`
		Longitude string `json:"lon"`
	}{}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Now parse the values
	c.Latitude, err = strconv.ParseFloat(data.Latitude, 64)
	if err != nil {
		return errors.Wrap(err, "Error while parsing float in coordinates")
	}
	c.Longitude, err = strconv.ParseFloat(data.Longitude, 64)
	if err != nil {
		return errors.Wrap(err, "Error while parsing float in coordinates")
	}

	return nil
}
