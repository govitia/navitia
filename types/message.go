package types

// A Message contains the text to be provided to the traveler.
type Message struct {
	Text    string   `json:"text"`    // The message to bring to the traveler
	Channel *Channel `json:"channel"` // The destination media for this Message.
}
