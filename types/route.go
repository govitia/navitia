package types

import (
	"encoding/json"
	"fmt"
)

// A Route represents a route: a Line can have several routes,
// that is several directions with potential junctions and different frequency for each.
// See http://doc.navitia.io/#public-transport-objects.
type Route struct {
	ID            ID             `json:"id"`             // Identifier of the route, eg: "route:RAT:M6"
	Name          string         `json:"name"`           // Name of the route
	Frequence     bool           `json:"is_frequence"`   // If the route has frequency or not. Can only be “False”, but may be “True” in the future
	Line          Line           `json:"line"`           // Line is the line it is connected to
	Direction     Container      `json:"direction"`      // Direction is the direction of the route (Place or POI)
	PhysicalModes []PhysicalMode `json:"physical_modes"` // PhysicalModes of the line
	GeoJSON       GeoJSON        `json:"geo_json"`
}

// jsonRoute define the JSON implementation of Route struct
// We define some of the value as pointers to the real values,
// allowing us to bypass copying in cases where we don't need to process the data.
type jsonRoute struct {
	ID        *ID        `json:"id"`
	Name      *string    `json:"name"`
	Line      *Line      `json:"line"`
	Direction *Container `json:"direction"`

	// Value to process
	Frequence string `json:"is_frequence"`
}

type GeoJSON struct {
	Type string `json:"type"`
}

// UnmarshalJSON implements json.Unmarshaller for Route
func (r *Route) UnmarshalJSON(b []byte) error {
	data := &jsonRoute{
		ID:        &r.ID,
		Name:      &r.Name,
		Line:      &r.Line,
		Direction: &r.Direction,
	}

	// Create the error generator
	gen := unmarshalErrorMaker{"Route", b}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return fmt.Errorf("error while unmarshalling Line")
	}

	// Now process the value
	switch {
	case data.Frequence == "true" || data.Frequence == "True":
		r.Frequence = true
	case data.Frequence == "false" || data.Frequence == "False":
		r.Frequence = false
	default:
		return gen.err(nil, "Frequence", "is_frequency", data.Frequence, `String is neither True, true, False or false`)
	}

	return nil
}
