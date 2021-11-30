package types

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestHexColor_UnmarshalJSON tests json unmarshaller implementation for HexColor type.
func TestHexColor_UnmarshalJSON(t *testing.T) {
    t.Parallel()
    var testColor HexColor

	require.NoError(t, json.Unmarshal([]byte(`"112233"`), &testColor))
	assert.Equal(t, uint32(0x11), testColor.R)
	assert.Equal(t, uint32(0x22), testColor.G)
	assert.Equal(t, uint32(0x33), testColor.B)

	require.NoError(t, json.Unmarshal([]byte(`"112233"`), &testColor))
	assert.Equal(t, uint32(0x11), testColor.R)
	assert.Equal(t, uint32(0x22), testColor.G)
	assert.Equal(t, uint32(0x33), testColor.B)
}

func TestHexColor_RGBA(t *testing.T) {
	t.Parallel()

	testColor := HexColor{R: 0x11, G: 0x22, B: 0x33}
	r, g, b, a := testColor.RGBA()
	assert.Equal(t, uint32(0x11), r)
	assert.Equal(t, uint32(0x22), g)
	assert.Equal(t, uint32(0x33), b)
	assert.Equal(t, uint32(0xFF), a)
}
