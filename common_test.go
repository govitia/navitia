package navitia

import (
	"encoding/json"
	"reflect"
	"testing"
)

// testUnmarshal is a helper to test unmarshalling for any value implementing results.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func testUnmarshal(t *testing.T, data typeTestData, resultsType reflect.Type) {
	t.Helper()
	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			// Create a pointer to a new value of the type indicated in resultsType
			// We know it is a results, so we assert it, this way we don't get any silent fails.
			var res = reflect.New(resultsType).Interface().(results)

			// We use encoding/json's unmarshaller, as we don't have one for this type
			err := json.Unmarshal(data, res)

			// We check that the result is what we expect:
			// 	If we expect no errors (correct == true) but we get one, the test has failed
			//	If we expect an error (correct == false) but we don't get one, the test has failes
			// 	In all other cases, the test is successful !
			if err != nil && correct {
				t.Errorf("expected no errors but got one: %v", err)
			} else if err == nil && !correct {
				t.Errorf("expected an error but didn't get one !")
			}
		}
	}

	// Create the sub functions (those will be the correct and incorrect version of this test)
	sub := func(data map[string][]byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			// If we have no data, we skip
			if len(data) == 0 {
				t.Skip("no data provided, skipping...")
			}

			// For all files provided
			for name, datum := range data {
				// Get the run function
				rfunc := rgen(datum, correct)

				// Run !
				t.Run(name, rfunc)
			}
		}
	}

	// Run !
	t.Run("correct", sub(data.correct, true))
	t.Run("incorrect", sub(data.incorrect, false))
}
