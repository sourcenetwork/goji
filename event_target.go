//go:build js

package goji

import "syscall/js"

func init() {
	EventTarget = eventTargetJS(js.Global().Get("EventTarget"))
}

type eventTargetJS js.Value

// EventTarget wraps the EventTarget global interface.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget
var EventTarget eventTargetJS

// New wraps the EventTarget constructor.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/EventTarget
func (e eventTargetJS) New() EventTargetValue {
	res := js.Value(e).New()
	return EventTargetValue(res)
}

// EventTargetValue is an instance of EventTarget.
type EventTargetValue js.Value

// AddEventListener wraps the EventTarget addEventListener instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (e EventTargetValue) AddEventListener(eventType string, listener js.Value, options js.Value) {
	js.Value(e).Call("addEventListener", eventType, listener, options)
}

// DispatchEvent wraps the EventTarget dispatchEvent instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent
func (e EventTargetValue) DispatchEvent(event js.Value) bool {
	return js.Value(e).Call("dispatchEvent", event).Bool()
}

// RemoveEventListener wraps the EventTarget removeEventListener instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
func (e EventTargetValue) RemoveEventListener(eventType string, listener js.Value, options js.Value) {
	js.Value(e).Call("removeEventListener", eventType, listener, options)
}

// EventListener returns a new event listener callback
// that calls the given func when an event is received.
func EventListener(fn func(event EventValue)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		event := EventValue(args[0])
		fn(event)
		return js.Undefined()
	})
}
