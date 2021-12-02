package types

type Address struct {
	ID          string `json:"id"`   // Identifier of the address
	Name        string `json:"name"` // Name of the address
	Label       string `json:"label"`
	Coord       Coord  `json:"coord"`        // Coordinates of the address
	HouseNumber int    `json:"house_number"` // House number of the address
	// Administrative regions of the address in which is the stop area
	AdministrativeRegions []*AdministrativeRegion `json:"administrative_regions"`
}
