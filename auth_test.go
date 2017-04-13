package navitia

import (
	"testing"
)

func Test_New(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}
	_, err := New(*apiKey)
	if err != nil {
		t.Fatalf("Error while creating new session: %v", err)
	}
}
