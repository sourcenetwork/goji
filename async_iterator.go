//go:build js

package goji

import (
	"syscall/js"
)

// AsyncIteratorOf wraps the given channel into an async iterator object.
func AsyncIteratorOf(from <-chan any) js.Value {
	object := js.Global().Get("Object")
	symbol := js.Global().Get("Symbol").Get("asyncIterator")
	next := js.FuncOf(func(this js.Value, args []js.Value) any {
		prom := PromiseOf(func(resolve, reject func(value js.Value)) {
			v, ok := <-from
			if !ok {
				resolve(js.ValueOf(map[string]any{"done": true}))
			} else {
				resolve(js.ValueOf(map[string]any{"done": false, "value": v}))
			}
		})
		return js.Value(prom)
	})
	value := js.FuncOf(func(this js.Value, args []js.Value) any {
		return js.ValueOf(map[string]any{"next": next})
	})
	return object.Call("defineProperty", object.New(), symbol, map[string]any{"value": value})
}

type AsyncIteratorResult struct {
	Value js.Value
	Error error
}

// ForAwaitOf is a helper that wraps an async iterator in a channel.
func ForAwaitOf(value js.Value) <-chan AsyncIteratorResult {
	object := js.Global().Get("Object")
	symbol := js.Global().Get("Symbol").Get("asyncIterator")
	desc := object.Call("getOwnPropertyDescriptor", value, symbol)
	iter := desc.Call("value")
	result := make(chan AsyncIteratorResult)
	go func() {
		defer close(result)
		for {
			prom := iter.Get("next").Invoke()
			res, err := Await(PromiseValue(prom))
			out := AsyncIteratorResult{Error: err}
			done := js.Undefined()
			if err == nil {
				done = res[0].Get("done")
				out.Value = res[0].Get("value")
			}
			if done.Type() == js.TypeBoolean && done.Bool() {
				return
			}
			result <- out
			if err != nil {
				return
			}
		}
	}()
	return result
}
