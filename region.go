package types

import "fmt"

// A Region holds information about a geographical region, including its ID, name & shape
type Region struct {
	ID     ID     `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`

	//DatasetCreation time.Time `json:"dataset_created_at"`
	//LastLoaded      time.Time `json:"last_load_at"`

	//ProductionStart time.Time `json:"start_production_date"`
	//ProductionEnd   time.Time `json:"end_production_date"`

	Error error `json:"error"`
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
