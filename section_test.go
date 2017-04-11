package types

import (
	"encoding/json"
	"strconv"
	"testing"
)

// TestSectionUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestSectionUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	input := testData["section"]
	if len(input) == 0 {
		t.Skip("No data to test")
	}

	// For each of them, let's run a subtest
	for i, reader := range input {
		// Create a name for this run
		name := strconv.Itoa(i)

		// Create the run function
		rfunc := func(t *testing.T) {
			var j = &Section{}
			dec := json.NewDecoder(reader)
			err := dec.Decode(j)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}
		}

		// Run !
		t.Run(name, rfunc)
	}
}
