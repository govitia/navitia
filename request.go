package navitia

import (
	"net/url"
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
