package pretty

import (
	"fmt"
	"io"

	"github.com/aabizri/navitia/types"
	"github.com/fatih/color"
)

// RegionConf stores configuration for pretty-printing
type RegionConf struct {
	Name           *color.Color
	ID             *color.Color
	Start          *color.Color
	End            *color.Color
	DateTimeLayout string
}

// DefaultRegionConf holds a default, quite good configuration
var DefaultRegionConf = RegionConf{
	Name:           color.New(color.FgBlue),
	ID:             color.New(),
	Start:          color.New(),
	End:            color.New(),
	DateTimeLayout: defaultDateTimeLayout,
}

const defaultDateTimeLayout = "02/01/2006"

// PrettyWrite writes a pretty-printed account of a navitia.PlacesResults to out.
func (conf RegionConf) PrettyWrite(r *types.Region, out io.Writer) error {
	format := "%s (id: %s) [%s-%s]"
	msg := fmt.Sprintf(format, conf.Name.Sprint(r.Name), conf.ID.Sprint(r.ID), conf.Start.Sprint(r.ProductionStart.Format(conf.DateTimeLayout)), conf.End.Sprint(r.ProductionEnd.Format(conf.DateTimeLayout)))
	out.Write([]byte(msg))

	return nil
}
