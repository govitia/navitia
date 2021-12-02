package types

type StopArea struct {
	ID                    string                  `json:"id"`
	Name                  string                  `json:"name"`
	Label                 string                  `json:"label"`
	Coord                 Coord                   `json:"coord"`
	AdministrativeRegions []*AdministrativeRegion `json:"administrative_regions"`
	StopPoints            []*StopPoint            `json:"stop_points"`
}
