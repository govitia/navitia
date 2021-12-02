package types

// PoiType
// /poi_types lists groups of point of interest.
// You will find classifications as theater, offices or bicycle rental station for example.
type PoiType struct {
	ID   string `json:"id"`   // Identifier of the poi type
	Name string `json:"name"` // Name of the poi type
}

// Poi = Point Of Interest
type Poi struct {
	ID     string  `json:"id"`   // Identifier of the poi
	Name   string  `json:"name"` // Name of the poi
	Label  string  `json:"label"`
	Type   PoiType `json:"type"`   // Type of the poi
	Stands Stands  `json:"stands"` // Information on the spots available, for BSS stations

}
