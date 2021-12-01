package types

// Place is a container containing either an admin, poi, address, stop_area, stop_point
type Place struct {
	ID                   string               `json:"id"`                    // The id of the embedded object
	Name                 string               `json:"name"`                  // The name of the embedded object
	Quality              int                  `json:"quality"`               // The quality of the place
	EmbeddedType         PlaceEmbeddedType    `json:"embedded_type"`         // The type of the embedded object
	AdministrativeRegion AdministrativeRegion `json:"administrative_region"` // Embedded administrative region
	StopArea             StopArea             `json:"stop_area"`             // Embedded Stop area
	Poi                  Poi                  `json:"poi"`                   // Embedded poi
	Address              Address              `json:"address"`               // Embedded address
	StopPoint            StopPoint            `json:"stop_point"`            // Embedded stop point
}
