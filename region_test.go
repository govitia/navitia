package navitia

import (
	"context"
	"encoding/json"
	"testing"
)

func Test_Regions(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	ctx := context.Background()
	req := RegionRequest{}

	// Run the query with GeoJSON
	t.Run("with_geojson", func(t *testing.T) {
		res, err := testSession.Regions(ctx, req)
		if err != nil {
			t.Fatalf("error in Regions: %v\n\tParameters: %#v\n\tReceived: %#v", err, req, res)
		}
	})

	// Run the query without GeoJSON
	t.Run("without_geojson", func(t *testing.T) {
		req := req
		req.DisableGeoJSON = true
		res, err := testSession.Regions(ctx, req)
		if err != nil {
			t.Fatalf("error in Regions: %v\n\tParameters: %#v\n\tReceived: %#v", err, req, res)
		}
	})
}

// Test_RegionResults_Unmarshal tests unmarshalling for RegionResults.
// As the unmarshalling is done by encoding/json, this allows us to check that the input can be reliably unmarshalled into the structure we have for it.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_RegionResults_Unmarshal(t *testing.T) {
	// Declare this test to be run in parallel
	t.Parallel()

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			var rr = &RegionResults{}

			// We use encoding/json's unmarshaller, as we don't have one for this type
			err := json.Unmarshal(data, rr)

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
	correct := testData["coverage"].correct

	// Get the incorrect files
	incorrect := testData["coverage"].incorrect

	// Run !
	t.Run("correct", sub(correct, true))
	t.Run("incorrect", sub(incorrect, false))
}
