//go:build js

package web_assembly

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

func init() {
	webAssembly = js.Global().Get("WebAssembly")
}

var webAssembly js.Value

// Compile is a wrapper for the WebAssembly compile static method.
//
// https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface/compile_static
func Compile(bufferSource js.Value) goji.PromiseValue {
	res := webAssembly.Call("compile", bufferSource)
	return goji.PromiseValue(res)
}

// Instantiate is a wrapper for the WebAssembly instantiate static method.
//
// https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface/instantiate_static
func Instantiate(bufferSourceOrModule js.Value, importObject js.Value) goji.PromiseValue {
	res := webAssembly.Call("instantiate", bufferSourceOrModule, importObject)
	return goji.PromiseValue(res)
}

// Validate is a wrapper for the WebAssembly validate static method.
//
// https://developer.mozilla.org/en-US/docs/WebAssembly/JavaScript_interface/validate_static
func Validate(bufferSource js.Value) bool {
	res := webAssembly.Call("validate", bufferSource)
	return res.Bool()
}
