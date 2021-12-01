package types

type Address struct {
	ID                    string                  `json:"id"`
	Name                  string                  `json:"name"`
	Label                 string                  `json:"label"`
	Coord                 Coord                   `json:"coord"`
	HouseNumber           int                     `json:"house_number"`
	AdministrativeRegions []*AdministrativeRegion `json:"administrative_regions"`
}
