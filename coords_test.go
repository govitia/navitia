package navitia

import (
	"context"
	"testing"

	"github.com/aabizri/navitia/types"
)

func Test_Coords(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	coords := types.Coordinates{Latitude: 48.847002, Longitude: 2.377310}

	// Create the root context
	ctx := context.Background()

	// Receive
	address, regionID, err := testSession.Coords(ctx, coords)
	if err != nil {
		t.Fatalf("error in (*Session).Coords: %v\n\tInput coordinates: %s\n\tReceived:\n\t\tAddress: %s\n\t\tRegion: %s\n", err, coords.ID(), address.Label, regionID)
	}
}
