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
	known  map[string]testPair // Pairs of files + known concrete go type
	corpus map[string][]byte   // List of corpus files by their names
	bench  map[string][]byte   // Descriptions -> File
}

// testData stores a map which maps each category to their data
var testData = make(map[string]typeTestData, len(typesList))

// this is the list of potential types
// must be lower case
var typesList = []string{
	"journey",
	"section",
	"region",
	"place",
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

// extractKnown extracts the list of testpairs
func extractKnown(path string) (map[string]testPair, error) {
	// List the files
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Create the data
	var testpairs = make(map[string]testPair, len(files))

	// For each of them, populate testpairs
	for _, finfo := range files {
		// Get the name
		name := finfo.Name()

		// Build the path
		filePath := filepath.Join(path, name)

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			return testpairs, err
		}

		// Read it all
		read, err := ioutil.ReadAll(file)
		if err != nil {
			return testpairs, err
		}

		// Assign it
		pair := testPair{
			raw: read,
		}
		testpairs[name] = pair
	}

	return testpairs, nil
}

// extractCorpus extracts the corpus map
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
		dirPath := filepath.Join(path, name)

		// Open the file
		file, err := os.Open(dirPath)
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

// extractBench extracts the bench map
func extractBench(path string) (map[string][]byte, error) {
	// List the files
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Create the data
	var bench = make(map[string][]byte, len(files))

	// For each of them, populate testpairs
	for _, finfo := range files {
		// Get the name
		name := finfo.Name()

		// Build the path
		filePath := filepath.Join(path, name)

		// Open the file
		file, err := os.Open(filePath)
		if err != nil {
			return bench, err
		}

		// Read it all
		read, err := ioutil.ReadAll(file)
		if err != nil {
			return bench, err
		}

		// Assign it
		bench[name] = read
	}

	return bench, nil
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
		if dinfo.IsDir() {
			switch dinfo.Name() {
			case "known":
				knownPath := filepath.Join(path, "known")
				data.known, err = extractKnown(knownPath)
				if err != nil {
					return data, err
				}
			case "corpus":
				corpusPath := filepath.Join(path, "corpus")
				data.corpus, err = extractCorpus(corpusPath)
				if err != nil {
					return data, err
				}
			case "bench":
				benchPath := filepath.Join(path, "bench")
				data.bench, err = extractBench(benchPath)
				if err != nil {
					return data, err
				}
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
		fmt.Printf("For %s we have:\n\tKnown:\t%d\n\tCorpus:\t%d\n\tBench:\t%d\n", path, len(data.known), len(data.corpus), len(data.bench))
		testData[name] = data
	}

	// Subloads
	err = loadPC()

	return err
}
