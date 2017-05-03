package navitia

import (
	"context"
	"testing"

	"github.com/aabizri/navitia/types"
)

func Test_PlacesNear(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	params := PlacesNearRequest{}
	coords := types.Coordinates{Latitude: 48.847002, Longitude: 2.377310}

	// Create the root context
	ctx := context.Background()

	// Receive
	res, err := testSession.PlacesNear(ctx, coords, params)
	if err != nil {
		t.Fatalf("error in PlacesNear: %v\n\tParameters: %#v\n\tReceived: %#v", err, params, res)
	}

	// Log
	t.Logf("received %d places", len(res.Places))
}
