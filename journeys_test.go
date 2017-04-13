package navitia

import (
	"github.com/aabizri/navitia/types"
	"testing"
)

// Return a journey request
// TODO: Add random
func helperGenerateJourneyRequest() JourneyRequest {
	return JourneyRequest{}
}

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

	params := helperGenerateJourneyRequest()
	coords := types.Coordinates{48.847002, 2.377310}
	params.From = coords

	res, err := testSession.Journeys(params)
	t.Logf("Got results: \n%s", res.String())
	if err != nil {
		t.Fatalf("Got error in Journey(): %v\n\tParameters: %#v", err, params)
	}
}
