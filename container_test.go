package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

var containers map[string]*Container

// loadPC loads the containers in their final form for testing
func loadPC() error {
	// Get the input
	known := testData["container"].known
	if len(known) == 0 {
		return nil
	}

	cs := make(map[string]*Container, len(known))
	// For each of them, unmarshal and add to containers
	for name, datum := range known {
		var c = &Container{}

		err := c.UnmarshalJSON(datum.raw)
		if err != nil {
			return fmt.Errorf("Error while unmarshalling: %v", err)
		}

		cs[name] = c
	}

	containers = cs

	return nil
}

// TestContainer_Object_NoCompare tests the Container.Object method
func TestContainer_Object_NoCompare(t *testing.T) {
	// Get the input
	data := containers
	if len(data) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run it in parallel
	rgen := func(c *Container) func(t *testing.T) {
		return func(t *testing.T) {
			// Get the object
			obj, err := c.Object()
			if err != nil {
				t.Errorf("Error while calling .Object(): %v", err)
			}

			// Log it
			t.Logf("Object: %#v", obj)
		}
	}

	// For each of them, let's run a subtest
	for name, datum := range data {
		// Create the run function
		rfunc := rgen(datum)

		// Run !
		t.Run(name, rfunc)
	}
}

// TestContainer_Check_NoCompare tests the Container.Check method
func TestContainer_Check_NoCompare(t *testing.T) {
	// Get the input
	data := containers
	if len(data) == 0 {
		t.Skip("No data to test")
	}

	// Create the run function generator, allowing us to run it in parallel
	rgen := func(pc *Container) func(t *testing.T) {
		return func(t *testing.T) {
			err := pc.Check()
			if err != nil {
				t.Errorf("Check gave us invalid results: %v", err)
			}
		}
	}

	// For each of them, let's run a subtest
	for name, datum := range data {
		// Create the run function
		rfunc := rgen(datum)

		// Run !
		t.Run(name, rfunc)
	}
}

// BenchmarkContainer_UnmarshalJSON benchmarks Container.UnmarshalJSON through benchmarks
func BenchmarkContainer_UnmarshalJSON(b *testing.B) {
	// Get the bench data
	data := testData["container"].bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in []byte) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Unmarshal a Journey
				var c = &Container{}
				_ = c.UnmarshalJSON(in)
			}
		}
	}

	// Loop over all corpus
	for name, datum := range data {
		// Get run function
		runFunc := runGen(datum)

		// Run it !
		b.Run(name, runFunc)
	}
}

// BenchmarkContainer_Check benchmarks Container.Check through subbenchmarks
func BenchmarkContainer_Check(b *testing.B) {
	// Get the bench data
	data := testData["container"].bench
	if len(data) == 0 {
		b.Skip("No data to test")
	}

	// Run function generator, allowing parallel run
	runGen := func(in Container) func(*testing.B) {
		return func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Call .Check
				_ = in.Check()
			}
		}
	}

	// Loop over all corpus
	for name, datum := range data {
		var c = Container{}

		err := json.Unmarshal(datum, &c)
		if err != nil {
			b.Errorf("Error while unmarshalling: %v", err)
		}

		// Get run function
		runFunc := runGen(c)

		// Run it !
		b.Run(name, runFunc)
	}
}
