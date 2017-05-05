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
		// We only want addresses, omit if you want to allow everything
		Types: []string{"address"},

		// We simply wish for six results.
		Count: 6,
	}

	// Execute it, query for "rue du caire"
	ctx := context.Background()
	res, _ := session.Places(ctx, "rue du caire", req)

	// Create a variable to store it
	var (
		myPlace   types.Container
		myAddress types.Address
	)

	// Check if there are enough results, and then assign the first element as your place
	if len(res.Places) == 0 {
		myPlace = res.Places[0]
		obj, _ := myPlace.Place()
		myAddress = obj.(types.Address)
	}

	// Print the name
	fmt.Printf("Name: %s, house number %d\n", myPlace.Name, myAddress.HouseNumber)
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
