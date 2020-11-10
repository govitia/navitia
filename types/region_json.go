package types

import (
	"encoding/json"

	"github.com/mb0/wkt"
	"github.com/pkg/errors"
	"github.com/twpayne/go-geom"
)

// UnmarshalJSON implements json.Unmarshaller for a Region
func (r *Region) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	data := &struct {
		ID     *ID     `json:"id"`
		Name   *string `json:"name"`
		Status *string `json:"status"`

		// This is mind-fuckery of the highest level.
		// While EVERY other geojson value returned by navitia is in standard format, THIS ONE, for NO GOOD REASON is coded in wkt...
		// See (http://en.wikipedia.org/wiki/Well-known_text).
		Shape string `json:"shape"`

		DatasetCreation string `json:"dataset_created_at"`
		LastLoaded      string `json:"last_load_at"`

		ProductionStart string `json:"start_production_date"`
		ProductionEnd   string `json:"end_production_date"`

		Error *string `json:"error"`
	}{
		ID:     &r.ID,
		Name:   &r.Name,
		Status: &r.Status,
		Error:  &r.Error,
	}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling journey")
	}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Region", b}

	// Now let's process the values
	// First the times
	r.DatasetCreation, err = parseDateTime(data.DatasetCreation)
	if err != nil {
		return gen.err(err, "DatasetCreation", "dataset_created_at", data.DatasetCreation, "Error while parsing datetime")
	}
	r.LastLoaded, err = parseDateTime(data.LastLoaded)
	if err != nil {
		return gen.err(err, "LastLoaded", "last_load_at", data.LastLoaded, "Error while parsing datetime")
	}
	r.ProductionStart, err = parseDateTime(data.ProductionStart)
	if err != nil {
		return gen.err(err, "ProductionStart", "start_production_date", data.ProductionStart, "Error while parsing datetime")
	}
	r.ProductionEnd, err = parseDateTime(data.ProductionEnd)
	if err != nil {
		return gen.err(err, "ProductionEnd", "end_production_date", data.ProductionEnd, "Error while parsing datetime")
	}

	// And now let's have some FUN, deal with the "shape" key.
	// First, let's check if the string isn't empty, cause that would be so awesome...
	if data.Shape != "" {
		// Parse the MKT
		out, err := wkt.Parse([]byte(data.Shape))
		if err != nil {
			return gen.err(err, "Shape", "shape", out, "error in wkt.Parse")
		}

		// Now, out should be a wkt.MultiPolygon
		wktmp, ok := out.(*wkt.MultiPolygon)
		if !ok {
			return gen.err(nil, "Shape", "shape", out, "expected out to be of type wkt.MultiPolygon, but it isn't !")
		}

		// Call our funny little function to convert that to a geom format
		mp, err := convertWktMPtoGeomMP(wktmp)
		if err != nil {
			return gen.err(err, "Shape", "shape", wktmp, "error while converting *wkt.MultiPolygon to *geom.MultiPolygon via convertWktMPtoGeomMP")
		}

		r.Shape = mp
	}

	return nil
}

// convertWktMPtoGeomMP converts a wkt MultiPolygon to a geom MultiPolygon
func convertWktMPtoGeomMP(in *wkt.MultiPolygon) (*geom.MultiPolygon, error) {
	// Now let's convert it to a geom format
	// First let's create the geom.MultiPolygon
	mp := geom.NewMultiPolygon(geom.XY)

	// Then let's iterate through the polygons, and convert each of them from wkt.Coord to geom.Coord
	multipolygonCoords := make([][][]geom.Coord, len(in.Polygons))
	for i, k := range in.Polygons {
		polygonCoords := make([][]geom.Coord, len(k))
		for j, l := range k {
			coords := make([]geom.Coord, len(l))
			for n, m := range l {
				coord := make(geom.Coord, 2)
				coord[0] = m.X
				coord[1] = m.Y
				coords[n] = coord
			}
			polygonCoords[j] = coords
		}
		multipolygonCoords[i] = polygonCoords
	}

	// Now assign it !
	mp, err := mp.SetCoords(multipolygonCoords)
	if err != nil {
		return mp, errors.Wrapf(err, "Error while setting coordinates")
	}
	return mp, err
}
