package navitia

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
)

// Paging holds potential Previous / Next functions
type Paging struct {
	// Next results
	Next func(ctx context.Context, s *Session, res results) error

	// Previous results
	Previous func(ctx context.Context, s *Session, res results) error
}

// createPagingFunc creates a paging func (either Previous or Next)
func createPagingFunc(url string) func(ctx context.Context, s *Session, res results) error {
	f := func(ctx context.Context, s *Session, res results) error {
		return s.requestURL(ctx, url, res)
	}
	return f
}

// UnmarshalJSON unmarshals a Paging type from a Links data structure
func (p *Paging) UnmarshalJSON(b []byte) error {
	var links []link
	err := json.Unmarshal(b, &links)
	if err != nil {
		return errors.Wrap(err, "error while unmarshalling links")
	}

	// Iterate through the links
	for _, l := range links {
		switch l.Type {
		case "next":
			p.Next = createPagingFunc(l.Href)
		case "previous":
			p.Previous = createPagingFunc(l.Href)
		}
	}

	// Return
	return nil
}

type link struct {
	Href      string
	Rel       string
	Templated bool
	Type      string
}
