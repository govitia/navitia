package navitia

import (
	"net/url"

	"github.com/govitia/navitia/types"
	"github.com/govitia/navitia/utils"
)

type DeparturesResults struct {
	Departures []types.Departure `json:"departures"`
	Paging     Paging            `json:"links"`
	Logging    `json:"-"`
	session    *Session
}

// Count returns the number of results available in a Departures
func (dr *DeparturesResults) Count() int {
	return len(dr.Departures)
}

// DeparturesRequest contain the parameters needed to make a departures
type DeparturesRequest struct {
	StopArea string
}

func (req DeparturesRequest) toURL() (url.Values, error) {
	rb := utils.NewRequestBuilder()

	rb.AddString("stop_area", req.StopArea)

	return rb.Values(), nil
}
