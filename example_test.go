package navitia_test

import (
	"context"
	"fmt"

	"github.com/aabizri/navitia"
	"github.com/aabizri/navitia/types"
)

var session *navitia.Session

func ExampleSession_Places() {
	// Create a request
	req := navitia.PlacesRequest{
		Query: "10 rue du caire, Paris",
		Types: []string{"address"},
		Count: 1,
	}

	// Execute it
	ctx := context.Background()
	res, _ := session.Places(ctx, req)

	// Create a variable to store it
	var myPlace types.Container

	// Check if there are enough results, and then assign the first element as your place
	if len(res.Places) != 0 {
		myPlace = res.Places[0]
	}

	// Print the name
	fmt.Printf("Name: %s\n", myPlace.Name)
}

var (
	myPlace        types.Container
	thatOtherPlace types.Container
)

func ExampleSession_Journeys() {
	// Create a request, having already found two corresponding places
	request := navitia.JourneyRequest{
		From: myPlace.ID,
		To:   thatOtherPlace.ID,
	}

	// Execute it
	ctx := context.Background()
	res, _ := session.Journeys(ctx, request)

	// Print the duration of the first journey
	fmt.Printf("Duration of journey #0: %s\n", res.Journeys[0].Duration.String())
}
