package navitia

import (
	"encoding/json"

	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

// UnmarshalJSON implements unmarshalling for ExploreResults
func (sasr *ExploreResults) UnmarshalJSON(b []byte) error {
	// first will hold the preliminary values
	first := make(map[string]json.RawMessage)

	// Unmarshal to first
	err := json.Unmarshal(b, &first)
	if err != nil {
		return errors.Wrap(err, "error in first-pass unmarshalling")
	}

	// Extract paging information and unmarshal it
	if links, ok := first["links"]; ok {
		err = json.Unmarshal(links, &sasr.Paging)
		if err != nil {
			return errors.Wrap(err, "error while unmarshalling links (paging info)")
		}
		// Remove "links" from the map
		delete(first, "links")
	} else {
		return errors.New("No \"links\" field found in returned data")
	}

	// Create a value
	var (
		recv   interface{}
		second json.RawMessage
	)

	// Switch on it
	for k := range first {
		switch k {
		case CommercialModesSelector:
			recv = []types.CommercialMode{}
		case LinesSelector:
			recv = []types.Line{}
		case NetworksSelector:
			recv = []types.Network{}
		case RoutesSelector:
			recv = []types.Route{}
		case StopAreasSelector:
			recv = []types.StopArea{}
		case StopPointsSelector:
			recv = []types.StopPoint{}
		case PhysicalModesSelector:
			recv = []types.PhysicalMode{}
		case CompaniesSelector:
			recv = []types.Company{}
		/*case VehicleJourneysSelector:
		recv = []types.VehicleJourneys*/
		case DisruptionsSelector:
			recv = []types.Disruption{}
		}

		// If we have found something, let's break
		if recv != nil {
			// Assign the raw json value to second
			second = first[k]
			break
		}
	}

	// If we have found nothing, return an error
	if recv == nil || second == nil || len(second) == 0 {
		return errors.New("error: no known key in response")
	}

	// Else, let's unmarshal
	err = json.Unmarshal(second, &recv)
	if err != nil {
		return errors.Wrap(err, "error in second pass of json.Unmarshal")
	}

	// Now assign it to sasr.PTObjects
	sasr.Objects = recv

	return nil
}
