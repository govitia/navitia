package types

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

// these are the types that can be embedded
const (
	embeddedStopArea       string = "stop_area"             // This is a Place & a PT Object
	embeddedPOI                   = "poi"                   // This is a place
	embeddedAddress               = "address"               // This is a place
	embeddedStopPoint             = "stop_point"            // This is a place
	embeddedAdmin                 = "administrative_region" // This is a place
	embeddedLine                  = "line"                  // This is a PT Object
	embeddedRoute                 = "route"                 // This is a PT Object
	embeddedNetwork               = "network"               // This is a PT Object
	embeddedCommercialMode        = "commercial_mode"       // This is a PT Object
	embeddedTrip                  = "trip"                  // This is a PT Object
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

	embeddedJSON json.RawMessage

	// embeddedObject acts as a cache, it is the only element guarded by the RWMutex
	embeddedObject Object

	// If multiple goroutines try to access embeddedObject while one is writing it, results are undefined.
	// So we guard against it through a mutex.
	mu *sync.RWMutex
}

// IsPlace returns true if the container's content is a Place
func (c *Container) IsPlace() bool {
	t := c.EmbeddedType
	return t == embeddedStopArea || t == embeddedPOI || t == embeddedAddress || t == embeddedStopPoint || t == embeddedAdmin
}

// IsPTObject returns true if the container's content is a PTObject
func (c *Container) IsPTObject() bool {
	t := c.EmbeddedType
	return t == embeddedStopArea || t == embeddedLine || t == embeddedRoute || t == embeddedNetwork || t == embeddedCommercialMode || t == embeddedTrip
}

// ErrInvalidContainer is returned after a check on a Container
type ErrInvalidContainer struct {
	// If the Container has a zero ID.
	NoID bool

	// If the PlaceContainer has an EmbeddedType yet non-empty embedded content.
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
		msg += "\n\tEmpty EmbeddedType yet non-empty embedded content"
		anomalies++
	}
	if err.UnknownEmbeddedType {
		msg += "\n\tUnknown EmbeddedType"
		anomalies++
	}

	return fmt.Sprintf(msg, anomalies)
}

// Empty returns true if the container is empty (zero value)
func (c *Container) Empty() bool {
	return c.ID == "" && c.Name == "" && c.EmbeddedType == "" && c.Quality == 0 && len(c.embeddedJSON) == 0 && c.embeddedObject == nil
}

// Check checks the validity of the Container. Returns an ErrInvalidContainer.
//
// An empty Container is valid. But those cases aren't:
// 	- If the Container has an empty ID.
// 	- If the Container has an empty EmbeddedType & a non-empty embedded struct inside.
// 	- If the Container has an unknown EmbeddedType.
func (c *Container) Check() error {
	// Check if the container is empty
	if c.Empty() {
		return nil
	}

	// Create the error to be populated
	err := ErrInvalidContainer{}

	// Check for zero ID
	err.NoID = (c.ID == "")

	// Check if the embedded type is empty & there is a non-empty embedded content inside, that's an error
	if c.EmbeddedType == "" && (len(c.embeddedJSON) != 0 || c.embeddedObject != nil) {
		err.NoEmbeddedType = true
		return err
	} else if c.EmbeddedType == "" { // Else, if the embedded type indicator is empty, the rest is useless
		return nil
	}

	// Else, check if the declared EmbeddedType is known.
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

// Object returns the Object contained in a Container.
// If the Container is empty, Object returns an error.
// Check() is run on the Container.
func (c *Container) Object() (Object, error) {
	if c.EmbeddedType == "" {
		return nil, nil
	}

	// If we already have an embedded object, return it
	c.mu.RLock()
	o := c.embeddedObject
	c.mu.RUnlock()
	if o != nil {
		return o, nil
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
	c.mu.Lock()
	c.embeddedObject = obj
	c.mu.Unlock()

	// Let's return it
	return obj, nil
}

// Place returns the Place contained in the container if that is what's inside
//
// If the Object isn't a Place or the Container is empty or invalid, Place returns an error
func (c *Container) Place() (Place, error) {
	// Check if its a Place
	if !c.IsPlace() {
		return nil, errors.Errorf("container's content isn't a Place")
	}

	// Return it then
	obj, err := c.Object()
	if err != nil {
		return nil, err
	}

	// Type assert
	return obj.(Place), nil
}

// PTObject returns the PTObject contained in the container if that is what's inside
//
// If the Object isn't a PTObject or the Container is empty or invalid, Place returns an error
func (c *Container) PTObject() (PTObject, error) {
	// Check if its a Place
	if !c.IsPTObject() {
		return nil, errors.Errorf("container's content isn't a PTObject")
	}

	// Return it then
	obj, err := c.Object()
	if err != nil {
		return nil, err
	}

	// Type assert
	return obj.(PTObject), nil
}
