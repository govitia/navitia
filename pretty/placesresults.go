package pretty

import (
	"bytes"
	"io"
	"strings"
	"sync"

	"github.com/fatih/color"

	"github.com/govitia/navitia"
	"github.com/govitia/navitia/types"
)

// PlacesResultsConf stores configuration for pretty-printing
type PlacesResultsConf struct {
	Count *color.Color
	Place ContainerConf
}

// DefaultPlacesResultsConf holds a default, quite good configuration
var DefaultPlacesResultsConf = PlacesResultsConf{
	Count: color.New(color.Italic),
	Place: DefaultContainerConf,
}

// PrettyWrite writes a pretty-printed account of a navitia.PlacesResults to out.
func (conf PlacesResultsConf) PrettyWrite(pr *navitia.PlacesResults, out io.Writer) error {
	// Buffers to line-up the reads, sequentially
	buffers := make([]io.Reader, pr.Len())

	// Waitgroup for each goroutine
	wg := sync.WaitGroup{}

	// Iterate through the places, printing them
	for i, p := range pr.Places {
		base := []byte(color.New(color.FgCyan).Sprintf("#%d: ", i))
		buf := bytes.NewBuffer(base)
		buffers[i] = buf

		// Increment the WaitGroup
		wg.Add(1)

		// Launch !
		go func(p types.Container) {
			defer wg.Done()

			if err := conf.Place.ContainerWrite(&p, buf); err != nil {
				panic(err)
			}

			_, err := buf.WriteString("\n")
			panic(err)
		}(p)
	}

	// Create the overall message
	msg := conf.Count.Sprintf("(%d places found)\n", pr.Len())

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
