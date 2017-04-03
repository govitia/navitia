package gonavitia

type PlaceID string

type Place struct {
	ID           PlaceID `json:"id"`
	Name         string  `json:"name"`
	Quality      uint    `json:"quality,omitempty"`
	EmbeddedType string  `json:"embedded_type"`

	// Four possibilities
	StopArea  StopArea `json:"stop_area,omitempty"`
	POI       POI      `json:"POI,omitempty"`
	Address   Address  `json:"address,omitempty"`
	StopPoint `json:"stop_point,omitempty"`
}

type StopArea struct {
}

type POI struct {
}

type Address struct {
}

type StopPoint struct {
}
