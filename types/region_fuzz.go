// +build gofuzz

package types

func FuzzRegion(data []byte) int {
	r := &Region{}

	// Let's unmarshal
	err := r.UnmarshalJSON(data)
	if err != nil {
		return 0
	}

	return 1
}
