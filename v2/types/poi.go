package types

// PoiType
// /poi_types lists groups of point of interest.
// You will find classifications as theater, offices or bicycle rental station for example.
type PoiType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Poi = Point Of Interest
type Poi struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Label  string  `json:"label"`
	Type   PoiType `json:"type"`
	Stands Stands  `json:"stands"`
}
