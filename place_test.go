package types

import (
	"encoding/json"
	"strconv"
	"testing"
)

// TestPlaceUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestPlaceUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	input := testData["place"]
	if len(input) == 0 {
		t.Skip("No data to test")
	}

	// For each of them, let's run a subtest
	for i, file := range input {
		// Create a name for this run
		var name string
		stat, err := file.Stat()
		if err != nil {
			t.Errorf("Error while retrieving name for pass %d: %v", i, err)
			name = strconv.Itoa(i)
		} else {
			name = stat.Name()
		}

		// Create the run function
		rfunc := func(t *testing.T) {
			var j = &PlaceCountainer{}
			dec := json.NewDecoder(file)
			err := dec.Decode(j)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}

			_, err = j.Place()
			if err != nil {
				t.Errorf("Error while retrieving Place(): %v", err)
			}
		}

		// Run !
		t.Run(name, rfunc)
	}
}
