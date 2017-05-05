package navitia

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/aabizri/navitia/testutils"
	"github.com/aabizri/navitia/types"
)

func Test_PlacesNearby_Online(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	params := PlacesNearbyRequest{
		Count: 1000, // We want the biggest count to cause the biggest stress
	}
	coords := types.Coordinates{Latitude: 48.847002, Longitude: 2.377310}

	// Create the root context
	ctx := context.Background()

	// Receive
	res, err := testSession.PlacesNearby(ctx, coords, params)
	if err != nil {
		t.Fatalf("error in PlacesNearby: %v\n\tParameters: %#v\n\tReceived: %#v", err, params, res)
	}

	// Log
	t.Logf("received %d places", len(res.Places))
}

// Test_PlacesNearbyResults_Unmarshal tests unmarshalling for PlacesNearbyResults.
// As the unmarshalling is done by encoding/json, this allows us to check that the input can be reliably unmarshalled into the structure we have for it.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_PlacesNearbyResults_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["placesnearby"], reflect.TypeOf(PlacesNearbyResults{}))
}

// Benchmark_PlacesNearbyResults_Unmarshal benchmarks PlacesNearbyResults unmarshalling via subbenchmarks
func Benchmark_PlacesNearbyResults_Unmarshal(b *testing.B) {
	// Get the bench data
	data := testData["placesnearby"].Bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a PlacesNearbyResults
				var ptr = &PlacesNearbyResults{}
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
