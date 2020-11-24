package types

// A Channel is a destination media for a message.
type Channel struct {
	ID          ID       `json:"id"`              // ID of the address
	ContentType string   `json:"content_type"`    // Content Type (text/html etc.) RFC1341.4
	Name        string   `json:"name"`            // Name of the channel
	Types       []string `json:"types,omitempty"` // Types ?
}
