package navitia

import (
	"context"
	"testing"

	"github.com/aabizri/navitia/types"
)

func Test_Places_Basic(t *testing.T) {
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
		params.Around = types.Coordinates{48.847002, 2.377310}
		res, err := testSession.Places(ctx, params)
		t.Logf("Got results: \n%s", res.String())
		if err != nil {
			t.Fatalf("Got error in Places(): %v\n\tParameters: %#v", err, params)
		}
	})
}

func Test_PlacesResultsUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	pairs := testData["places"].known
	if len(pairs) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte) func(t *testing.T) {
		return func(t *testing.T) {
			var pr = &PlacesResults{}

			err := pr.UnmarshalJSON(data)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}

			str := pr.String()
			t.Log("\n" + str)
		}
	}

	// For each of them, let's run a subtest
	for name, pair := range pairs {
		// Get the run function
		rfunc := rgen(pair.raw)

		// Run !
		t.Run(name, rfunc)
	}
}
