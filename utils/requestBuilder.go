package utils

import (
	"net/url"
	"strconv"
	"time"

	"github.com/govitia/navitia/types"
)

// RequestBuilder define a basic http request params builder for Navitia.
type RequestBuilder struct {
	params *url.Values
}

// NewRequestBuilder create and return a new instance of RequestBuilder
func NewRequestBuilder() RequestBuilder {
	return RequestBuilder{params: &url.Values{}}
}

// AddUInt add an unigned integer to the request.
func (rb RequestBuilder) AddUInt(key string, amount uint) {
	rb.params.Add(key, strconv.FormatUint(uint64(amount), 10))
}

// AddUInt add a signed integer to the request.
func (rb RequestBuilder) AddInt(key string, amount int) {
	rb.params.Add(key, strconv.FormatInt(int64(amount), 10))
}

// AddUInt add a floating point number to the request.
func (rb RequestBuilder) AddFloat64(key string, amount float64) {
	rb.params.Add(key, strconv.FormatFloat(amount, 'f', 3, 64))
}

// AddString add a string to the request.
func (rb RequestBuilder) AddString(key string, value string) {
	if value != "" {
		rb.params.Add(key, value)
	}
}

// AddString add a string slice to the request. Add nothing if len of the given slice is equal to 0.
func (rb RequestBuilder) AddStringSlice(key string, values []string) {
	if len(values) != 0 {
		for _, val := range values {
			rb.params.Add(key, val)
		}
	}
}

// AddIDSlice add an ID slice to the request.
func (rb RequestBuilder) AddIDSlice(key string, ids []types.ID) {
	if len(ids) != 0 {
		for _, id := range ids {
			rb.params.Add(key, string(id))
		}
	}
}

// AddMode add a mode list to the request
func (rb RequestBuilder) AddMode(key string, modes []string) {
	if len(modes) != 0 {
		for _, mode := range modes {
			rb.params.Add(key, mode)
		}
	}
}

// AddDate add a date time to the request (YYYYMMDDThhmmss)
func (rb RequestBuilder) AddDateTime(key string, date time.Time) {
	if !date.IsZero() {
		rb.params.Add(key, date.Format(types.DateTimeFormat))
	}
}

// Values return value of url.Values
func (rb RequestBuilder) Values() url.Values {
	return *rb.params
}