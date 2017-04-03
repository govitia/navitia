package gonavitia

import (
	"flag"
)

const skipNoKey string = "No api key supplied, skipping (provide one using -key flag)"

var (
	apiKey      = flag.String("key", "", "API Key to use for testing")
	testSession *Session
)

// Initialize testing function
func init() {
	// Populate flags
	flag.Parse()

	// Create session
	if *apiKey != "" {
		var err error
		testSession, err = New(*apiKey)
		if err != nil {
			panic(err)
		}
	}
}
