package types

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestJourneyUnmarshall(t *testing.T) {
	// First let's open the file
	path := filepath.Join(testDataPath, "journeys-A.json")
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("Error while opening file: %v", err)
	}

	// Now let's unmarshal this
	var j = &Journey{}
	dec := json.NewDecoder(f)
	err = dec.Decode(j)
	if err != nil {
		t.Errorf("Error while unmarshalling: %v", err)
	}

	t.Logf("%#v", *j)
}
