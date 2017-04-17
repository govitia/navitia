// +build gofuzz

package types

import "encoding/json"

func FuzzPlaceContainer(data []byte) int {
	var pc = &PlaceContainer{}

	// Let's unmarshal, this is not our job so "bleh"
	err := json.Unmarshal(data, pc)
	if err != nil {
		return 0
	}

	// Let's test the .Place method !
	// No need to check before as .Check is called in .Place !
	place, err := pc.Place()
	if err != nil {
		return 0
	}

	// If we have an empty PlaceContainer but a non-nil place, panic !
	// But if we have both empty PlaceContainer and nil place, exit with 1, as this is the expected behaviour.
	if pc.IsEmpty() {
		if place != nil {
			panic("Error: empty PlaceContainer but non-nil place")
		}
		return 1
	}

	// Few methods to test
	_ = place.PlaceName()
	_ = place.PlaceID()
	_ = place.PlaceType()
	_ = place.String()

	return 1
}
