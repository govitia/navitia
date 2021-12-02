package types

import (
	"time"
)

type Period struct {
	Begin time.Time `json:"begin"`
	End   time.Time `json:"end"`
}
