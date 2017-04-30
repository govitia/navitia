package navitia

import (
	"context"
	"reflect"
	"sort"
	"testing"

	"github.com/aabizri/navitia/testutils"
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
		if err != nil {
			t.Fatalf("error in Places: %v\n\tParameters: %#v\n\tReceived: %#v", err, params, res)
		}
	})

	// Run a search with proximity
	t.Run("proximity", func(t *testing.T) {
		params.Around = types.Coordinates{Latitude: 48.847002, Longitude: 2.377310}
		res, err := testSession.Places(ctx, params)
		if err != nil {
			t.Fatalf("error in Places: %v\n\tParameters: %#v\n\tReceived: %#v", err, params, res)
		}
	})
}

func Test_Scope_Places_Sort(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	params := PlacesRequest{
		Query: "avenue de la libération",
	}

	// Create the root context
	ctx := context.Background()

	// Run a simple search
	res, err := testSession.Scope("fr-idf").Places(ctx, params)
	if err != nil {
		t.Fatalf("error in Places: %v\n\tParameters: %#v\n\tReceived: %#v", err, params, res)
	}

	// Check if sorted
	if !sort.IsSorted(sort.Reverse(res)) {
		t.Errorf("PlacesResults isn't sorted !")
	}
}

// Test_PlacesResults_Unmarshal tests unmarshalling for PlacesResults.
// As the unmarshalling is done by encoding/json, this allows us to check that the input can be reliably unmarshalled into the structure we have for it.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_PlacesResults_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["places"], reflect.TypeOf(PlacesResults{}))
}
