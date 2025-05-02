//go:build js

package goji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForAwaitOf(t *testing.T) {
	input := make(chan any)
	go func() {
		defer close(input)
		input <- true
		input <- false
	}()
	iter := AsyncIteratorOf(input)
	output := ForAwaitOf(iter)

	res := <-output
	assert.NoError(t, res.Error)
	assert.True(t, res.Value.Bool())

	res = <-output
	assert.NoError(t, res.Error)
	assert.False(t, res.Value.Bool())

	_, ok := <-output
	assert.False(t, ok)
}
