//go:build js

package goji

import (
	"sync"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventTargetAddListener(t *testing.T) {
	var wait sync.WaitGroup

	// create an event listener to receive the event
	listener := EventListener(func(event EventValue) {
		defer wait.Done()
		assert.Equal(t, "test", event.Type())
	})
	defer listener.Release()

	// create an event to dispatch
	event := Event.New("test", js.Undefined())

	// setup a target and listener for the event
	target := EventTarget.New()
	target.AddEventListener(event.Type(), listener.Value, js.Undefined())

	wait.Add(1)
	dispatched := target.DispatchEvent(js.Value(event))
	wait.Wait()

	assert.True(t, dispatched)
}
