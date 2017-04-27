package types

import (
	"reflect"
	"testing"
)

// Test_Section_Unmarshal tests unmarshalling for Section.
// As the unmarshalling is done in-house, this allows us to check that the custom UnmarshalJSON function correctly
//
// This launches both a "correct" and "incorrect" subtest, allowing us to test both cases.
// 	If we expect no errors but we get one, the test fails
//	If we expect an error but we don't get one, the test fails
func Test_Section_Unmarshal(t *testing.T) {
	testUnmarshal(t, testData["section"], reflect.TypeOf(Section{}))
}
