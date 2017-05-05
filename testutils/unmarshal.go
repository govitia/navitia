package testutils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// UnmarshalTest is a helper to test unmarshalling for any value implementing results.
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func UnmarshalTest(t *testing.T, data *TestData, resultsType reflect.Type) {
	// Skip if no data is provided
	if data == nil {
		t.Skip("no data provided, skipping...")
	}

	// Create the run function generator, allowing us to run this in parallel
	rgen := func(data []byte, correct bool) func(t *testing.T) {
		return func(t *testing.T) {
			// Declare this test to be run in parallel
			t.Parallel()

			// Create a pointer to a new value of the type indicated in resultsType
			var res interface{} = reflect.New(resultsType).Interface()

			// We use encoding/json's unmarshaller, which internally uses the type's UnmarshalJSON if it exists
			err := json.Unmarshal(data, res)

			// We check that the result is what we expect:
			// 	If we expect no errors (correct == true) but we get one, the test has failed
			//	If we expect an error (correct == false) but we don't get one, the test has failes
			// 	In all other cases, the test is successful !
			if err != nil && correct {
				var msg string
				switch concrete := err.(type) {
				case *json.SyntaxError:
					msg += fmt.Sprintf("syntax error: syntax error after %d bytes", concrete.Offset)
				case *json.UnmarshalTypeError:
					msg += fmt.Sprintf("type error: JSON value (%s) not appropriate for the type (%s) of the key (json field: \"%s\") at offset %d",
						concrete.Value,
						concrete.Type.Name(),
						concrete.Field,
						concrete.Offset,
					)
				}
				if msg != "" {
					msg += ": "
				}
				t.Errorf("expected no errors but got one: "+msg+"%v", msg, err)
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
	t.Run("correct", sub(data.Correct, true))
	t.Run("incorrect", sub(data.Incorrect, false))
}
