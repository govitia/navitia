package navitia

import (
	"net/http"
	"path"
	"time"

	"github.com/govitia/navitia/types"
)

const (
	// DefaultAPIProtocol is the protocol to be used
	DefaultAPIProtocol = "https"

	// DefaultAPIHostname is the known Navitia API hostname
	DefaultAPIHostname = "api.navitia.io"

	// DefaultAPIVersion is the used API Version
	DefaultAPIVersion = "v1"

	defaultAPIURL = DefaultAPIProtocol + "://" + DefaultAPIHostname + "/" + DefaultAPIVersion

	// Maximum size of response in bytes
	// 10 megabytes
	maxSize int64 = 10 * (1000 * 1000)
)

var defaultClient = &http.Client{}

// Session holds a current session, it is thread-safe
type Session struct {
	APIKey string
	APIURL string

	client  *http.Client
	created time.Time
}

// New creates a new session given an API Key.
// It acts as a convenience wrapper to NewCustom.
//
// Warning: No Timeout is indicated in the default http client, and as such, it is strongly advised to use NewCustom with a custom *http.Client !
func New(key string) (*Session, error) {
	return NewCustom(key, path.Clean(defaultAPIURL), defaultClient)
}

// NewCustom creates a custom new session given an API key, URL to api base & http client
func NewCustom(key, url string, client *http.Client) (*Session, error) {
	return &Session{
		APIKey:  key,
		APIURL:  url,
		created: time.Now(),
		client:  client,
	}, nil
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
