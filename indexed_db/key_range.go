//go:build js

package indexed_db

import "syscall/js"

func init() {
	KeyRange = keyRange(js.Global().Get("IDBKeyRange"))
}

// KeyRange wraps the IDBKeyRange global interface.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange
var KeyRange keyRange

type keyRange js.Value

// Bound wraps the IDBKeyRange bound static method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/bound_static
func (k keyRange) Bound(lower, upper js.Value, lowerOpen, upperOpen bool) KeyRangeValue {
	res := js.Value(k).Call("bound", lower, upper, lowerOpen, upperOpen)
	return KeyRangeValue(res)
}

// LowerBound wraps the IDBKeyRange lowerBound static method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/lowerBound_static
func (k keyRange) LowerBound(lower js.Value, open bool) KeyRangeValue {
	res := js.Value(k).Call("lowerBound", lower, open)
	return KeyRangeValue(res)
}

// Only wraps the IDBKeyRange only static method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/only_static
func (k keyRange) Only(value js.Value) KeyRangeValue {
	res := js.Value(k).Call("only", value)
	return KeyRangeValue(res)
}

// UpperBound wraps the IDBKeyRange upperBound static method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/upperBound_static
func (k keyRange) UpperBound(upper js.Value, open bool) KeyRangeValue {
	res := js.Value(k).Call("upperBound", upper, open)
	return KeyRangeValue(res)
}

// KeyRangeValue is an instance of IDBKeyRange.
type KeyRangeValue js.Value

// Lower returns the IDBKeyRange lower property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/lower
func (k KeyRangeValue) Lower() js.Value {
	return js.Value(k).Get("lower")
}

// LowerOpen returns the IDBKeyrange lowerOpen property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/lowerOpen
func (k KeyRangeValue) LowerOpen() bool {
	return js.Value(k).Get("lowerOpen").Bool()
}

// Upper returns the IDBKeyRange upper property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/upper
func (k KeyRangeValue) Upper() js.Value {
	return js.Value(k).Get("upper")
}

// UpperOpen returns the IDBKeyRange upperOpen property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/upperOpen
func (k KeyRangeValue) UpperOpen() bool {
	return js.Value(k).Get("upperOpen").Bool()
}

// Includes wraps the IDBKeyRange includes instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBKeyRange/includes
func (k KeyRangeValue) Includes(key js.Value) bool {
	return js.Value(k).Call("includes", key).Bool()
}
