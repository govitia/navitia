package types

import "testing"

// TestLineUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestLineUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	pairs := testData["line"].known
	if len(pairs) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte) func(t *testing.T) {
		return func(t *testing.T) {
			var l = &Line{}

			err := l.UnmarshalJSON(data)
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

// BenchmarkLineUnmarshal benchmarks Line unmarshalling via subbenchmarks
func BenchmarkLineUnmarshal(b *testing.B) {
	// Get the bench data
	data := testData["line"].bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var l = &Line{}
				_ = l.UnmarshalJSON(in)
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
