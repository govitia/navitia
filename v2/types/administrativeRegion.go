package types

// AdministrativeRegion
// Cities are mainly on the 8 level, dependent on the country
// (http://wiki.openstreetmap.org/wiki/Tag:boundary%3Dadministrative)
type AdministrativeRegion struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Label of the administrative region.
	// The name is directly taken from the data whereas
	// the label is something we compute for better traveler information.
	// If you don't know what to display, display the label.
	Label   string `json:"label"`
	Coord   Coord  `json:"coord"`
	Level   int    `json:"level"`
	ZipCode string `json:"zip_code"`
}
