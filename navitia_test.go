package navitia

import (
	"context"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
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
	tester *testing.T
}

func (mc mockClient) Do(req *http.Request) (*http.Response, error) {
	time.Sleep(10 * time.Millisecond)
	resp := httptest.NewRecorder().Result()
	return resp, nil
}

func mockSession(t *testing.T) *Session {
	return &Session{
		APIURL: NavitiaAPIURL,
		client: mockClient{t},
	}
}
