//go:build js

package goji

import "syscall/js"

func init() {
	WebAssembly = webAssemblyJS(js.Global().Get("WebAssembly"))
}

type webAssemblyJS js.Value

// WebAssembly is a wrapper for the WebAssembly global object.
//
// https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface
var WebAssembly webAssemblyJS

// Compile is a wrapper for the WebAssembly compile static method.
//
// https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface/compile_static
func (w webAssemblyJS) Compile(bufferSource js.Value) promiseValue {
	res := js.Value(w).Call("compile", bufferSource)
	return promiseValue(res)
}

// Instantiate is a wrapper for the WebAssembly instantiate static method.
//
// https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface/instantiate_static
func (w webAssemblyJS) Instantiate(bufferSourceOrModule js.Value, importObject js.Value) promiseValue {
	res := js.Value(w).Call("instantiate", bufferSourceOrModule, importObject)
	return promiseValue(res)
}

// Validate is a wrapper for the WebAssembly validate static method.
//
// https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface/validate_static
func (w webAssemblyJS) Validate(bufferSource js.Value) bool {
	res := js.Value(w).Call("validate", bufferSource)
	return res.Bool()
}
