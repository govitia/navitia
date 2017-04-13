package types

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"
)

// TestJourneyString tests the String method
func TestJourneyString(t *testing.T) {
	from := Address{
		ID:    "2.399803859568057;48.88150165806373",
		Label: "54 Boulevard d'Algérie (Paris)",
		Name:  "Boulevard d'Algérie",
	}

	to := Address{
		ID:    "2.344404;48.835114",
		Label: "54 Boulevard Arago (Paris)",
		Name:  "Boulevard Arago",
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
		Display: DisplayInformations{
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

	want := "Boulevard d'Algérie (11/04 @ 21:33) --(50m18s)--> Boulevard Arago (11/04 @ 22:24)\n\t0: Boulevard d'Algérie (11/04 @ 21:33) --(Métro 11 | 50m18s)--> Boulevard Arago (11/04 @ 22:24)"

	if journey.String() != want {
		t.Error("Output of String isn't what was expected")
	}

	t.Logf("For journey we have: %s", journey.String())
}

// TestJourneyUnmarshal_NoCompare tries to unmarshal all json test data for this type, but doesn't compare its response to a known correct output.
func TestJourneyUnmarshal_NoCompare(t *testing.T) {
	// Get the input
	input := testData["journey"]
	if len(input) == 0 {
		t.Skip("No data to test")
	}

	// For each of them, let's run a subtest
	for i, reader := range input {
		// Create a name for this run
		name := strconv.Itoa(i)

		// Create the run function
		rfunc := func(t *testing.T) {
			var j = &Journey{}
			dec := json.NewDecoder(reader)
			err := dec.Decode(j)
			if err != nil {
				t.Errorf("Error while unmarshalling: %v", err)
			}
		}

		// Run !
		t.Run(name, rfunc)
	}
}
