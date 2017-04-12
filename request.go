package gonavitia

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/url"
)

type query interface {
	toURL() (url.Values, error)
}

type results interface {
	creating()
	sending()
	parsing()
}

// request does a request given a url, query and results to populate
func (s *Session) request(url string, query query, results results) error {
	// Store creation time
	results.creating()

	// Get the request
	req, err := s.newRequest(url)
	if err != nil {
		return errors.Wrap(err, "error while creating request")
	}

	// Encode the parameters
	values, err := query.toURL()
	if err != nil {
		return errors.Wrap(err, "error while retrieving url values to be encoded")
	}
	req.URL.RawQuery = values.Encode()

	// Execute the request
	resp, err := s.client.Do(req)
	results.sending()

	// Check it
	if err != nil {
		return errors.Wrap(err, "error while executing request")
	}
	if resp.StatusCode != 200 {
		return parseRemoteError(resp)
	}

	// Parse it
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(results)
	if err != nil {
		return errors.Wrap(err, "JSON decoding failed")
	}
	results.parsing()

	// Return
	return err
}
