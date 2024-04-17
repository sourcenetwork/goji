//go:build js

package goji

import "syscall/js"

func init() {
	Event = eventJS(js.Global().Get("Event"))
}

type eventJS js.Value

// Event is a wrapper for the Event global interface.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event
var Event eventJS

// New wraps the Event constructor.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/Event
func (e eventJS) New(eventType string, options js.Value) eventValue {
	res := js.Value(e).New(eventType, options)
	return eventValue(res)
}

type eventValue js.Value

// Bubbles returns the Event bubbles property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/bubbles
func (e eventValue) Bubbles() bool {
	return js.Value(e).Get("bubbles").Bool()
}

// Cancelable returns the Event cancelable property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/cancelable
func (e eventValue) Cancelable() bool {
	return js.Value(e).Get("cancelable").Bool()
}

// Composed returns the Event composed property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/composed
func (e eventValue) Composed() bool {
	return js.Value(e).Get("composed").Bool()
}

// CurrentTarget returns the Event currentTarget property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/currentTarget
func (e eventValue) CurrentTarget() eventTargetValue {
	res := js.Value(e).Get("currentTarget")
	return eventTargetValue(res)
}

// DefaultPrevented returns the Event defaultPrevented property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/defaultPrevented
func (e eventValue) DefaultPrevented() bool {
	return js.Value(e).Get("defaultPrevented").Bool()
}

// EventPhase returns the Event eventPhase property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/eventPhase
func (e eventValue) EventPhase() int {
	return js.Value(e).Get("eventPhase").Int()
}

// IsTrusted returns the Event isTrusted property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/isTrusted
func (e eventValue) IsTrusted() bool {
	return js.Value(e).Get("isTrusted").Bool()
}

// Target returns the Event target property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (e eventValue) Target() eventTargetValue {
	res := js.Value(e).Get("target")
	return eventTargetValue(res)
}

// TimeStamp returns the Event timeStamp property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/timeStamp
func (e eventValue) TimeStamp() float64 {
	return js.Value(e).Get("timeStamp").Float()
}

// Type returns the Event type property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/type
func (e eventValue) Type() string {
	return js.Value(e).Get("type").String()
}

// ComposedPath wraps the Event composedPath instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/composedPath
func (e eventValue) ComposedPath() js.Value {
	return js.Value(e).Call("composedPath")
}

// PreventDefault wraps the Event preventDefault instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault
func (e eventValue) PreventDefault() {
	js.Value(e).Call("preventDefault")
}

// StopImmediatePropagation wraps the Event stopImmediatePropagation instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopImmediatePropagation
func (e eventValue) StopImmediatePropagation() {
	js.Value(e).Call("stopImmediatePropagation")
}

// StopPropagation wraps the Event stopPropagation instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopPropagation
func (e eventValue) StopPropagation() {
	js.Value(e).Call("stopPropagation")
}
