package types

import (
	"fmt"
	//"github.com/paulmach/go.geojson"
	"time"
)

// A Region holds information about a geographical region, including its ID, name & shape
type Region struct {
	ID     ID
	Name   string
	Status string

	Shape string // Temporary: Should be a geojson geometry object, why isn't it ?

	DatasetCreation time.Time
	LastLoaded      time.Time

	ProductionStart time.Time
	ProductionEnd   time.Time

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
