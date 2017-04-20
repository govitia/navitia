package types

import (
	"testing"
	"time"
)

// TestJourneyString shocases the String method without first unmarshalling
func TestJourneyString(t *testing.T) {
	from := Container{
		ID:   "2.399803859568057;48.88150165806373",
		Name: "54 Boulevard d'Algérie (Paris)",
		embeddedObject: Address{
			ID:    "2.399803859568057;48.88150165806373",
			Label: "54 Boulevard d'Algérie (Paris)",
			Name:  "Boulevard d'Algérie",
		},
	}

	to := Container{
		ID:   "2.344404;48.835114",
		Name: "54 Boulevard Arago (Paris)",
		embeddedObject: Address{
			ID:    "2.344404;48.835114",
			Label: "54 Boulevard Arago (Paris)",
			Name:  "Boulevard Arago",
		},
	}

	departure, err := time.Parse("2006-01-02T15:04:05", "2017-04-11T21:33:55")
	if err != nil {
		t.Fatal(err)
	}
	arrival, err := time.Parse("2006-01-02T15:04:05", "2017-04-11T22:24:13")
	if err != nil {
		t.Fatal(err)
	}

	section := Section{
		From:      from,
		To:        to,
		Departure: departure,
		Arrival:   arrival,
		Duration:  time.Duration(3018) * time.Second,
		Display: Display{
			Label:        "11",
			PhysicalMode: "Métro",
		},
	}

	journey := Journey{
		From:      from,
		To:        to,
		Departure: departure,
		Arrival:   arrival,
		Duration:  time.Duration(3018) * time.Second,
		Sections:  []Section{section},
	}

	t.Logf("For journey we have: %s", journey.String())
}

// TestJourneyUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestJourneyUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	pairs := testData["journey"].known
	if len(pairs) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte) func(t *testing.T) {
		return func(t *testing.T) {
			var j = &Journey{}

			err := j.UnmarshalJSON(data)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}

			str := j.String()
			t.Log("\n" + str)
		}
	}

	// For each of them, let's run a subtest
	for name, pair := range pairs {
		// Get the run function
		rfunc := rgen(pair.raw)

		// Run !
		t.Run(name, rfunc)
	}
}

// BenchmarkJourney_UnmarshalJSON benchmarks Journey unmarshalling via subbenchmarks
func BenchmarkJourney_UnmarshalJSON(b *testing.B) {
	// Get the bench data
	data := testData["journey"].bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a Journey
				var j = &Journey{}
				_ = j.UnmarshalJSON(in)
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
