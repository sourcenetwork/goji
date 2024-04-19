//go:build js

package goji

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventConstructor(t *testing.T) {
	event := Event.New("test")
	assert.Equal(t, "test", event.Type())
	assert.False(t, event.Bubbles())
	assert.False(t, event.Cancelable())
	assert.False(t, event.Composed())
}

func TestEventConstructorWithOptions(t *testing.T) {
	options := []eventOption{
		EventOptions.WithBubbles(true),
		EventOptions.WithCancelable(true),
		EventOptions.WithComposed(true),
	}

	event := Event.New("test", options...)
	assert.Equal(t, "test", event.Type())
	assert.True(t, event.Bubbles())
	assert.True(t, event.Cancelable())
	assert.True(t, event.Composed())
}
