package types

type PtObjectType string

const (
	POTNetwork        PtObjectType = "network"
	POTCommercialMode PtObjectType = "commercial_mode"
	POTLine           PtObjectType = "line"
	POTRoute          PtObjectType = "route"
	POTStopPoint      PtObjectType = "stop_point"
	POTStopArea       PtObjectType = "stop_area"
	POTTrip           PtObjectType = "trip"
)

// PtObject is a container containing either a network,
// commercial_mode, line, route, stop_point, stop_area, trip
type PtObject struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Quality        int             `json:"quality"`
	EmbeddedType   PtObjectType    `json:"embedded_type"`
	StopArea       *StopArea       `json:"stop_area,omitempty"`
	StopPoint      *StopPoint      `json:"stop_point,omitempty"`
	Network        *Network        `json:"network,omitempty"`
	CommercialMode *CommercialMode `json:"commercial_mode,omitempty"`
	Line           *Line           `json:"line,omitempty"`
	Route          *Route          `json:"route,omitempty"`
	Trip           *Trip           `json:"trip,omitempty"`
}
