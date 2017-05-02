package navitia

import (
	"net/http"
	"time"

	"github.com/aabizri/navitia/types"
	"github.com/pkg/errors"
)

const (
	// NavitiaAPIURL is the navitia.io API URL.
	// It is also the default one when New is used.
	NavitiaAPIURL = "https://api.navitia.io/v1"

	// SNCFAPIURL is the SNCF (French National Railway Company) API URL.
	SNCFAPIURL = "https://api.sncf.com/v1"

	// Maximum size of response in bytes
	// 10 megabytes
	maxSize int64 = 10 * (1000 * 1000)
)

// Session holds a current session, it is thread-safe
type Session struct {
	apiKey string
	apiURL string

	client interface {
		Do(req *http.Request) (*http.Response, error)
	}

	created time.Time
}

// New creates a new session given an API Key.
// It acts as a convenience wrapper to NewCustom.
//
// Warning: No Timeout is indicated in the default http client, and as such, it is strongly advised to use NewCustom with a custom *http.Client !
func New(key string) (*Session, error) {
	return NewCustom(key, http.DefaultClient)
}

// SetAPIURL sets an APIURL in a session
func SetAPIURL(APIURL string) func(*Session) error {
	return func(s *Session) error {
		if s == nil {
			return errors.New("nil session")
		}
		s.apiURL = APIURL
		return nil
	}
}

// NewCustom creates a custom new session given an API key, URL to api base & http client.
// It can also be given additional configuration functions.
func NewCustom(key string, client *http.Client, options ...func(*Session) error) (*Session, error) {
	// Establish the basic value
	s := &Session{
		apiKey:  key,
		created: time.Now(),
		client:  client,
	}

	// Iterate through options
	for i, f := range options {
		err := f(s)
		if err != nil {
			return nil, errors.Wrapf(err, "NewCustom: error while parsing option %d", i)
		}
	}

	// Return
	return s, nil
}

// A Scope is a coverage-scoped question, allowing you to query information about a specific region.
//
// It is needed for every non-global request you wish to make, and helps have better results with some global request too!
type Scope struct {
	region  types.ID
	session *Session
}

// Scope creates a coverage-scoped session given a region ID.
func (s *Session) Scope(region types.ID) *Scope {
	return &Scope{region: region, session: s}
}
