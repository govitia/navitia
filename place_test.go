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

// TestPlaceCountainer_Place_NoCompare tests the PlaceCountainer.Place method
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

// TestPlaceCountainer_Check_NoCompare tests the PlaceCountainer.Check method
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

// BenchmarkPlaceCountainerCheck benchmarks Place.Check through subbenchmarks
func BenchmarkPlaceCountainerCheck(b *testing.B) {
	// Get the bench data
	data := testData["place"].bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in PlaceCountainer) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Call .Check
				_ = in.Check()
			}
		}
	}

	// Loop over all corpus
	for name, datum := range data {
		var pc = PlaceCountainer{}

		err := json.Unmarshal(datum, &pc)
		if err != nil {
			b.Errorf("Error while unmarshalling: %v", err)
		}

		// Get run function
		runFunc := runGen(pc)

		// Run it !
		b.Run(name, runFunc)
	}
}
