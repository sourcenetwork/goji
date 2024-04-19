//go:build js

package goji

import (
	"encoding/json"
	"syscall/js"
)

func init() {
	JSON = jsonJS(js.Global().Get("JSON"))
}

// JSON is a wrapper for the JSON global object.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON
var JSON jsonJS

type jsonJS js.Value

// Parse is a wrapper for the JSON parse static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/parse
func (j jsonJS) Parse(text string) js.Value {
	return js.Value(j).Call("parse", text)
}

// Stringify is a wrapper for the JSON stringify static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify
func (j jsonJS) Stringify(value js.Value) string {
	return js.Value(j).Call("stringify", value).String()
}

// MarshalJS marshals the given value into a js.Value.
func MarshalJS(v any) (js.Value, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return js.Undefined(), err
	}
	return JSON.Parse(string(data)), nil
}

// UnmarshalJS unmarshals the given js.Value into the given pointer.
func UnmarshalJS(value js.Value, v any) error {
	text := JSON.Stringify(value)
	return json.Unmarshal([]byte(text), v)
}
