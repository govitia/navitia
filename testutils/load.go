/*
Package testutils implements utils for testing JSON no-compare unmarshalling.
*/
package testutils

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"

	"strings"

	"io"

	"github.com/pkg/errors"
)

// FileSizeLimit is the maximum size of files we accept, 10 megabytes currently
const FileSizeLimit = 30 * (1000 * 1000)

// A TestPair is a pair of JSON data along with the go representation they should have
type TestPair struct {
	Raw     []byte
	Correct interface{}
}

// TestData countains test data relative to their filenames for a specific category
type TestData struct {
	Known     map[string]TestPair // Hardcoded map of pairs of JSON data / the Go representation they should have
	Correct   map[string][]byte   // Map of known correct files to be parsed (they should _not_ return an error)
	Incorrect map[string][]byte   // Map of known incorrect files to be parsed (they _should_ return an error)
	Bench     map[string][]byte   // Descriptions -> File
}

/*
Load loads from a directory the TestData, and returns a map[categoryName]TestData.

The path parameter should be absolute.

The categories parameter should be a list of expected directory names, in the example hierarchy below, it would be: []string{"foo","bar"}.
Only the directories matching the given category list will be read.

This is what the directory should look like
	directory/
		foo/
			correct/
				a.json // The extension needs to be .json
				yolo.json
			incorrect/
				empty.json
				corruptedName.json
			bench/
				heavy.json
				regular.json
				light.json
		bar/
			...
*/
func Load(ctx context.Context, mainDirPath string, categories []string) (map[string]*TestData, error) {
	// Create the output
	output := make(map[string]*TestData, len(categories))

	// Read the subdirectories
	categoryDirectories, err := ioutil.ReadDir(mainDirPath)
	if err != nil {
		return nil, errors.Wrapf(err, "Load: error while reading testdata dir (%s)", mainDirPath)
	}

	// Create a map of categories towards a channel of *TestData
	rets := make(map[string](chan *TestData), len(categories))

	// Iterate through the category directories
	for _, catDir := range categoryDirectories {
		// If it is not a directory, skip
		if !catDir.IsDir() {
			continue
		}

		// If the name doesn't match one of the listed categories, continue
		// Note: SearchStrings works because ioutil.ReadDir returns a list sorted in the alphabetical order.
		if j := sort.SearchStrings(categories, catDir.Name()); j == len(categories) {
			continue
		}

		// Create the channel, blocking
		c := make(chan *TestData)

		// Add it to the channel map
		rets[catDir.Name()] = c

		// Create the absolute path for the category directory
		path := filepath.Join(mainDirPath, catDir.Name())

		// Create the category context, with an additional timeout
		cctx, cancel := context.WithTimeout(ctx, 12*time.Second)
		defer cancel()

		// Launch the goroutine
		go func(ctx context.Context, catDirPath string, c chan *TestData) {
			// Call the loadCategory function
			td, err := loadCategory(ctx, catDirPath)
			if err != nil {
				// Do something else than panic, please ?
				panic(err)
			}

			// Send the result via the channel
			c <- td
		}(cctx, path, c)
	}

	// Select on the goroutines
	for name, c := range rets {
		select {
		// Add the received category testdata to the output
		case td := <-c:
			output[name] = td
		// In case the context gets canceled
		case <-ctx.Done():
			return output, ctx.Err()
		}
	}

	// Return
	return output, nil
}

// loadCategory loads a category, "foo" in the example of Load
func loadCategory(ctx context.Context, catDirPath string) (*TestData, error) {
	// Create the output TestData
	output := new(TestData)

	// Read the subdirectories
	subDirectories, err := ioutil.ReadDir(catDirPath)
	if err != nil {
		return nil, errors.Wrapf(err, "loadCategory: error while reading subdirectories in %s", catDirPath)
	}

	// Iterate through them
	for i, sub := range subDirectories {
		// If this is not a directory, skip
		if !sub.IsDir() {
			continue
		}

		// Check for context cancelation
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// Build the subdirectory's path
		corpusPath := filepath.Join(catDirPath, sub.Name())

		// Switch to populate output
		switch sub.Name() {
		case "correct":
			output.Correct, err = loadCorpus(ctx, corpusPath)
		case "incorrect":
			output.Incorrect, err = loadCorpus(ctx, corpusPath)
		case "bench":
			output.Bench, err = loadCorpus(ctx, corpusPath)
		default: // If none of them, skip
			continue
		}

		// Check for error
		if err != nil {
			return nil, errors.Wrapf(err, "loadCategory: error while loading corpus in subdirectory #%d (%s)", i, corpusPath)
		}
	}

	// Done !
	return output, nil
}

// loadCorpus loads a corpus from a category's given subdirectory, "correct", "incorrect" or "bench"
func loadCorpus(ctx context.Context, path string) (map[string][]byte, error) {
	// Read the directory
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, errors.Wrapf(err, "loadCorpus: error while reading subdirectories of %s", path)
	}

	// Create the output
	output := make(map[string][]byte, len(files))

	// Iterate through the files, and load them up !
	for i, finfo := range files {
		// If this is a directory or it doesn't contain .json, skip
		if finfo.IsDir() || !strings.Contains(finfo.Name(), ".json") {
			continue
		}

		// Build the path
		fpath := filepath.Join(path, finfo.Name())

		// Open the file
		file, err := os.Open(fpath)
		if err != nil {
			return nil, errors.Wrapf(err, "loadCorpus: error while opening file #%d (%s)", i, fpath)
		}

		// Limit the reader
		reader := io.LimitReader(file, FileSizeLimit)

		// Read all
		data, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, errors.Wrapf(err, "loadCorpus: error while reading from file #%d (%s)", i, fpath)
		}

		// Copy it to output
		output[finfo.Name()] = data
	}

	// Return that all
	return output, nil
}
