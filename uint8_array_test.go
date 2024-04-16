//go:build js

package goji

import (
	"crypto/rand"
	"io"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUint8New(t *testing.T) {
	value := Uint8Array.New(js.ValueOf(10))
	assert.Equal(t, 10, value.Length())
}

func TestUint8ArrayFromBytes(t *testing.T) {
	r := io.LimitReader(rand.Reader, 64)
	bytes, err := io.ReadAll(r)
	require.NoError(t, err)

	value := Uint8ArrayFromBytes(bytes)
	assert.Equal(t, 64, value.Length())

	other := BytesFromUint8Array(js.Value(value))
	assert.Equal(t, bytes, other)
}
