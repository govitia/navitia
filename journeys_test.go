package navitia

import (
	"context"
	"reflect"
	"testing"

	"github.com/aabizri/navitia/testutils"
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
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_JourneysResults_Unmarshal(t *testing.T) {
	testutils.UnmarshalTest(t, testData["journeys"], reflect.TypeOf(JourneyResults{}))
}
