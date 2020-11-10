// +build gofuzz

package types

func FuzzContainer(data []byte) int {
	c := &Container{}

	// Let's unmarshal, this is not our job so "bleh"
	err := c.UnmarshalJSON(data)
	if err != nil {
		return 0
	}

	// Let's check it !
	err = c.Check()
	if err != nil {
		return 0
	}

	return 1
}
