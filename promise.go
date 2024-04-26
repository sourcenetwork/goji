//go:build js

package goji

import (
	"sync"
	"syscall/js"
)

func init() {
	Promise = promiseJS(js.Global().Get("Promise"))
}

type promiseJS js.Value

// Promise is a wrapper for the Promise global object.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise
var Promise promiseJS

// New wraps the Promise constructor.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/Promise
func (p promiseJS) New(executor js.Func) PromiseValue {
	res := js.Value(p).New(executor)
	return PromiseValue(res)
}

// All wraps the Promise all static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all
func (p promiseJS) All(iterable js.Value) PromiseValue {
	res := js.Value(p).Call("all", iterable)
	return PromiseValue(res)
}

// AllSettled wraps the Promise allSettled static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/allSettled
func (p promiseJS) AllSettled(iterable js.Value) PromiseValue {
	res := js.Value(p).Call("allSettled", iterable)
	return PromiseValue(res)
}

// Any wraps the Promise any static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/any
func (p promiseJS) Any(iterable js.Value) PromiseValue {
	res := js.Value(p).Call("any", iterable)
	return PromiseValue(res)
}

// Race wraps the Promise race static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/race
func (p promiseJS) Race(iterable js.Value) PromiseValue {
	res := js.Value(p).Call("race", iterable)
	return PromiseValue(res)
}

// Reject wraps the Promise reject static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/reject
func (p promiseJS) Reject(value js.Value) PromiseValue {
	res := js.Value(p).Call("reject", value)
	return PromiseValue(res)
}

// Resolve wraps the Promise resolve static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/resolve
func (p promiseJS) Resolve(value js.Value) PromiseValue {
	res := js.Value(p).Call("resolve", value)
	return PromiseValue(res)
}

// PromiseValue is an instance of Promise.
type PromiseValue js.Value

// Then wraps the Promise then prototype method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/then
func (v PromiseValue) Then(onFulfilled js.Func) PromiseValue {
	res := js.Value(v).Call("then", onFulfilled)
	return PromiseValue(res)
}

// Catch wraps the Promise catch prototype method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/catch
func (v PromiseValue) Catch(onRejected js.Func) PromiseValue {
	res := js.Value(v).Call("catch", onRejected)
	return PromiseValue(res)
}

// Finally wraps the Promise finally prototype method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/finally
func (v PromiseValue) Finally(onFinally js.Value) PromiseValue {
	res := js.Value(v).Call("catch", onFinally)
	return PromiseValue(res)
}

// PromiseOf is a helper function that wraps the given func in a promise.
func PromiseOf(fn func(resolve, reject func(value js.Value))) PromiseValue {
	var executor js.Func
	executor = js.FuncOf(func(this js.Value, args []js.Value) any {
		executor.Release()
		resolve := func(value js.Value) {
			args[0].Invoke(value)
		}
		reject := func(value js.Value) {
			args[1].Invoke(value)
		}
		fn(resolve, reject)
		return js.Undefined()
	})
	return Promise.New(executor)
}

// Await is a helper function that waits for a promise to resolve or reject
// and returns the results and an error value.
func Await(prom PromiseValue) (res []js.Value, err error) {
	var wait sync.WaitGroup

	onFulfilled := js.FuncOf(func(this js.Value, args []js.Value) any {
		defer wait.Done()
		res = args
		return js.Undefined()
	})
	defer onFulfilled.Release()

	onRejected := js.FuncOf(func(this js.Value, args []js.Value) any {
		defer wait.Done()
		err = ErrorValue(args[0])
		return js.Undefined()
	})
	defer onRejected.Release()

	wait.Add(1)
	prom.Then(onFulfilled).Catch(onRejected)
	wait.Wait()

	return
}

// Async is a helper function that wraps the given func in a promise that
// resolves when no error is returned or rejects when an error is returned.
func Async(fn func(this js.Value, args []js.Value) (js.Value, error)) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		prom := PromiseOf(func(resolve, reject func(value js.Value)) {
			res, err := fn(this, args)
			if err != nil {
				reject(js.Value(Error.New(err.Error())))
			} else {
				resolve(res)
			}
		})
		return js.Value(prom)
	})
}
