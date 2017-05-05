package navitia

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	"github.com/aabizri/navitia/testutils"
	"github.com/aabizri/navitia/types"
)

func Test_ConnectionsSA_Online(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	region := types.ID("fr-idf")
	resources := []types.ID{
		"stop_area:OIF:SA:59346",
	}

	// Create the context
	ctx := context.Background()

	// Common request (for now)
	req := ConnectionsRequest{
		Count: 1000, // We want the biggest count to cause the biggest stress
	}

	// Create the run function generator, allowing us to run this in parallel
	//
	// Creates two versions: one calling DeparturesSA the other ArrivalsSA
	rgen := func(region types.ID, resource types.ID) (func(t *testing.T), func(t *testing.T)) {
		depFunc := func(t *testing.T) {
			res, err := testSession.Scope(region).DeparturesSA(ctx, req, resource)

			if err != nil {
				t.Errorf("error in DeparturesSA: %v\n\tResource: %s\n\tParameters: %#v\n\tReceived: %#v", err, resource, req, res)
			}
		}
		arrFunc := func(t *testing.T) {
			res, err := testSession.Scope(region).ArrivalsSA(ctx, req, resource)

			if err != nil {
				t.Errorf("error in ArrivalsSA: %v\n\tResource: %s\n\tParameters: %#v\n\tReceived: %#v", err, resource, req, res)
			}
		}
		return depFunc, arrFunc
	}

	// For each of them, let's run a subtest
	for i, sa := range resources {
		// Get the run function
		depFunc, arrFunc := rgen(region, sa)

		// Create the name
		name := strconv.Itoa(i)

		// Run !
		t.Run(name+"_departures", depFunc)
		t.Run(name+"_arrivals", arrFunc)
	}

}

// Test_ConnectionsResults_Unmarshal tests unmarshalling for ConnectionsResults.
// As the unmarshalling is done in-house, this allows us to check that the custom UnmarshalJSON function correctly
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_ConnectionsResults_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["connections"], reflect.TypeOf(ConnectionsResults{}))
}

// Benchmark_ConnectionsResults_Unmarshal benchmarks ConnectionsResults unmarshalling via subbenchmarks
func Benchmark_ConnectionsResults_Unmarshal(b *testing.B) {
	// Get the bench data
	data := testData["connections"].Bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a ConnectionsResults
				var ptr = &ConnectionsResults{}
				_ = ptr.UnmarshalJSON(in)
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

func Test_ConnectionsResults_Unmarshal_Compare(t *testing.T) {
	equalFunc := func(x, y interface{}) bool {
		a, ok := x.(*ConnectionsResults)
		if !ok {
			return false
		}

		b, ok := y.(*ConnectionsResults)
		if !ok {
			return false
		}

		if len(a.Connections) != len(b.Connections) {
			return false
		}

		/*for i := 0; i < len(a.Connections); i++ {
			// compare a.Connections[i] and b.Connections[i]
		}*/

		return true
	}
	testutils.UnmarshalAndCompare(t, knownConnections, reflect.TypeOf(ConnectionsResults{}), equalFunc)
}

var knownConnections = map[string]testutils.TestPair{
	"one": {
		Raw: []byte(`
{
	"departures": [{
		"display_informations": {
			"code": "4"
		}
	},
	{
		"display_informations": {
			"code": "4"
		}
	}
	]
}`),
		Correct: &ConnectionsResults{
			Connections: []types.Connection{
				{
					Display: types.Display{
						Code: "4",
					},
				},
				{
					Display: types.Display{
						Code: "4",
					},
				},
			},
		},
	},
}
