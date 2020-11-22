package pretty

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/fatih/color"

	"github.com/govitia/navitia/types"
)

const timeLayout = "15:04"

// JourneyConf stores configuration for pretty-printing a types.Journey.
type JourneyConf struct {
	Departure      *color.Color
	Arrival        *color.Color
	Duration       *color.Color
	DateTimeLayout string

	Section SectionConf
}

// DefaultJourneyConf holds a default, quite good configuration.
var DefaultJourneyConf = JourneyConf{
	Departure:      color.New(color.FgRed),
	Arrival:        color.New(color.FgRed),
	Duration:       color.New(color.FgMagenta),
	DateTimeLayout: timeLayout,
	Section:        DefaultSectionConf,
}

// PrettyWrite writes a pretty-printed types.Journey to out.
func (conf JourneyConf) PrettyWrite(j *types.Journey, out io.Writer) error {
	// Build the envellope
	const msgFmt = "%s ➡️ %s | %s\n"
	msg := fmt.Sprintf(
		msgFmt,
		conf.Departure.Sprint(j.Departure.Format(conf.DateTimeLayout)),
		conf.Arrival.Sprint(j.Arrival.Format(conf.DateTimeLayout)),
		conf.Duration.Sprint(j.Duration.String()),
	)

	// Buffers to line-up the reads, sequentially
	buffers := make([]io.Reader, len(j.Sections))

	// Waitgroup so that we wait until all goroutines have completd
	wg := sync.WaitGroup{}

	// Iterate through the journeys, printing them
	for i, s := range j.Sections {
		buf := &bytes.Buffer{}
		buffers[i] = buf

		// Increment the WaitGroup
		wg.Add(1)

		// Launch !
		go func(s types.Section) {
			defer wg.Done()
			err := conf.Section.PrettyWrite(&s, buf)

			// TODO: Deal with errors
			_ = err
		}(s)
	}

	// Create the reader
	readers := append([]io.Reader{strings.NewReader(msg)}, buffers...)
	reader := io.MultiReader(readers...)

	// Wait for completion
	wg.Wait()

	// Copy the new reader to the given output
	_, err := io.Copy(out, reader)
	return err
}
