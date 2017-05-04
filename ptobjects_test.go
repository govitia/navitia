package navitia

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/aabizri/navitia/testutils"
	"github.com/aabizri/navitia/types"
)

func Test_PTObjects(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	scope := testSession.Scope("fr-idf")
	query := "chapelle"
	params := PTObjectsRequest{
		// Types takes on the default value
		Types: []string{types.EmbeddedNetwork, types.EmbeddedCommercialMode, types.EmbeddedLine, types.EmbeddedRoute, types.EmbeddedStopArea},

		// Enable Geo
		Geo: true,

		// Set count to 1000
		Count: 1000,
	}

	// Create the root context
	ctx := context.Background()

	// Launch
	res, err := scope.PTObjects(ctx, query, params)
	if err != nil {
		t.Fatalf("error in PTObjects: %v\n\tQuery: %s\n\tParameters: %#v\n\tReceived: %#v", err, query, params, res)
	}
}

// Test_PTObjectsResults_Unmarshal tests unmarshalling for PTObjectsResults.
// As the unmarshalling is done by ourselves, this allows us to check that our unmarshaller works correctly.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_PTObjectsResults_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["ptobjects"], reflect.TypeOf(PTObjectsResults{}))
}

// Benchmark_PTObjectsResults_Unmarshal benchmarks PTObjectsResults unmarshalling via subbenchmarks
func Benchmark_PTObjectsResults_Unmarshal(b *testing.B) {
	// Get the bench data
	data := testData["ptobjects"].Bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a PTObjectsResults
				var ptr = &PTObjectsResults{}
				_ = json.Unmarshal(in, ptr)
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
