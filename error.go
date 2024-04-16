//go:build js

package goji

import "syscall/js"

func init() {
	Error = errorJS(js.Global().Get("Error"))
}

type errorJS js.Value

// Error is a wrapper for the Error global object.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error
var Error errorJS

// New wraps the Error constructor.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error/Error
func (e errorJS) New(message string) ErrorValue {
	res := js.Value(e).New(message)
	return ErrorValue(res)
}

var _ (error) = (ErrorValue)(js.Undefined())

// ErrorValue is an instance of
type ErrorValue js.Value

// Error wraps the Error toString prototype method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error/toString
func (v ErrorValue) Error() string {
	res := js.Value(v).Call("toString")
	return res.String()
}
