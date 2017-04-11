package types

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	testDataPathFlag = flag.String("path", "./testdata", "Directory of test data")
	testDataPath     string
)

// testData stores io.Readers for each of the files found in the directory
var testData = make(map[string][]io.ReadCloser, len(typesList))

// this is the list of potential types
// must be lower case
var typesList = []string{
	"journey",
	"section",
	"coverage",
}

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

	err := load()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded: %#v\n", testData)
}

func load() error {
	files, err := ioutil.ReadDir(testDataPath)
	if err != nil {
		return err
	}

	// Now we'll iterate through the files, checking if their name include a type, and if so adding them to testData
	for _, finfo := range files {
		// Get the name
		fileName := finfo.Name()
		for _, typeName := range typesList {
			// If the fileName includes the typeName, then add it to the list
			if strings.Contains(fileName, typeName) {
				// Open the file
				fpath := filepath.Join(testDataPath, fileName)
				freader, err := os.Open(fpath)
				if err != nil {
					return err
				}

				// Add it to the slice
				testData[typeName] = append(testData[typeName], freader)

				// Break off
				break
			}
		}
	}

	return nil
}
