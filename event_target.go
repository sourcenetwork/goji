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
func (e eventTargetJS) New() eventTargetValue {
	res := js.Value(e).New()
	return eventTargetValue(res)
}

type eventTargetValue js.Value

// AddEventListener wraps the EventTarget addEventListener instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (e eventTargetValue) AddEventListener(eventType string, listener js.Value, options js.Value) {
	js.Value(e).Call("addEventListener", eventType, listener, options)
}

// DispatchEvent wraps the EventTarget dispatchEvent instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/dispatchEvent
func (e eventTargetValue) DispatchEvent(event js.Value) bool {
	return js.Value(e).Call("dispatchEvent", event).Bool()
}

// RemoveEventListener wraps the EventTarget removeEventListener instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/removeEventListener
func (e eventTargetValue) RemoveEventListener(eventType string, listener js.Value, options js.Value) {
	js.Value(e).Call("removeEventListener", eventType, listener, options)
}
