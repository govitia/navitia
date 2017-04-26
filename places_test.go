package navitia

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aabizri/navitia/types"
)

func Test_Places(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	params := PlacesRequest{
		Query: "avenue",
	}

	// Create the root context
	ctx := context.Background()

	// Run a simple search
	t.Run("simple", func(t *testing.T) {
		res, err := testSession.Places(ctx, params)
		t.Logf("Got results: \n%s", res.String())
		if err != nil {
			t.Fatalf("Got error in Places(): %v\n\tParameters: %#v", err, params)
		}
	})

	// Run a search with proximity
	t.Run("proximity", func(t *testing.T) {
		params.Around = types.Coordinates{Latitude: 48.847002, Longitude: 2.377310}
		res, err := testSession.Places(ctx, params)
		t.Logf("Got results: \n%s", res.String())
		if err != nil {
			t.Fatalf("Got error in Places(): %v\n\tParameters: %#v", err, params)
		}
	})
}

// Test_PlacesResults_Unmarshal tests unmarshalling for PlacesResults.
// As the unmarshalling is done by encoding/json, this allows us to check that the input can be reliably unmarshalled into the structure we have for it.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_PlacesResults_Unmarshal(t *testing.T) {
	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			var pr = &PlacesResults{}

			// We use encoding/json's unmarshaller, as we don't have one for this type
			err := json.Unmarshal(data, pr)

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
	correct := testData["places"].correct

	// Get the incorrect files
	incorrect := testData["places"].incorrect

	// Run !
	t.Run("correct", sub(correct, true))
	t.Run("incorrect", sub(incorrect, false))
}
