package navitia

import (
	//"fmt"
	"testing"
)

func Test_Coverage(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	res, err := testSession.Coverage(0)
	t.Logf("Received res: %v", *res)
	if err != nil {
		t.Fatalf("Got error in Coverage(%d): %v", 0, err)
	}
}

/*
func Test_Coverage_Complete(t *testing.T) {
	if *apiKey == "" {
		t.Skip(skipNoKey)
	}

	var i uint = 3
	for i <= 1200 {
		subtest := func(t *testing.T) {
			t.Parallel()
			res, err := testSession.Coverage(i)
			t.Logf("Received res: %v", *res)
			if err != nil {
				t.Fatalf("Got error in Coverage(%d): %v", 0, err)
			}

			// Check for length
			if len(res.Regions) > int(i) {
				t.Errorf("Amount of regions returned superior to given limit")
			}
		}
		t.Run(fmt.Sprintf("With a paging limit of %d items", i),subtest)

		i = i*i
	}
}*/
