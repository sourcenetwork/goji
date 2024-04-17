//go:build js

package goji

import (
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventConstructor(t *testing.T) {
	event := Event.New("test", js.Undefined())
	assert.Equal(t, "test", event.Type())
	assert.False(t, event.Bubbles())
	assert.False(t, event.Cancelable())
	assert.False(t, event.Composed())
}

func TestEventConstructorWithOptions(t *testing.T) {
	options := js.ValueOf(map[string]any{
		"bubbles":    true,
		"cancelable": true,
		"composed":   true,
	})

	event := Event.New("test", options)
	assert.Equal(t, "test", event.Type())
	assert.True(t, event.Bubbles())
	assert.True(t, event.Cancelable())
	assert.True(t, event.Composed())
}
