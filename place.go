package gonavitia

import "./types"

// TODO: Rename Place to PlaceResults as a way to parse the results

type internalPlace struct {
	ID           types.PlaceID `json:"id"`
	Name         string        `json:"name"`
	Quality      uint          `json:"quality,omitempty"`
	EmbeddedType string        `json:"embedded_type"`

	// Four possibilities
	StopArea             types.StopArea             `json:"stop_area,omitempty"`
	POI                  types.POI                  `json:"POI,omitempty"`
	Address              types.Address              `json:"address,omitempty"`
	StopPoint            types.StopPoint            `json:"stop_point,omitempty"`
	AdministrativeRegion types.AdministrativeRegion `json:"administrative_region"`
}
