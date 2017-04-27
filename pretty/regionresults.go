package pretty

import (
	"bytes"
	"io"
	"strings"
	"sync"

	"github.com/aabizri/navitia"
	"github.com/aabizri/navitia/types"
	"github.com/fatih/color"
)

// RegionResultsConf stores configuration for pretty-printing
type RegionResultsConf struct {
	Count  *color.Color
	Region RegionConf
}

// DefaultRegionResultsConf holds a default, quite good configuration
var DefaultRegionResultsConf = RegionResultsConf{
	Count:  color.New(color.Italic),
	Region: DefaultRegionConf,
}

// PrettyWrite writes a pretty-printed account of a navitia.RegionResults to out.
func (conf RegionResultsConf) PrettyWrite(rr *navitia.RegionResults, out io.Writer) error {
	// Buffers to line-up the reads, sequentially
	buffers := make([]io.Reader, len(rr.Regions))

	// Waitgroup for each goroutine
	wg := sync.WaitGroup{}

	// Iterate through the places, printing them
	for i, r := range rr.Regions {
		var base = []byte(color.New(color.FgCyan).Sprintf("#%d: ", i))
		buf := bytes.NewBuffer(base)
		buffers[i] = buf

		// Increment the WaitGroup
		wg.Add(1)

		// Launch !
		go func(r types.Region) {
			defer wg.Done()
			err := conf.Region.PrettyWrite(&r, buf)
			_, err = buf.WriteString("\n")

			// TODO: Deal with errors
			_ = err
		}(r)
	}

	// Create the overall message
	msg := conf.Count.Sprintf("(%d regions found)\n", len(rr.Regions))

	// Create the reader
	readers := append([]io.Reader{strings.NewReader(msg)}, buffers...)
	reader := io.MultiReader(readers...)

	// Wait for completion
	wg.Wait()

	// Copy the new reader to the given output
	_, err := io.Copy(out, reader)

	// End
	return err
}
