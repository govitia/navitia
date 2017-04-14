package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

var placeCountainers map[string]*PlaceCountainer

func loadPC() error {
	// Get the input
	data := testData["place"].known
	if len(data) == 0 {
		return nil
	}

	pcs := make(map[string]*PlaceCountainer, len(data))
	// For each of them, unmarshal and add to placeCountainers
	for name, datum := range data {
		var pc = PlaceCountainer{}

		err := json.Unmarshal(datum.raw, &pc)
		if err != nil {
			return fmt.Errorf("Error while unmarshalling: %v", err)
		}

		pcs[name] = &pc
	}

	placeCountainers = pcs

	return nil
}

// TestPlaceCountainer_Place_NoCompare tries to unmarshal all json test data for this type, and then call a few functions on it.
func TestPlaceCountainer_Place_NoCompare(t *testing.T) {
	// Get the input
	data := placeCountainers
	if len(data) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run it in parallel
	rgen := func(pc *PlaceCountainer) func(t *testing.T) {
		return func(t *testing.T) {
			place, err := pc.Place()
			if err != nil {
				t.Errorf("Error while calling .Place(): %v", err)
			}

			placeName := place.PlaceName()
			placeType := place.PlaceType()
			placeID := place.PlaceID()

			_, _, _ = placeName, placeType, placeID
		}
	}

	// For each of them, let's run a subtest
	for name, datum := range data {
		// Create the run function
		rfunc := rgen(datum)

		// Run !
		t.Run(name, rfunc)
	}
}

func TestPlaceCountainer_Check_NoCompare(t *testing.T) {
	// Get the input
	data := placeCountainers
	if len(data) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run it in parallel
	rgen := func(pc *PlaceCountainer) func(t *testing.T) {
		return func(t *testing.T) {
			err := pc.Check()
			if err != nil {
				t.Errorf("Check gave us invalid results: %v", err)
			}
		}
	}

	// For each of them, let's run a subtest
	for name, datum := range data {
		// Create the run function
		rfunc := rgen(datum)

		// Run !
		t.Run(name, rfunc)
	}
}
