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

// results is implemented by every Result type
type results interface {
	creating()
	sending()
	parsing()
}

// requestURL requests a url, with the query already encoded in, and decodes the result in res.
func (s *Session) requestURL(ctx context.Context, url string, res results) error {
	// Store creation time
	res.creating()

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.Wrapf(err, "couldn't create new request (for %s)", url)
	}

	// Add context to the request
	req = req.WithContext(ctx)

	// Add basic auth
	req.SetBasicAuth(s.apiKey, "")

	// Execute the request
	resp, err := s.client.Do(req)
	res.sending()

	// Check the response
	switch {
	case err == context.Canceled:
		return err
	case err != nil:
		return errors.Wrap(err, "error while executing request")
	case resp.StatusCode != 200:
		return parseRemoteError(resp)
	}

	// Defer the close
	defer resp.Body.Close()

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
	err = dec.Decode(res)
	if err != nil {
		return errors.Wrap(err, "JSON decoding failed")
	}
	res.parsing()

	// Return
	return err
}

// request does a request given a url, query and results to populate
func (s *Session) request(ctx context.Context, baseURL string, query query, res results) error {
	// Encode the parameters
	values, err := query.toURL()
	if err != nil {
		return errors.Wrap(err, "error while retrieving url values to be encoded")
	}
	url := baseURL + "?" + values.Encode()

	// Call requestURL
	return s.requestURL(ctx, url, res)
}
