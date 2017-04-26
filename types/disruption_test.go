package types

import "testing"

// TestDisruptionUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestDisruptionUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	pairs := testData["disruption"].known
	if len(pairs) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte) func(t *testing.T) {
		return func(t *testing.T) {
			var d = &Disruption{}

			err := d.UnmarshalJSON(data)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}
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
