package types

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	testDataPathFlag = flag.String("path", "./testdata", "Directory of test data")
	testDataPath     string
)

// testData stores io.Readers for each of the files found in the directory
var testData = make(map[string][]*os.File, len(typesList))

// this is the list of potential types
// must be lower case
var typesList = []string{
	"journey",
	"section",
	"region",
	"place",
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
	subdirs, err := ioutil.ReadDir(testDataPath)
	if err != nil {
		return errors.Wrapf(err, "Error while reading %s's files", testDataPath)
	}

	// Iterate through the subdirs
	for _, dinfo := range subdirs {
		if dinfo.IsDir() {
			dpath := filepath.Join(testDataPath, dinfo.Name())
			files, err := ioutil.ReadDir(dpath)
			if err != nil {
				return errors.Wrapf(err, "error while reading %s's files", dpath)
			}

			// Now we'll iterate through the files, checking if their name include a type, and if so adding them to testData
			for _, finfo := range files {
				// Get the name
				fileName := finfo.Name()

				// Open it
				path := filepath.Join(testDataPath, dinfo.Name(), fileName)
				f, err := os.Open(path)
				if err != nil {
					return err
				}

				// Add it
				testData[dinfo.Name()] = append(testData[dinfo.Name()], f)
			}
		}
	}

	return nil
}
