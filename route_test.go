package types

import (
	"encoding/json"
	"testing"
)

// TestRouteUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestRouteUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	pairs := testData["route"].known
	if len(pairs) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte) func(t *testing.T) {
		return func(t *testing.T) {
			var r = &Route{}

			err := json.Unmarshal(data, r)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}

			t.Logf("Result: %#v", r)
		}
	}

	// For each of them, let's run a subtest
	for name, pair := range pairs {
		// Get the run function
		rfunc := rgen(pair.raw)

		// Run !
		t.Run(name, rfunc)
	}
}
