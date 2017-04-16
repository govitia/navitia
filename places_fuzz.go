// +build gofuzz

package navitia

func FuzzPlaces(data []byte) int {
	var pr = &PlacesResults{}

	// Let's unmarshal
	err := pr.UnmarshalJSON(data)
	if err != nil {
		return 0
	}

	// Now that it is unmarshalled, let's the string method !
	_ = pr.String()

	return 1
}
