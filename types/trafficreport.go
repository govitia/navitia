package types

// A TrafficReport made of a network, an array of lines and an array of stop_areas.
// Named "traffic_report" in the Navitia doc
//
// See http://doc.navitia.io/#traffic-reports
// TODO: Add the internal links.
type TrafficReport struct {
	// Main object (network) and links within its own disruptions
	Network Network `json:"network"`

	// List of all disrupted Lines from the network
	Lines []Line `json:"lines"`

	// List of all disrupted StopAreas from the network
	StopAreas []StopArea `json:"stop_areas"`
}
