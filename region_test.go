package navitia

import (
	"context"
	"testing"
)

func Test_Regions(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	ctx := context.Background()
	req := RegionRequest{}

	// Run the query with GeoJSON
	t.Run("with_geojson", func(t *testing.T) {
		res, err := testSession.Regions(ctx, req)
		t.Logf("Received res: %v", *res)
		if err != nil {
			t.Fatalf("Got error in Regions: %v", err)
		}
	})

	// Run the query without GeoJSON
	t.Run("without_geojson", func(t *testing.T) {
		req := req
		req.DisableGeoJSON = true
		res, err := testSession.Regions(ctx, req)
		t.Logf("Received res: %v", *res)
		if err != nil {
			t.Fatalf("Got error in Regions: %v", err)
		}
	})
}
