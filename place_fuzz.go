// +build gofuzz

package types

import "encoding/json"

func FuzzPlaceCountainer(data []byte) int {
	var pc = &PlaceCountainer{}

	// Let's unmarshal, this is not our job so "bleh"
	err := json.Unmarshal(data, pc)
	if err != nil {
		return 0
	}

	// Now that it is unmarshalled, let's test the Place method !
	place, err := pc.Place()
	if err != nil {
		return 0
	}

	place.PlaceName()
	place.PlaceID()
	place.PlaceType()

	return 1
}
