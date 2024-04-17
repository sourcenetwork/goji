//go:build js

package goji

import "syscall/js"

func init() {
	Uint8Array = uint8ArrayJS(js.Global().Get("Uint8Array"))
}

type uint8ArrayJS js.Value

// Uint8Array is a wrapper for the Uint8Array global object.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Uint8Array
var Uint8Array uint8ArrayJS

// New is a wrapper for the Uint8Array constructor.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Uint8Array/Uint8Array
func (a uint8ArrayJS) New(length js.Value) uint8ArrayInstance {
	res := js.Value(a).New(length)
	return uint8ArrayInstance(res)
}

type uint8ArrayInstance js.Value

// Length is a wrapper for Uint8Array length property.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/TypedArray/length
func (a uint8ArrayInstance) Length() int {
	return js.Value(a).Get("length").Int()
}

// Uint8ArrayFromBytes is a helper function that copies the given
// byte slice into a new Uint8Array.
func Uint8ArrayFromBytes(src []byte) uint8ArrayInstance {
	len := js.ValueOf(len(src))
	dst := Uint8Array.New(len)
	js.CopyBytesToJS(js.Value(dst), src)
	return dst
}

// BytesFromUint8Array is a helper function that copies the given
// Uint8Array into a new byte slice.
func BytesFromUint8Array(src js.Value) []byte {
	len := uint8ArrayInstance(src).Length()
	dst := make([]byte, len)
	js.CopyBytesToGo(dst, src)
	return dst
}