package types

import (
	"fmt"
	"time"

	"github.com/twpayne/go-geom"
)

// A Region holds information about a geographical region, including its ID, name & shape
type Region struct {
	// Identifier of the region
	ID ID

	// Name of the region
	Name string

	// Status of the dataset
	Status string

	// Shape of the region.
	// You can use it to check if a particular coordinate is within that MultiPolygon
	Shape *geom.MultiPolygon

	// When was the DataSet created ?
	DatasetCreation time.Time
	// When was it last loaded at navitia.io's end ?
	LastLoaded time.Time

	// When did production start ?
	ProductionStart time.Time
	// When did or when will it stop ?
	ProductionEnd time.Time

	// An error in the dataset.
	// This comes from the server, not from this package.
	Error string
}

// String stringifies a region
func (r Region) String() string {
	format := `ID: %s
Name: %s
Status: %s
Error: %v
`
	return fmt.Sprintf(format, r.ID, r.Name, r.Status, r.Error)
}
