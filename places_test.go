package navitia

import (
	"context"
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/aabizri/navitia/testutils"
	"github.com/aabizri/navitia/types"
)

func Test_Places_Online(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	query := "avenue"
	opt := PlacesRequest{
		Count: 1000, // We want the biggest count to cause the biggest stress
	}

	// Create the root context
	ctx := context.Background()

	// Run a simple search
	t.Run("simple", func(t *testing.T) {
		res, err := testSession.Places(ctx, query, opt)
		if err != nil {
			t.Fatalf("error in Places: %v\n\tParameters: %#v\n\tReceived: %#v", err, opt, res)
		}
	})

	// Run a search with proximity
	t.Run("proximity", func(t *testing.T) {
		opt.Around = types.Coordinates{Latitude: 48.847002, Longitude: 2.377310}
		res, err := testSession.Places(ctx, query, opt)
		if err != nil {
			t.Fatalf("error in Places: %v\n\tParameters: %#v\n\tReceived: %#v", err, opt, res)
		}
	})
}

func Test_Scope_Places_Sort(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	query := "avenue de la libération"
	opt := PlacesRequest{
		Count: 1000,
	}

	// Create the root context
	ctx := context.Background()

	// Run a simple search
	res, err := testSession.Scope("fr-idf").Places(ctx, query, opt)
	if err != nil {
		t.Fatalf("error in Places: %v\n\tParameters: %#v\n\tReceived: %#v", err, opt, res)
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

// Benchmark_PlacesResults_Unmarshal benchmarks PlacesResults unmarshalling via subbenchmarks
func Benchmark_PlacesResults_Unmarshal(b *testing.B) {
	// Get the bench data
	data := testData["places"].Bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a PlacesResults
				var ptr = &PlacesResults{}
				_ = json.Unmarshal(in, ptr)
			}
		}
	}

	// Loop over all corpus
	for name, datum := range data {
		// Get run function
		runFunc := runGen(datum)

		// Run it !
		b.Run(name, runFunc)
	}
}
