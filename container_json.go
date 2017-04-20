package types

import (
	"encoding/json"
	"sync"

	"github.com/pkg/errors"
)

// UnmarshalJSON satisfies the json.Unmarshaller interface
func (c *Container) UnmarshalJSON(b []byte) error {
	// Set up a mutex
	c.mu = &sync.RWMutex{}

	// Unmarshal into a map
	data := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return errors.Wrap(err, "Couldn't unmarshal into a map")
	}

	// From a map, extract the ID, Name & EmbeddedType
	if id, ok := data["id"]; ok {
		err := json.Unmarshal(id, &c.ID)
		if err != nil {
			return err
		}
	}
	if name, ok := data["name"]; ok {
		err := json.Unmarshal(name, &c.Name)
		if err != nil {
			return err
		}
	}
	if embeddedType, ok := data["embedded_type"]; ok {
		err := json.Unmarshal(embeddedType, &c.EmbeddedType)
		if err != nil {
			return err
		}
	}
	if quality, ok := data["quality"]; ok {
		err := json.Unmarshal(quality, &c.Quality)
		if err != nil {
			return err
		}
	}
	// Now, unmarshal the last one
	if embedded, ok := data[c.EmbeddedType]; ok {
		c.embeddedJSON = embedded
	}

	return nil
}
