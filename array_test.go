//go:build js

package goji

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewArray(t *testing.T) {
	a1 := Array.New()
	assert.Equal(t, 0, a1.Length())

	a2 := Array.New(5)
	assert.Equal(t, 5, a2.Length())

	a3 := Array.New("one", "two")
	assert.Equal(t, 2, a3.Length())
}

func TestArrayAt(t *testing.T) {
	arr := Array.Of("one", "two", "three")
	assert.Equal(t, "one", arr.At(0).String())
	assert.Equal(t, "two", arr.At(1).String())
	assert.Equal(t, "three", arr.At(2).String())
}

func TestIsArray(t *testing.T) {
	arr := Array.New()
	assert.True(t, Array.IsArray(js.Value(arr)))
	assert.False(t, Array.IsArray(map[string]any{}))
	assert.False(t, Array.IsArray(1))
}

func TestArrayReverse(t *testing.T) {
	arr := Array.Of("one", "two", "three").Reverse()
	assert.Equal(t, "three", arr.At(0).String())
	assert.Equal(t, "two", arr.At(1).String())
	assert.Equal(t, "one", arr.At(2).String())
}

func TestArrayJoin(t *testing.T) {
	res := Array.Of("one", "two", "three").Join("+")
	assert.Equal(t, "one+two+three", res)
}
