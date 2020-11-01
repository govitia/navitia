package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type journeyRequest struct {
	Journeys []json.RawMessage `json:"journeys"`
}

func (r journeyRequest) messages() []json.RawMessage {
	return r.Journeys
}

type placeRequest struct {
	Places []json.RawMessage `json:"places"`
}

func (r placeRequest) messages() []json.RawMessage {
	return r.Places
}

type coverageRequest struct {
	Regions []json.RawMessage `json:"regions"`
}

func (r coverageRequest) messages() []json.RawMessage {
	return r.Regions
}

type request interface {
	messages() []json.RawMessage
}

var (
	originFlag      = flag.String("from", "../../testdata", "Original directory")
	destinationFlag = flag.String("to", "./", "Destination directory")
	originPath      string
	destinationPath string
)

var equivalence = map[string]string{
	"journeys": "journey",
	"places":   "place",
	"coverage": "region",
}

// load loads the origin directory and files
func load(path string) (map[string][]*os.File, error) {
	subDirs, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("Error while listing subdirs !: %v\n", err)
		return nil, err
	}

	originFiles := make(map[string][]*os.File, len(subDirs))

	for _, dinfo := range subDirs {
		dirName := dinfo.Name()
		if !dinfo.IsDir() {
			fmt.Printf("Skipping %s...\n", dirName)
		} else {
			fmt.Printf("Processing %s directory...\n", dirName)
			files, err := ioutil.ReadDir(filepath.Join(originPath, dirName))
			if err != nil {
				fmt.Printf("Error while reading %s directory !: %v\n", dirName, err)
				return originFiles, err
			}
			for _, finfo := range files {
				fmt.Printf("\tProcessing %s...\n", finfo.Name())
				fname := finfo.Name()
				if fname[len(fname)-4:] != "json" {
					fmt.Printf("\t\tSkipping\n")
				} else {
					path := filepath.Join(originPath, dirName, fname)
					f, err := os.Open(path)
					if err != nil {
						fmt.Printf("\t\tError while opening file %s! : %v\n", path, err)
					}
					originFiles[dirName] = append(originFiles[dirName], f)
				}
			}
		}
	}

	return originFiles, nil
}

func main() {
	flag.Parse()

	if filepath.IsAbs(*originFlag) {
		originPath = *originFlag
	} else {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error while retrieving working directory, please retry with an absolute path: %v\n", err)
			return
		}
		originPath = filepath.Join(wd, *originFlag)
	}

	if filepath.IsAbs(*destinationFlag) {
		destinationPath = *destinationFlag
	} else {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error while retrieving working directory, please retry with an absolute path: %v\n", err)
			return
		}
		destinationPath = filepath.Join(wd, *destinationFlag)
	}

	originFiles, err := load(originPath)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	// For each of them, process them
	for cat, files := range originFiles {
		cat = equivalence[cat]
		fmt.Printf("Printing %s...\n", cat)
		for _, file := range files {
			fmt.Printf("Dealing with one origin-file...\n")
			// Prepare decoder
			dec := json.NewDecoder(file)

			// Create hosting structure & decode to it
			var req request
			switch cat {
			case "journey":
				tmp := &journeyRequest{}
				// Decode to it
				err := dec.Decode(tmp)
				if err != nil {
					stat, _ := file.Stat()
					fmt.Printf("Error while decoding %s: %v\n", stat.Name(), err)
					return
				}
				req = tmp
			case "place":
				tmp := &placeRequest{}
				// Decode to it
				err := dec.Decode(tmp)
				if err != nil {
					stat, _ := file.Stat()
					fmt.Printf("Error while decoding %s: %v\n", stat.Name(), err)
					return
				}
				req = tmp
			case "region":
				tmp := &coverageRequest{}
				// Decode to it
				err := dec.Decode(tmp)
				if err != nil {
					stat, _ := file.Stat()
					fmt.Printf("Error while decoding %s: %v\n", stat.Name(), err)
					return
				}
				req = tmp
			default:
				fmt.Printf("Incorrect category")
				return
			}

			// Get file stat
			stat, err := file.Stat()
			if err != nil {
				fmt.Printf("Error while retrieving file stat: %v\n", err)
				continue
			}

			// Now for each Journey, create a new file and write to it
			for i, message := range req.messages() {
				// Create the file name
				nname := fmt.Sprintf("%s%d.json", stat.Name()[:len(stat.Name())-5], i)
				npath := filepath.Join(destinationPath, cat, nname)

				// Create the file
				nfile, err := os.Create(npath)
				if err != nil {
					fmt.Printf("Error while creating file %s: %v\n", npath, err)
				}

				// Write to it
				enc := json.NewEncoder(nfile)
				enc.SetIndent("", "\t")
				err = enc.Encode(message)
				if err != nil {
					fmt.Printf("Error while writing to file %s: %v\n", npath, err)
				}
			}

		}
	}

}
