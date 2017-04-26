package pretty

import (
	"fmt"
	"io"

	"github.com/aabizri/navitia/types"
	"github.com/fatih/color"
)

// ContainerConf stores configuration for use in ContainerWrite
type ContainerConf struct {
	Quality *color.Color
	Type    *color.Color
	Name    *color.Color
}

// DefaultContainerConf holds a default, quite good configuration
var DefaultContainerConf = ContainerConf{
	Quality: color.New(color.FgMagenta),
	Type:    color.New(color.FgGreen),
	Name:    color.New(color.FgBlue),
}

var placeTypeToName = map[string]string{
	"address":               "Address",
	"poi":                   "Point Of Interest",
	"stop_area":             "Stop Area",
	"stop_point":            "Stop Point",
	"administrative_region": "Administrative Region",
}

// ContainerWrite writes a pretty-printed account of a types.Container to out.
func (conf ContainerConf) ContainerWrite(c *types.Container, out io.Writer) error {
	var msg string
	if c.Quality != 0 {
		msg = conf.Quality.Sprintf("[%d%%] ", c.Quality)
	}
	msg += fmt.Sprintf("(%s)\t%s", conf.Type.Sprint(placeTypeToName[c.EmbeddedType]), conf.Name.Sprint(c.Name))
	out.Write([]byte(msg))

	return nil
}
