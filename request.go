package navitia

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
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
func (s *Session) request(ctx context.Context, url string, query query, results results) error {
	// Store creation time
	results.creating()

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.Wrapf(err, "couldn't create new request (for %s)", url)
	}

	// Add context to the request
	req = req.WithContext(ctx)

	// Add basic auth
	req.SetBasicAuth(s.APIKey, "")

	// Encode the parameters
	values, err := query.toURL()
	if err != nil {
		return errors.Wrap(err, "error while retrieving url values to be encoded")
	}
	req.URL.RawQuery = values.Encode()

	// Execute the request
	resp, err := s.client.Do(req)
	results.sending()

	// Check the response
	if err != nil {
		return errors.Wrap(err, "error while executing request")
	}
	if resp.StatusCode != 200 {
		return parseRemoteError(resp)
	}

	// Check for cancellation
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	// Limit the reader
	reader := io.LimitReader(resp.Body, maxSize)

	// Parse the now limited body
	dec := json.NewDecoder(reader)
	err = dec.Decode(results)
	if err != nil {
		return errors.Wrap(err, "JSON decoding failed")
	}
	results.parsing()

	// Return
	return err
}
