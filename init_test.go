package navitia

import (
	"flag"
	"net/http"
)

const skipNoKey = "No api key supplied, skipping (provide one using -key flag)"

var (
	apiKey      = flag.String("key", "", "API Key to use for testing")
	testSession *Session
)

// Initialise testing function.
func init() {
	// Populate flags
	flag.Parse()

	// Create session
	if *apiKey != "" {
		var err error
		testSession, err = NewCustom(*apiKey, "http://api.navitia.io/v1", http.DefaultClient)
		if err != nil {
			panic(err)
		}
	}
}
