package navitia

import (
	"context"
	"testing"
)

func TestRequestContextCancellation(t *testing.T) {
	// Create the already-canceled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	// Create the mock session
	mock := &Session{client: newMockClient(t)}

	// Launch
	err := mock.request(ctx, "https://api.navitia.io/v1/", &PlacesRequest{}, &PlacesResults{})
	if err != context.Canceled {
		t.Errorf("request didn't cancel request, got instead: %v", err)
	}
}
