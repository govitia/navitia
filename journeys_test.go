package navitia

import (
	"context"
	"testing"

	"github.com/aabizri/navitia/types"
)

func Test_JourneyRequest_toUrl(t *testing.T) {
	// First an empty struct
	req, err := JourneyRequest{}.toURL()
	if err != nil {
		t.Errorf("failure: toURL returned error: %v", err)
	}
	if len(req) != 0 {
		t.Errorf("failure: toURL created fields for non-specified parameters")
	}
	t.Logf("Result: %v", req)
}

func Test_Journeys(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	ctx := context.Background()

	params := JourneyRequest{}
	coords := types.Coordinates{48.847002, 2.377310}
	params.From = coords

	res, err := testSession.Journeys(ctx, params)
	t.Logf("Got results: \n%s", res.String())
	if err != nil {
		t.Fatalf("Got error in Journey(): %v\n\tParameters: %#v", err, params)
	}
}

func Test_Journeys_Paging(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	ctx := context.Background()

	params := JourneyRequest{
		From: types.Coordinates{Latitude: 48.842716, Longitude: 2.384471}, // 110 Avenue Daumesnil (Paris)
		To:   types.Coordinates{Latitude: 48.867305, Longitude: 2.352005}, // 10 Rue du Caire (Paris)
	}

	res, err := testSession.Journeys(ctx, params)
	t.Logf("Got results: \n%s", res.String())
	t.Logf("Paging: %#v", res.Paging)
	if err != nil {
		t.Fatalf("Got error in Journey(): %v\n\tParameters: %#v", err, params)
	}

	var paginated *JourneyResults = res
	for i := 0; paginated.Paging.Next != nil && i < 6; i++ {
		p := JourneyResults{}
		err = paginated.Paging.Next(ctx, testSession, &p)
		t.Logf("Next n°%d results:\n%s", i, p.String())
		if err != nil {
			t.Fatalf("Got error in Paging.Next (pass %d): %v", i, err)
		}
		paginated = &p
	}
}
