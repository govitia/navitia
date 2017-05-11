package navitia

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/aabizri/navitia/testutils"
	"github.com/aabizri/navitia/types"
)

func Test_StopSchedules_Online(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	var (
		opts = StopSchedulesRequest{
		//Count: 1000, // We want the biggest count to cause the biggest stress
		//Depth: 3,    // Same reasoning
		}
		scope            = testSession.Scope("fr-idf")
		resType          = "stop_area"
		resID   types.ID = "stop_area:OIF:SA:59566"
		coords           = types.Coordinates{Latitude: 48.842716, Longitude: 2.384471} // 110 Avenue Daumesnil (Paris)
	)

	// Create the root context
	ctx := context.Background()

	t.Run("StopSchedules", func(t *testing.T) {
		res, err := scope.StopSchedules(ctx, resID, opts)
		if err != nil {
			t.Fatalf("error in StopSchedules: %v\n\tResource ID: %s\n\tParameters: %#v\n\tReceived: %#v", err, resID, opts, res)
		}
	})

	t.Run("StopSchedulesExplicit", func(t *testing.T) {
		res, err := scope.StopSchedulesExplicit(ctx, resType, resID, opts)
		if err != nil {
			t.Fatalf("error in StopSchedulesExplicit: %v\n\tResource Type: %s\n\tResource ID: %s\n\tParameters: %#v\n\tReceived: %#v", err, resType, resID, opts, res)
		}
	})

	t.Run("StopSchedulesCoords", func(t *testing.T) {
		res, err := scope.StopSchedulesCoords(ctx, coords, opts)
		if err != nil {
			t.Fatalf("error in StopSchedules: %v\n\tCoordinates: %s\n\tParameters: %#v\n\tReceived: %#v", err, coords.ID(), opts, res)
		}
	})

}

// Test_StopSchedulesResults_Unmarshal tests unmarshalling for StopSchedulesResults.
// As the unmarshalling is done by encoding/json, this allows us to check that our struct correctly fits the data.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_StopSchedulesResults_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["stop_schedules"], reflect.TypeOf(StopSchedulesResults{}))
}

// Benchmark_StopSchedulesResults_Unmarshal benchmarks StopSchedulesResults unmarshalling via subbenchmarks
func Benchmark_StopSchedulesResults_Unmarshal(b *testing.B) {
	// Get the bench data
	data := testData["stop_schedules"].Bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a StopSchedulesResults
				var er = &StopSchedulesResults{}
				_ = json.Unmarshal(in, er)
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
