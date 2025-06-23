//go:build js

package goji

import (
	"sync"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventTargetDispatch(t *testing.T) {
	var wait sync.WaitGroup

	// create an event listener to receive the event
	listener := EventListener(func(event EventValue) {
		defer wait.Done()
		assert.Equal(t, "test", event.Type())
	})
	defer listener.Release()

	// create an event to dispatch
	event := Event.New("test")

	// setup a target and listener for the event
	target := EventTarget.New()
	target.AddEventListener(event.Type(), listener.Value)

	wait.Add(1)
	dispatched := target.DispatchEvent(js.Value(event))
	wait.Wait()

	assert.True(t, dispatched)
}

func TestEventTargetDispatchCustom(t *testing.T) {
	var wait sync.WaitGroup

	// create an event listener to receive the event
	listener := EventListener(func(event EventValue) {
		defer wait.Done()
		custom := CustomEventValue(event)
		assert.Equal(t, "test", custom.Event().Type())
		assert.Equal(t, js.ValueOf(1), custom.Detail())
	})
	defer listener.Release()

	// create an event to dispatch
	event := CustomEvent.New("test", js.ValueOf(1))

	// setup a target and listener for the event
	target := EventTarget.New()
	target.AddEventListener(event.Event().Type(), listener.Value)

	wait.Add(1)
	dispatched := target.DispatchEvent(js.Value(event))
	wait.Wait()

	assert.True(t, dispatched)
}
