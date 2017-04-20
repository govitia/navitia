package types

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// these are the types that can be embedded
const (
	embeddedStopArea       string = "stop_area"
	embeddedPOI                   = "poi"
	embeddedAddress               = "address"
	embeddedStopPoint             = "stop_point"
	embeddedAdmin                 = "administrative_region"
	embeddedLine                  = "line"
	embeddedRoute                 = "route"
	embeddedNetwork               = "network"
	embeddedCommercialMode        = "commercial_mode"
	embeddedTrip                  = "trip"
)

var embeddedTypes = [...]string{
	embeddedStopArea,
	embeddedPOI,
	embeddedAddress,
	embeddedStopPoint,
	embeddedAdmin,
	embeddedLine,
	embeddedRoute,
	embeddedNetwork,
	embeddedCommercialMode,
	embeddedTrip,
}

// An Object is what is contained by a Container
type Object interface{}

// A Container holds an Object, which can be a Place or a PT Object
type Container struct {
	ID           ID
	Name         string
	EmbeddedType string
	Quality      int

	embeddedJSON   json.RawMessage
	embeddedObject Object
}

// ErrInvalidContainer is returned after a check on a Container
type ErrInvalidContainer struct {
	// If the Container has a zero ID.
	NoID bool

	// If the PlaceContainer has a zero EmbeddedType.
	NoEmbeddedType bool

	// If the PlaceContainer has an unknown EmbeddedType
	UnknownEmbeddedType bool
}

// Error satisfies the error interface
func (err ErrInvalidContainer) Error() string {
	// Count the number of anomalies
	var anomalies uint

	msg := "Error: Invalid non-empty PlaceContainer (%d anomalies):"

	if err.NoID {
		msg += "\n\tNo ID specified"
		anomalies++
	}
	if err.NoEmbeddedType {
		msg += "\n\tEmpty EmbeddedType"
		anomalies++
	}
	if err.UnknownEmbeddedType {
		msg += "\n\tUnknown EmbeddedType"
		anomalies++
	}

	return fmt.Sprintf(msg, anomalies)
}

func (c Container) Empty() bool {
	return c.ID == "" && c.Name == "" && c.EmbeddedType == "" && c.Quality == 0 && len(c.embeddedJSON) == 0 && c.embeddedObject == nil
}

// Check checks the validity of the Container. Returns an ErrInvalidContainer.
//
// An empty Container is valid. But those cases aren't:
// 	- If the Container has an empty ID.
// 	- If the Container has an empty EmbeddedType.
// 	- If the Container has an unknown EmbeddedType.
func (c Container) Check() error {
	// Check if the container is empty
	if c.Empty() {
		return nil
	}

	// Create the error to be populated
	err := ErrInvalidContainer{}

	// Check for zero ID
	err.NoID = (c.ID == "")

	// Check if the embedded type is empty
	if c.EmbeddedType == "" {
		err.NoEmbeddedType = true
		return err
	}

	// Check if the declared EmbeddedType is known.
	var known bool
	for _, ket := range embeddedTypes {
		if c.EmbeddedType == ket {
			known = true
			break
		}
	}
	err.UnknownEmbeddedType = !known

	// Check if there's any change
	emptyErr := ErrInvalidContainer{}
	if err != emptyErr {
		return err
	}

	return nil
}

// Object returns the Object contained in a Container
// If the Container is empty, Object returns an error.
// Check() is run on the Container.
func (c *Container) Object() (Object, error) {
	if c.EmbeddedType == "" {
		return nil, nil
	}

	// If we already have an embedded object, return it
	if obj := c.embeddedObject; obj != nil {
		return obj, nil
	}

	// Create the receiver
	var obj Object

	// Switch through
	switch c.EmbeddedType {
	case embeddedStopArea:
		obj = &StopArea{}
	case embeddedPOI:
		obj = &POI{}
	case embeddedAddress:
		obj = &Address{}
	case embeddedStopPoint:
		obj = &StopPoint{}
	case embeddedAdmin:
		obj = &Admin{}
	case embeddedLine:
		obj = &Line{}
	case embeddedRoute:
		obj = &Route{}
	case embeddedNetwork:
		obj = &Network{}
	case embeddedCommercialMode:
		obj = &CommercialMode{}
	case embeddedTrip:
		obj = &Trip{}
	default:
		return nil, errors.Errorf("no known embedded type indicated (we have \"%s\"), can't return a place !", c.EmbeddedType)
	}

	// Unmarshal into the receiver
	err := json.Unmarshal(c.embeddedJSON, obj)
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't unmarshal the embedded type (%s)", c.EmbeddedType)
	}

	// Let's add it to the container
	c.embeddedObject = obj

	// Let's return it
	return obj, nil
}
