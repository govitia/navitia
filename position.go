package types

// PositionID is implemented by every type coding for some position
type PositionID interface {
	FormatURL() (string, error)
}
