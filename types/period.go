package types

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// Period of effect
type Period struct {
	Begin time.Time `json:"begin"`
	End   time.Time `json:"end"`
}

// UnmarshalJSON implements json.Unmarshaller for a Period
func (p *Period) UnmarshalJSON(b []byte) error {
	// First let's create the analogous structure
	// We define some of the value as pointers to the real values, allowing us to bypass copying in cases where we don't need to process the data
	data := &struct {
		// Those we will process
		Begin string `json:"begin"`
		End   string `json:"end"`
	}{}

	// Let's create the error generator
	gen := unmarshalErrorMaker{"Period", b}

	// Now unmarshall the raw data into the analogous structure
	err := json.Unmarshal(b, data)
	if err != nil {
		return errors.Wrap(err, "Error while unmarshalling Period")
	}

	// Now we process the times.
	p.Begin, err = parseDateTime(data.Begin)
	if err != nil {
		return gen.err(err, "Begin", "begin", data.Begin, "parseDateTime failed")
	}
	p.End, err = parseDateTime(data.End)
	if err != nil {
		return gen.err(err, "End", "end", data.End, "parseDateTime failed")
	}

	// Finished !
	return nil
}
