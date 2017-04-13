// +build gofuzz

package types

func FuzzJourney(data []byte) int {
	var j = &Journey{}

	// Let's unmarshal
	err := j.UnmarshalJSON(data)
	if err != nil {
		return 0
	}

	// Now that it is unmarshalled, let's the string method !
	_ = j.String()

	return 1
}
