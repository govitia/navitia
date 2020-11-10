package pretty

import (
	"bytes"
	"io"
	"sync"

	"github.com/fatih/color"

	"github.com/govitia/navitia"
	"github.com/govitia/navitia/types"
)

// JourneyResultsConf stores configuration for pretty-printing a navitia.JourneyResults
type JourneyResultsConf struct {
	Count   *color.Color
	Journey JourneyConf
}

// DefaultJourneyResultsConf holds a default, quite good configuration
var DefaultJourneyResultsConf = JourneyResultsConf{
	Count:   color.New(color.FgBlack),
	Journey: DefaultJourneyConf,
}

// PrettyWrite writes a pretty-printed navitia.JourneyResults to out
func (conf JourneyResultsConf) PrettyWrite(jr *navitia.JourneyResults, out io.Writer) error {
	// Buffers to line-up the reads, sequentially
	buffers := make([]io.Reader, jr.Count())
	// Waitgroup for each goroutine
	wg := sync.WaitGroup{}

	// Iterate through the journeys, printing them
	for i, j := range jr.Journeys {
		buf := &bytes.Buffer{}
		buffers[i] = buf

		// Increment the WaitGroup
		wg.Add(1)

		// Launch !
		go func(j types.Journey) {
			defer wg.Done()
			err := conf.Journey.PrettyWrite(&j, buf)

			// TODO: Deal with the errors
			_ = err
		}(j)
	}

	// Create the reader
	reader := io.MultiReader(buffers...)

	// Wait for completion
	wg.Wait()

	// Copy the new reader to the given output
	_, err := io.Copy(out, reader)
	return err
}
