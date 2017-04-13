package navitia

import (
	"github.com/aabizri/navitia/types"
	"testing"
)

func Test_Places(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	params := PlacesRequest{
		Query: "avenue",
	}

	// Run a simple search
	t.Run("simple", func(t *testing.T) {
		res, err := testSession.Places(params)
		t.Logf("Got results: \n%s", res.String())
		if err != nil {
			t.Fatalf("Got error in Places(): %v\n\tParameters: %#v", err, params)
		}
	})

	// Run a search with proximity
	t.Run("proximity", func(t *testing.T) {
		params.Around = types.Coordinates{48.847002, 2.377310}
		res, err := testSession.Places(params)
		t.Logf("Got results: \n%s", res.String())
		if err != nil {
			t.Fatalf("Got error in Places(): %v\n\tParameters: %#v", err, params)
		}
	})
}
