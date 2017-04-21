package types

// A Route represents a route: a Line can have several routes, that is several directions with potential embranchments and different frequency for each.
//
// See http://doc.navitia.io/#public-transport-objects
type Route struct {
	// Identifier of the route
	// For example: "route:RAT:M6"
	ID ID

	// Name of the route
	// For example:"Nation - Charles de Gaule Etoile"
	Name string

	// Frequence is true when the route has frequency, if it doesn't it stays false
	Frequence bool

	// Line is the line it is connected to
	Line Line

	// Direction is the direction of the route
	// Note: As direction is a Place, it can be a POI in some data
	Direction Container
}
