package types

// Coord defines a coordinate struct with latitude and longitude.
// Lots of object are geographically localized.
type Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

