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

	var listener js.Func
	listener = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer wait.Done()
		listener.Release()
		return js.Undefined()
	})

	// create an event to dispatch
	event := Event.New("Test", js.Undefined())

	// setup a target and listener for the event
	target := EventTarget.New()
	target.AddEventListener(event.Type(), listener.Value, js.Undefined())

	wait.Add(1)
	dispatched := target.DispatchEvent(js.Value(event))
	wait.Wait()

	assert.True(t, dispatched)
}
