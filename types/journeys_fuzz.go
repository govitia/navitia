// +build gofuzz

package types

func FuzzJourney(data []byte) int {
	j := &Journey{}

	// Let's unmarshal
	err := j.UnmarshalJSON(data)
	if err != nil {
		return 0
	}

	return 1
}
