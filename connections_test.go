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
				t.Errorf("Got error: %v", err)
			}
		}
		arrFunc := func(t *testing.T) {
			res, err := testSession.ArrivalsSA(ctx, req, region, resource)
			t.Log(res)
			if err != nil {
				t.Errorf("Got error: %v", err)
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
