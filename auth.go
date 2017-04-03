package gonavitia

import (
	"github.com/pkg/errors"
	"net/http"
	"time"
)

const (
	// APIProtocol is the protocol to be used
	DefaultAPIProtocol = "https"

	// APIHostname is the known Navitia API hostname
	DefaultAPIHostname = "api.navitia.io"

	// APIVersion is the used API Version
	DefaultAPIVersion = "v1"

	defaultAPIURL = DefaultAPIProtocol + "://" + DefaultAPIHostname + "/" + DefaultAPIVersion
)

var defaultClient = &http.Client{}

// Session holds a current session, it is thread-safe
type Session struct {
	APIKey string
	APIURL string

	client  *http.Client
	created time.Time
}

// New creates a new session given an API Key
// It acts as a convenience wrapper to NewCustom
func New(key string) (*Session, error) {
	return NewCustom(key, defaultAPIURL, defaultClient)
}

// NewCustom creates a custom new session given an API key, URL to api base & http client
func NewCustom(key string, url string, client *http.Client) (*Session, error) {
	return &Session{
		APIKey:  key,
		APIURL:  url,
		created: time.Now(),
		client:  client,
	}, nil
}

// UseClient sets a given *http.Client to be used for further queries
func (s *Session) UseClient(client *http.Client) {
	s.client = client
}

// newRequest creates a newRequest with the correct auth set
func (s *Session) newRequest(url string) (*http.Request, error) {
	// Create the request
	newReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return newReq, errors.Wrapf(err, "couldn't create new request (for %s)", url)
	}

	// Add basic auth
	newReq.SetBasicAuth(s.APIKey, "")

	return newReq, err
}
