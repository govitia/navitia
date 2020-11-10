package types

import (
	"reflect"
	"testing"
)

// Test_Journey_Unmarshal tests unmarshalling for Journey.
// As the unmarshalling is done in-house, this allows us to check that the custom UnmarshalJSON function correctly
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_Journey_Unmarshal(t *testing.T) {
	testUnmarshal(t, testData["journey"], reflect.TypeOf(Journey{}))
}

// BenchmarkJourney_UnmarshalJSON benchmarks Journey unmarshalling via subbenchmarks
func BenchmarkJourney_UnmarshalJSON(b *testing.B) {
	// Get the bench data
	data := testData["journey"].bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a Journey
				j := &Journey{}
				_ = j.UnmarshalJSON(in)
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
