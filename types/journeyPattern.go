package types

// JourneyPattern A journey pattern is an ordered list of stop points.
// Two vehicles that serve exactly the same stop points in
// exactly the same order belong to to the same journey pattern.
type JourneyPattern struct {
	ID   string
	Name string
}
