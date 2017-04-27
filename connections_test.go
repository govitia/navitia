package navitia

import (
	"context"
	"reflect"
	"strconv"
	"testing"

	"github.com/aabizri/navitia/types"
)

func TestConnectionsSA(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	region := types.ID("fr-idf")
	resources := []types.ID{
		"stop_area:OIF:SA:59346",
		"stop_area:OIF:SA:59586",
	}

	// Create the context
	ctx := context.Background()

	// Common request (for now)
	req := ConnectionsRequest{}

	// Create the run function generator, allowing us to run this in parallel
	//
	// Creates two versions: one calling DeparturesSA the other ArrivalsSA
	rgen := func(region types.ID, resource types.ID) (func(t *testing.T), func(t *testing.T)) {
		depFunc := func(t *testing.T) {
			res, err := testSession.Scope(region).DeparturesSA(ctx, req, resource)
			t.Log(res)
			if err != nil {
				t.Errorf("error in DeparturesSA: %v\n\tResource: %s\n\tParameters: %#v\n\tReceived: %#v", err, resource, req, res)
			}
		}
		arrFunc := func(t *testing.T) {
			res, err := testSession.Scope(region).ArrivalsSA(ctx, req, resource)
			t.Log(res)
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
	testUnmarshal(t, testData["connections"], reflect.TypeOf(ConnectionsResults{}))
}
