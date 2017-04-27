package navitia

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aabizri/navitia/types"
)

func Test_JourneyRequest_toUrl(t *testing.T) {
	// Declare this test to be run in parallel
	t.Parallel()

	req, err := JourneyRequest{}.toURL()
	if err != nil {
		t.Fatalf("error in JourneyRequest.ToURL: %v\n\tReceived: %#v", err, req)
	}
	if len(req) != 0 {
		t.Fatalf("error in JourneyRequest.ToURL: toURL created fields for non-specified parameters\n\tReceived: %#v", req)
	}
}

func Test_Journeys(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	ctx := context.Background()

	req := JourneyRequest{}
	coords := types.Coordinates{Latitude: 48.847002, Longitude: 2.377310}
	req.From = coords.ID()

	res, err := testSession.Journeys(ctx, req)
	if err != nil {
		t.Fatalf("error in Journeys: %v\n\tParameters: %#v\n\tReceived: %#v", err, req, res)
	}
}

func Test_Journeys_Paging(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	ctx := context.Background()

	params := JourneyRequest{
		From: types.Coordinates{Latitude: 48.842716, Longitude: 2.384471}.ID(), // 110 Avenue Daumesnil (Paris)
		To:   types.Coordinates{Latitude: 48.867305, Longitude: 2.352005}.ID(), // 10 Rue du Caire (Paris)
	}

	res, err := testSession.Journeys(ctx, params)
	if err != nil {
		t.Fatalf("error in initial call to Journeys: %v\n\tParameters: %#v\n\tReceived: %#v", err, params, res)
	}

	var i uint
	for i = 0; res.Paging.Next != nil && i < 6; i++ {
		p := JourneyResults{}
		err = res.Paging.Next(ctx, testSession, &p)
		if err != nil {
			t.Fatalf("error in call #%d to res.Paging.Next: %v\n\tReceived: %#v", i, err, p)
		}
		res = &p
	}
	t.Logf("Paging finished with %d iterations", i)
}

// Test_JourneysResults_Unmarshal tests unmarshalling for JourneyResults.
// As the unmarshalling is done by encoding/json, this allows us to check that the input can be reliably unmarshalled into the structure we have for it.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_JourneysResults_Unmarshal(t *testing.T) {
	// Declare this test to be run in parallel
	t.Parallel()

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			var jr = &JourneyResults{}

			// We use encoding/json's unmarshaller, as we don't have one for this type
			err := json.Unmarshal(data, jr)

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
			// Declare this test to be run in parallel
			t.Parallel()

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
	correct := testData["journeys"].correct

	// Get the incorrect files
	incorrect := testData["journeys"].incorrect

	// Run !
	t.Run("correct", sub(correct, true))
	t.Run("incorrect", sub(incorrect, false))
}
