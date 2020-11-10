package types

import (
	"time"

	"github.com/twpayne/go-geom"
)

// A Region holds information about a geographical region, including its ID, name & shape.
type Region struct {
	// Identifier of the region
	ID ID `json:"id"`

	// Name of the region
	Name string `json:"name"`

	// Status of the dataset
	Status string `json:"status"`

	// Shape of the region.
	// You can use it to check if a particular coordinate is within that MultiPolygon
	Shape *geom.MultiPolygon `json:"shape"`

	// When was the DataSet created ?
	DatasetCreation time.Time `json:"dataset_creation"`
	// When was it last loaded at navitia.io's end ?
	LastLoaded time.Time `json:"last_loaded"`

	// When did production start ?
	ProductionStart time.Time `json:"production_start"`
	// When did or when will it stop ?
	ProductionEnd time.Time `json:"production_end"`

	// An error in the dataset.
	// This comes from the server, not from this package.
	Error string `json:"error"`
}
