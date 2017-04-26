package types

import "testing"

// TestIDCheck checks if ID.Check returns an error when given an empty ID
func TestIDCheck(t *testing.T) {
	id := ID("")
	if id.Check() == nil {
		t.Errorf("Received no error even though we expect one")
	}
}
