package types

type StopPoint struct {
	ID                    string                  `json:"id"`
	Name                  string                  `json:"name"`
	Coord                 Coord                   `json:"coord"`
	AdministrativeRegions []*AdministrativeRegion `json:"administrative_regions"`
	Equipments            []Equipment             `json:"equipments"`
	StopArea              *StopArea               `json:"stop_areas"`
}
