package testutils

import (
	"encoding/json"
	"reflect"
	"testing"
)

// UnmarshalAndCompare unmarshals a map[string]TestPair (which is the type of TestData.Known), and compares the unmarshalled data to the expected one.
//
// equalFunc is a function that is used to compare the unmarshalled value to the expected one
func UnmarshalAndCompare(t *testing.T, data map[string]TestPair, resultsType reflect.Type, equalFunc func(x, y interface{}) bool) {
	// if we have no data, skip
	if data == nil || len(data) == 0 {
		t.Skip("no data provided, skipping...")
	}

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(pair TestPair) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			// If the data is empty, skip
			// We don't check for len(pair.Raw)==0 in case that is a testcase
			if pair.Raw == nil {
				t.Skip("no raw data, skipping...")
			}
			if pair.Correct == nil {
				t.Skip("no correct value to compare to, skipping;..")
			}

			// Create a pointer to a new value of the type indicated in resultsType
			var res interface{} = reflect.New(resultsType).Interface()

			// We use encoding/json's unmarshaller, which internally uses the type's UnmarshalJSON if it exists
			err := json.Unmarshal(pair.Raw, res)
			if err != nil {
				t.Fatalf("error while unmarshalling: %v", err)
			}

			// Now we call equalFunc
			if ok := equalFunc(res, pair.Correct); !ok {
				t.Errorf("error: unmarshalled value not equal to correct value !")
			}
		}
	}

	// Loop around
	for name, pair := range data {
		t.Run(name, rgen(pair))
	}
}
