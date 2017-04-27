package types

import (
	"fmt"
	"strings"
	"testing"
)

// Test_Region_Unmarshal tests unmarshalling for Region.
// As the unmarshalling is done in-house, this allows us to check that the custom UnmarshalJSON function correctly
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_Region_Unmarshal(t *testing.T) {
	// Declare this test to be run in parallel
	t.Parallel()

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			var r = &Region{}

			// We use encoding/json's unmarshaller, as we don't have one for this type
			err := r.UnmarshalJSON(data)

			// We check that the result is what we expect:
			// 	If we expect no errors (correct == true) but we get one, the test has failed
			//	If we expect an error (correct == false) but we don't get one, the test has failes
			// 	In all other cases, the test is successful !
			if err != nil && correct {
				t.Errorf("expected no errors but got one: %v", err)
			} else if err == nil && !correct {
				t.Errorf("expected an error but didn't get one !")
			}
		}
	}

	// Create the sub functions (those will be the correct and incorrect version of this test)
	sub := func(data map[string][]byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// If we have no data, we skip
			if len(data) == 0 {
				t.Skip("no data provided, skipping...")
			}

			// For all files provided
			for name, datum := range data {
				// Get the run function
				rfunc := rgen(datum, correct)

				// Run !
				t.Run(name, rfunc)
			}
		}
	}

	// Get the correct files
	correct := testData["region"].correct

	// Get the incorrect files
	incorrect := testData["region"].incorrect

	// Run !
	t.Run("correct", sub(correct, true))
	t.Run("incorrect", sub(incorrect, false))
}

// TestRegionUnmarshal_ShapeInvalidMKT tests known invalid MKT (well-known text) -encoded Region.Shape inputs for (*Region).UnmarshalJSON
func TestRegionUnmarshal_ShapeInvalidMKT(t *testing.T) {
	// Shapes
	var shapes = [...]string{
		"MULTIPOLYGON(((-11.33535 51.29165,0,0,-11.33535 51.29165))",
		"MULTIPOLYGON((",
		"MULTIPOLYGON(((-11.33535 51.29165,0,0,-11.33535 51.29165)",
		"MULTIPOLYGON(",
		"POLYGON(",
		"MULTIPOLYGON(((0",
	}

	// Run
	for i, s := range shapes {
		in := []byte(fmt.Sprintf(`{"shape": "%s"}`, s))
		r := &Region{}
		err := r.UnmarshalJSON(in)
		if err == nil {
			t.Errorf("No error in run #%d even though we expected one", i)
		} else if !strings.Contains(err.Error(), "EOF") {
			t.Errorf("Unexpected error in run #%d with [%s]: %v", i, in, err)
		}
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
