package types

import (
	"testing"
)

// TestRegionUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestRegionUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	data := testData["region"].known
	if len(data) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run it in parallel
	rgen := func(in []byte) func(t *testing.T) {
		return func(t *testing.T) {
			var r = &Region{}

			err := r.UnmarshalJSON(in)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}
		}
	}

	// For each of them, let's run a subtest
	for name, datum := range data {
		// Get the run function
		rfunc := rgen(datum.raw)

		// Run !
		t.Run(name, rfunc)
	}
}

// BenchmarkRegionUnmarshal benchmarks Region unmarshalling via subbenchmarks
func BenchmarkRegionUnmarshal(b *testing.B) {
	// Get the bench data
	data := testData["region"].bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var r = &Region{}
				_ = r.UnmarshalJSON(in)
			}
		}
	}

	// Loop over all corpus
	for name, datum := range data {
		// Get run function
		runFunc := runGen(datum)

		// Run it !
		b.Run(name, runFunc)
	}
}
