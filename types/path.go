package types

import "time"

// A PathSegment (called Path item in the Navitia API) is a part of a path.
type PathSegment struct {
	// The Length of the segment
	Length uint `json:"length"`

	// The Name of the way corresponding to the segment
	Name string `json:"name"`

	// The duration in seconds of the segment
	Duration time.Duration `json:"duration"`

	// The angle in degree between the previous segment and this segment
	// = 0 Means going straight
	// < 0 Means turning left
	// > 0 Means turning right
	Direction int `json:"direction"`
}
