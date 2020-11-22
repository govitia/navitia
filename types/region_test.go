package types

import (
	"fmt"
	"reflect"
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
	testUnmarshal(t, testData["region"], reflect.TypeOf(Region{}))
}

// TestRegionUnmarshal_ShapeInvalidMKT tests known invalid MKT (well-known text) -encoded Region.Shape inputs for (*Region).UnmarshalJSON.
func TestRegionUnmarshal_ShapeInvalidMKT(t *testing.T) {
	// Shapes
	shapes := [...]string{
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

// BenchmarkRegionUnmarshal benchmarks Region unmarshalling via subbenchmarks.
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
				r := &Region{}
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
