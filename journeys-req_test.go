package gonavitia

import "testing"

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
