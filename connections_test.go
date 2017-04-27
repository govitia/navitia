package navitia

import (
	"context"
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
			res, err := testSession.DeparturesSA(ctx, req, region, resource)
			t.Log(res)
			if err != nil {
				t.Errorf("error in DeparturesSA: %v\n\tResource: %s\n\tParameters: %#v\n\tReceived: %#v", err, resource, req, res)
			}
		}
		arrFunc := func(t *testing.T) {
			res, err := testSession.ArrivalsSA(ctx, req, region, resource)
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
	// Declare this test to be run in parallel
	t.Parallel()

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			var cr = &ConnectionsResults{}

			// We use encoding/json's unmarshaller, as we don't have one for this type
			err := cr.UnmarshalJSON(data)

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
	correct := testData["connections"].correct

	// Get the incorrect files
	incorrect := testData["connections"].incorrect

	// Run !
	t.Run("correct", sub(correct, true))
	t.Run("incorrect", sub(incorrect, false))
}
