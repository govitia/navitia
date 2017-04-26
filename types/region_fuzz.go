// +build gofuzz

package types

func FuzzRegion(data []byte) int {
	var r = &Region{}

	// Let's unmarshal
	err := r.UnmarshalJSON(data)
	if err != nil {
		return 0
	}

	// Now that it is unmarshalled, let's the string method !
	_ = r.String()

	return 1
}
