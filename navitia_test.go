package navitia

import (
	"context"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"time"

	"github.com/aabizri/navitia/testutils"
)

var (
	testDataPathFlag = flag.String("path", "./testdata", "Directory of test data")
	testDataPath     string
)

func init() {
	flag.Parse()
	// If the given path is absolute, then use it as-is
	if filepath.IsAbs(*testDataPathFlag) {
		testDataPath = *testDataPathFlag
	} else {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		testDataPath = filepath.Join(wd, *testDataPathFlag)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	defer cancel()

	data, err := testutils.Load(ctx, testDataPath, typesList)
	if err != nil {
		panic(err)
	}

	testData = data
}

// testData stores a map which maps each category to their data
var testData = make(map[string]*testutils.TestData, len(typesList))

// this is the list of potential types
// must be lower case
var typesList = []string{
	"journeys",
	"coverage",
	"places",
	"connections",
}

// A mockClient implements the correct interface for the http client in the session, always responding with 200 and non-nil but empty body.
type mockClient struct {
	// StatusCode is the status code that should be returned
	StatusCode int

	// Return error
	Error error

	// Whether or not the context should or should not be checked for cancellation
	ContextCancelling bool

	// What data should we return
	Data []byte

	// Duration to sleep on
	Sleep time.Duration
}

func (mc mockClient) Do(req *http.Request) (*http.Response, error) {
	// First, sleep
	time.Sleep(mc.Sleep)

	// If we're told to return an error
	if mc.Error != nil {
		return nil, mc.Error
	}

	// Check for cancellation
	if mc.ContextCancelling {
		ctx := req.Context()
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
	}

	// Produce the response
	rec := httptest.NewRecorder()

	// Write the header if not zero
	if code := mc.StatusCode; code != 0 {
		rec.WriteHeader(code)
	}

	// Write the data if not nil
	if mc.Data != nil {
		_, err := rec.Write(mc.Data)
		if err != nil {
			return nil, err
		}
	}

	// Obtain results & return
	resp := rec.Result()
	return resp, nil
}

func mockSession(mc mockClient) *Session {
	return &Session{
		APIURL: NavitiaAPIURL,
		client: mc,
	}
}
