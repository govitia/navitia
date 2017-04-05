package gonavitia

type PositionID interface {
	formatURL() (string, error)
}
