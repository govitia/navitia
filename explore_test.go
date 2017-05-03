package navitia

import (
	"context"
	"reflect"
	"testing"

	"github.com/aabizri/navitia/testutils"
)

func Test_Explore(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	params := ExploreRequest{}
	scope := testSession.Scope("fr-idf")

	// Create the root context
	ctx := context.Background()

	res, err := scope.Explore(ctx, LinesSelector, params)
	if err != nil {
		t.Fatalf("error in Explore: %v\n\tParameters: %#v\n\tReceived: %#v", err, params, res)
	}

}

// Test_ExploreResults_Unmarshal tests unmarshalling for ExploreResults.
// As the unmarshalling is done by ourselves, this allows us to check that our unmarshaller works correctly.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_ExploreResults_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["explore"], reflect.TypeOf(ExploreResults{}))
}

// Benchmark_ExploreResults_Unmarshal benchmarks ExploreResults unmarshalling via subbenchmarks
func Benchmark_ExploreResults_Unmarshal(b *testing.B) {
	// Get the bench data
	data := testData["explore"].Bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a ExploreResults
				var er = &ExploreResults{}
				_ = er.UnmarshalJSON(in)
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
