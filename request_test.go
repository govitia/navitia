package navitia

import (
	"context"
	"testing"
)

// TestRequestContextCancellation tests that request is correctly canceled when the context is canceled.
//
// It uses a mockSession, which has a client that doesn't cancel when provided with a canceled context,
// allowing us to correctly test the behaviour we're looking for.
func TestRequestContextCancellation(t *testing.T) {
	// Create the already-canceled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Create the mock session
	mock := mockSession(t)

	// Launch
	err := mock.request(ctx, "https://api.navitia.io/v1/", &PlacesRequest{}, &PlacesResults{})
	if err != context.Canceled {
		t.Errorf("request didn't cancel request, got instead: %v", err)
	}
}
