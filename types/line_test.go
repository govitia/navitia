package types

import (
	"reflect"
	"testing"

	"github.com/aabizri/navitia/testutils"
)

// Test_Line_Unmarshal tests unmarshalling for Line.
// As the unmarshalling is done in-house, this allows us to check that the custom UnmarshalJSON function correctly
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_Line_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["line"], reflect.TypeOf(Line{}))
}

// BenchmarkLineUnmarshal benchmarks Line unmarshalling via subbenchmarks
func BenchmarkLineUnmarshal(b *testing.B) {
	// Get the bench data
	data := testData["line"].Bench
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
