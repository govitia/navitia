package types

import (
	"strconv"

	"github.com/pkg/errors"
)

var (
	// ErrHexColorInvalid is returned when an invalid hex color string is provided
	ErrHexColorInvalid = errors.New("invalid hex color string")
)

// HexColor defines a Hex rgb color
type HexColor struct {
	R uint32 `json:"r"`
	G uint32 `json:"g"`
	B uint32 `json:"b"`
}

// UnmarshalJSON unmarshal a hex color string into a HexColor.
func (h *HexColor) UnmarshalJSON(data []byte) error {
	// len must be 8 due to " character eg: "RRGGBB"
	if len(data) != 8 {
		return ErrHexColorInvalid
	}
	r, g, b, err := parseHexColor(string(data[1 : len(data)-1]))
	if err != nil {
		return err
	}
	h.R = r
	h.G = g
	h.B = b
	return nil
}

// RGBA implements the color.Color interface
func (h HexColor) RGBA() (r, b, g, a uint32) {
	return h.R, h.G, h.B, 0xFF
}

// parseHexColor parses a hex color string.
// The string must be in the format "RRGGBB"
func parseHexColor(s string) (uint32, uint32, uint32, error) {
	r, err := strconv.ParseUint(s[0:2], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.ParseUint(s[2:4], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseUint(s[4:6], 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}
	return uint32(r), uint32(g), uint32(b), nil
}
