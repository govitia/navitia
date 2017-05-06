package navitia

import "time"

// Logging stores logging info
type Logging struct {
	Created  time.Time
	Sent     time.Time
	Received time.Time
}

// creating stores creation time
func (l *Logging) creating() {
	l.Created = time.Now()
}

// sending stores sending time
func (l *Logging) sending() {
	l.Sent = time.Now()
}

// parsing stores parsing time
func (l *Logging) parsing() {
	l.Received = time.Now()
}

// Waiting returns the time spent waiting for a response from the server
func (l *Logging) Waiting() time.Duration {
	return l.Received.Sub(l.Sent)
}
