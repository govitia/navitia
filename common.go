package gonavitia

// Paging olds key paging information
type Paging struct {
	// The page number of this response
	PageNumber uint

	// The number of items on this page
	ItemsRetrievedOnLastRequest uint

	// The number of items retrieved, including anterior paging actions
	ItemsRetrievedToal uint

	// The total number of items we could get
	ItemsTotal uint

	// The URL to call to request the previous page
	previousURL string

	// The URL to call to request the next page
	nextURL string
}

// results is satisfied by nearly every result
type results interface {
	// retrieve retrieves the specified url and parses it
	retrieve(url string) (results, error)

	// add appends the given result to the one worked on
	add(newres results) error

	// paging retrieves the paging struct
	paging() Paging
}

// next retrieves the next elements
func next(res results) (results, error) {
	newResults, err := res.retrieve(res.paging().nextURL)
	return newResults, err
}

// previous retrieves the previous elements
func previous(res results) (results, error) {
	newResults, err := res.retrieve(res.paging().previousURL)
	return newResults, err
}

// previousAppend retrieves the previous elements & append them to the current one
func previousAppend(res results) error {
	newResults, err := res.retrieve(res.paging().previousURL)
	if err != nil {
		return err
	}

	err = res.add(newResults)
	return err
}

// nextAppend retrieves the next elements & append them to the current one
func nextAppend(res results) error {
	newResults, err := res.retrieve(res.paging().nextURL)
	if err != nil {
		return err
	}

	err = res.add(newResults)
	return err
}

type datetime string
