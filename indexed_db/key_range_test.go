//go:build js

package indexed_db

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyRangeBound(t *testing.T) {
	key := KeyRange.Bound(js.ValueOf("A"), js.ValueOf("F"), true, false)
	assert.Equal(t, "A", key.Lower().String())
	assert.Equal(t, "F", key.Upper().String())
	assert.Equal(t, true, key.LowerOpen())
	assert.Equal(t, false, key.UpperOpen())

	assert.True(t, key.Includes(js.ValueOf("B")))
	assert.False(t, key.Includes(js.ValueOf("Z")))
}

func TestKeyRangeLowerBound(t *testing.T) {
	key := KeyRange.LowerBound(js.ValueOf(5), false)
	assert.Equal(t, 5, key.Lower().Int())
	assert.Equal(t, js.Undefined(), key.Upper())
	assert.Equal(t, false, key.LowerOpen())
	assert.Equal(t, true, key.UpperOpen())

	assert.True(t, key.Includes(js.ValueOf(5)))
	assert.False(t, key.Includes(js.ValueOf(0)))
}

func TestKeyRangeUpperBound(t *testing.T) {
	key := KeyRange.UpperBound(js.ValueOf(3), false)
	assert.Equal(t, js.Undefined(), key.Lower())
	assert.Equal(t, 3, key.Upper().Int())
	assert.Equal(t, true, key.LowerOpen())
	assert.Equal(t, false, key.UpperOpen())

	assert.True(t, key.Includes(js.ValueOf(0)))
	assert.False(t, key.Includes(js.ValueOf(4)))
}

func TestKeyRangeOnly(t *testing.T) {
	key := KeyRange.Only(js.ValueOf("K"))
	assert.Equal(t, "K", key.Lower().String())
	assert.Equal(t, "K", key.Upper().String())
	assert.Equal(t, false, key.LowerOpen())
	assert.Equal(t, false, key.UpperOpen())

	assert.True(t, key.Includes(js.ValueOf("K")))
	assert.False(t, key.Includes(js.ValueOf("J")))
}
