package gonavitia

import (
	"encoding/json"
	"github.com/aabizri/gonavitia/types"
	"github.com/pkg/errors"
)

func (res *PlacesResults) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	data := &struct {
		Places []types.PlaceCountainer `json:"places"`
	}{}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling PlacesResults")
	}

	// Now iterate through the places and populate res.Places
	res.Places = make([]types.Place, len(data.Places))
	for i, pc := range data.Places {
		place, err := pc.Place()
		if err != nil {
			return errors.Wrap(err, "Error while retrieving Place")
		}
		res.Places[i] = place
	}

	// Return
	return nil
}
