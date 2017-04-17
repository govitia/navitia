package navitia

import (
	"testing"
)

func Test_Coverage(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	res, err := testSession.Coverage(0)
	t.Logf("Received res: %v", *res)
	if err != nil {
		t.Fatalf("Got error in Coverage(%d): %v", 0, err)
	}
}
