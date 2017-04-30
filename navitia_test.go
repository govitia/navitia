package navitia

import (
	"context"
	"flag"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"net/http/httputil"

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

type mockClient struct {
	tester *testing.T
}

func (mc mockClient) Do(req *http.Request) (*http.Response, error) {
	data, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		return nil, err
	}

	time.Sleep(10 * time.Millisecond)
	mc.tester.Log(data)
	return &http.Response{}, nil
}

func newMockClient(t *testing.T) interface {
	Do(req *http.Request) (*http.Response, error)
} {
	return mockClient{t}
}
