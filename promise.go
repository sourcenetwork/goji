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
func (p promiseJS) New(executor js.Func) promiseValue {
	res := js.Value(p).New(executor)
	return promiseValue(res)
}

// All wraps the Promise all static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/all
func (p promiseJS) All(iterable js.Value) promiseValue {
	res := js.Value(p).Call("all", iterable)
	return promiseValue(res)
}

// AllSettled wraps the Promise allSettled static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/allSettled
func (p promiseJS) AllSettled(iterable js.Value) promiseValue {
	res := js.Value(p).Call("allSettled", iterable)
	return promiseValue(res)
}

// Any wraps the Promise any static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/any
func (p promiseJS) Any(iterable js.Value) promiseValue {
	res := js.Value(p).Call("any", iterable)
	return promiseValue(res)
}

// Race wraps the Promise race static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/race
func (p promiseJS) Race(iterable js.Value) promiseValue {
	res := js.Value(p).Call("race", iterable)
	return promiseValue(res)
}

// Reject wraps the Promise reject static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/reject
func (p promiseJS) Reject(value js.Value) promiseValue {
	res := js.Value(p).Call("reject", value)
	return promiseValue(res)
}

// Resolve wraps the Promise resolve static method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/resolve
func (p promiseJS) Resolve(value js.Value) promiseValue {
	res := js.Value(p).Call("resolve", value)
	return promiseValue(res)
}

type promiseValue js.Value

// Then wraps the Promise then prototype method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/then
func (v promiseValue) Then(onFulfilled js.Func) promiseValue {
	res := js.Value(v).Call("then", onFulfilled)
	return promiseValue(res)
}

// Catch wraps the Promise catch prototype method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/catch
func (v promiseValue) Catch(onRejected js.Func) promiseValue {
	res := js.Value(v).Call("catch", onRejected)
	return promiseValue(res)
}

// Finally wraps the Promise finally prototype method.
//
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise/finally
func (v promiseValue) Finally(onFinally js.Value) promiseValue {
	res := js.Value(v).Call("catch", onFinally)
	return promiseValue(res)
}

// PromiseOf is a helper function that wraps the given func in a promise.
func PromiseOf(fn func(resolve, reject func(value js.Value))) promiseValue {
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
func Await(prom promiseValue) (res []js.Value, err error) {
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
