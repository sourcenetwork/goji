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
func (e eventJS) New(eventType string, opts ...eventOption) EventValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(e).New(eventType, options)
		return EventValue(res)

	default:
		res := js.Value(e).New(eventType)
		return EventValue(res)
	}
}

// EventValue is an instance of Event.
type EventValue js.Value

// Bubbles returns the Event bubbles property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/bubbles
func (e EventValue) Bubbles() bool {
	return js.Value(e).Get("bubbles").Bool()
}

// Cancelable returns the Event cancelable property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/cancelable
func (e EventValue) Cancelable() bool {
	return js.Value(e).Get("cancelable").Bool()
}

// Composed returns the Event composed property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/composed
func (e EventValue) Composed() bool {
	return js.Value(e).Get("composed").Bool()
}

// CurrentTarget returns the Event currentTarget property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/currentTarget
func (e EventValue) CurrentTarget() EventTargetValue {
	res := js.Value(e).Get("currentTarget")
	return EventTargetValue(res)
}

// DefaultPrevented returns the Event defaultPrevented property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/defaultPrevented
func (e EventValue) DefaultPrevented() bool {
	return js.Value(e).Get("defaultPrevented").Bool()
}

// EventPhase returns the Event eventPhase property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/eventPhase
func (e EventValue) EventPhase() int {
	return js.Value(e).Get("eventPhase").Int()
}

// IsTrusted returns the Event isTrusted property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/isTrusted
func (e EventValue) IsTrusted() bool {
	return js.Value(e).Get("isTrusted").Bool()
}

// Target returns the Event target property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/target
func (e EventValue) Target() EventTargetValue {
	res := js.Value(e).Get("target")
	return EventTargetValue(res)
}

// TimeStamp returns the Event timeStamp property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/timeStamp
func (e EventValue) TimeStamp() float64 {
	return js.Value(e).Get("timeStamp").Float()
}

// Type returns the Event type property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/type
func (e EventValue) Type() string {
	return js.Value(e).Get("type").String()
}

// ComposedPath wraps the Event composedPath instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/composedPath
func (e EventValue) ComposedPath() js.Value {
	return js.Value(e).Call("composedPath")
}

// PreventDefault wraps the Event preventDefault instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/preventDefault
func (e EventValue) PreventDefault() {
	js.Value(e).Call("preventDefault")
}

// StopImmediatePropagation wraps the Event stopImmediatePropagation instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopImmediatePropagation
func (e EventValue) StopImmediatePropagation() {
	js.Value(e).Call("stopImmediatePropagation")
}

// StopPropagation wraps the Event stopPropagation instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/stopPropagation
func (e EventValue) StopPropagation() {
	js.Value(e).Call("stopPropagation")
}

// EventOptions is used to set event listener options.
var EventOptions = &eventOptions{}

type eventOptions struct{}

type eventOption func(value js.Value)

// WithBubbles sets the bubbles options.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/Event#bubbles
func (e eventOptions) WithBubbles(enabled bool) eventOption {
	return func(value js.Value) {
		value.Set("bubbles", enabled)
	}
}

// WithCancelable sets the cancelable option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/Event#cancelable
func (e eventOptions) WithCancelable(enabled bool) eventOption {
	return func(value js.Value) {
		value.Set("cancelable", enabled)
	}
}

// WithComposed sets the composed option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/Event/Event#composed
func (e eventOptions) WithComposed(enabled bool) eventOption {
	return func(value js.Value) {
		value.Set("composed", enabled)
	}
}
