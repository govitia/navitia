package pretty

import (
	"fmt"
	"io"

	"github.com/aabizri/navitia/types"
	"github.com/fatih/color"
)

var modeEmoji = map[string]string{
	string(types.PhysicalModeAir):               "✈️",
	string(types.PhysicalModeBoat):              "⛴️",
	string(types.PhysicalModeBus):               "🚍",
	string(types.PhysicalModeBusRapidTransit):   "🚍",
	string(types.PhysicalModeCoach):             "🚍",
	string(types.PhysicalModeFerry):             "⛴️",
	string(types.PhysicalModeFunicular):         "🚞",
	string(types.PhysicalModeLocalTrain):        "🚆",
	string(types.PhysicalModeLongDistanceTrain): "🚆",
	string(types.PhysicalModeMetro):             "🚇",
	string(types.PhysicalModeRapidTransit):      "🚍",
	string(types.PhysicalModeShuttle):           "🚐",
	string(types.PhysicalModeTaxi):              "🚖",
	string(types.PhysicalModeTrain):             "🚆",
	string(types.PhysicalModeTramway):           "🚊",

	// Because the API doesn't always return predictable returns, we have aliases
	"Métro": "🚇",
	"Bus":   "🚍",

	// Classic Modes: Walking, biking or bikesharing
	string(types.ModeWalking):   "🚶",
	string(types.ModeBike):      "🚴",
	string(types.ModeBikeShare): "🚴",
}

// SectionConf stores configuration for pretty-printing a types.Section
type SectionConf struct {
	Mode     *color.Color
	Duration *color.Color
	From     *color.Color
	To       *color.Color
	Emoji    bool
}

// DefaultSectionConf holds a default, quite good configuration
var DefaultSectionConf = SectionConf{
	Mode:     color.New(color.FgGreen),
	Duration: color.New(color.FgMagenta),
	From:     color.New(color.FgBlue),
	To:       color.New(color.FgBlue),
}

// PrettyWrite writes a pretty-printed types.Section to out
func (conf SectionConf) PrettyWrite(s *types.Section, out io.Writer) error {
	// if there's no from or no to, finish now
	if s.From.Name == "" || s.To.Name == "" {
		return nil
	}

	var middle string
	switch {
	case s.Mode != "":
		middle = modeEmoji[string(s.Mode)]
	case s.Display.PhysicalMode != "":
		middle = modeEmoji[string(s.Display.PhysicalMode)] + s.Display.Label
	}
	msg := fmt.Sprintf("\t%s (%s)\t%s➡️%s\n", conf.Mode.Sprint(middle), conf.Duration.Sprint(s.Duration.String()), conf.From.Sprint(s.From.Name), conf.To.Sprint(s.To.Name))

	_, err := out.Write([]byte(msg))
	return err
}
