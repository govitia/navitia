package types

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
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

	err := load()
	if err != nil {
		panic(err)
	}
}

type testPair struct {
	raw     []byte
	correct interface{}
}

// Convert testing mechanism to known compare + corpus runs
type typeTestData struct {
	known     map[string]testPair // Hardcoded map of pairs of JSON data / the Go representation they should have
	correct   map[string][]byte   // Map of known correct files to be parsed (they should _not_ return an error)
	incorrect map[string][]byte   // Map of known incorrect files to be parsed (they _should_ return an error)
	bench     map[string][]byte   // Descriptions -> File
}

// testData stores a map which maps each category to their data
var testData = make(map[string]typeTestData, len(typesList))

// this is the list of potential types
// must be lower case
var typesList = []string{
	"journey",
	"coverage",
	"container",
	"route",
	"line",
	"network",
	"company",
}

// listCategoryDirs retrieves the subdirectories under the main testdata directory
func listCategoryDirs(path string) ([]os.FileInfo, error) {
	mainSubdirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "Error while reading %s's files", path)
	}

	var subDirsInfo []os.FileInfo

	// Iterate through the subdirs
	for _, dinfo := range mainSubdirs {
		// We have a category !
		if dinfo.IsDir() {
			subDirsInfo = append(subDirsInfo, dinfo)
		}
	}

	return subDirsInfo, nil
}

// extractCorpus extracts a corpus
func extractCorpus(path string) (map[string][]byte, error) {
	// List the files
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Create the data
	var corpus = make(map[string][]byte, len(files))

	// For each of them, populate testpairs
	for _, finfo := range files {
		// Get the name
		name := finfo.Name()

		// Build the path
		filePath := filepath.Join(path, name)

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			return corpus, err
		}

		// Read it all
		read, err := ioutil.ReadAll(file)
		if err != nil {
			return corpus, err
		}

		// Assign it
		corpus[name] = read
	}

	return corpus, nil
}

// getPertinentSubdirs, given a dir in a category subdirectory, returns the awaited values
func getCategory(path string) (typeTestData, error) {
	// Create the data
	var data = typeTestData{}

	// List the subdirs
	subdirs, err := ioutil.ReadDir(path)
	if err != nil {
		return data, err
	}

	// Iterate through the subdirs
	for _, dinfo := range subdirs {
		if !dinfo.IsDir() {
			continue // skip this iteration
		}
		switch dinfo.Name() {
		case "correct":
			correctPath := filepath.Join(path, "correct")
			data.correct, err = extractCorpus(correctPath)
			if err != nil {
				return data, err
			}
		case "incorrect":
			incorrectPath := filepath.Join(path, "incorrect")
			data.incorrect, err = extractCorpus(incorrectPath)
			if err != nil {
				return data, err
			}
		case "bench":
			benchPath := filepath.Join(path, "bench")
			data.bench, err = extractCorpus(benchPath)
			if err != nil {
				return data, err
			}
		}
	}
	// return
	return data, nil
}

// load loads the file structing into the testData
func load() error {

	subDirsInfo, err := listCategoryDirs(testDataPath)
	if err != nil {
		return err
	}

	// For each of them, call getCategory
	for _, dinfo := range subDirsInfo {
		name := dinfo.Name()
		path := filepath.Join(testDataPath, name)

		data, err := getCategory(path)
		if err != nil {
			return err
		}
		fmt.Printf("For %s we have:\n\tCorrect:\t%d\n\tIncorrect:\t%d\n\tBench:\t%d\n", path, len(data.correct), len(data.incorrect), len(data.bench))
		testData[name] = data
	}

	// Load containers
	err = loadContainers()

	return err
}
