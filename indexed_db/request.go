//go:build js

package indexed_db

import (
	"errors"
	"sync"
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

const (
	// The error handler is executed when an error
	// caused a request to fail.
	ErrorEvent = "error"
	// The success event is fired when an IDBRequest succeeds.
	SuccessEvent = "success"
)

// RequestResult is the type union for request results.
type RequestResult interface {
	js.Value | DatabaseValue | CursorValue | CursorWithValue
}

// RequestValue is an instance of IDBRequest.
type RequestValue[T RequestResult] js.Value

// Error returns the IDBRequest error property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/error
func (r RequestValue[T]) Error() js.Value {
	return js.Value(r).Get("error")
}

// Result returns the IDBRequest result property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/result
func (r RequestValue[T]) Result() T {
	res := js.Value(r).Get("result")
	return T(res)
}

// Source returns the IDBRequest source property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/source
func (r RequestValue[T]) Source() js.Value {
	return js.Value(r).Get("source")
}

// ReadyState returns the IDBRequest readyState property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/readyState
func (r RequestValue[T]) ReadyState() string {
	return js.Value(r).Get("readyState").String()
}

// Transaction returns the IDBRequest transaction property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/transaction
func (r RequestValue[T]) Transaction() TransactionValue {
	res := js.Value(r).Get("transaction")
	return TransactionValue(res)
}

// EventTarget returns the EventTarget for the request.
func (r RequestValue[T]) EventTarget() goji.EventTargetValue {
	return goji.EventTargetValue(r)
}

// Await is a helper that waits for a request and returns the result and error.
func Await[T RequestResult](request RequestValue[T]) (res T, err error) {
	var wait sync.WaitGroup

	onSuccess := goji.EventListener(func(event goji.EventValue) {
		defer wait.Done()
		res = request.Result()
	})
	defer onSuccess.Release()

	onError := goji.EventListener(func(event goji.EventValue) {
		defer wait.Done()
		err = errors.New(request.Error().Get("message").String())
	})
	defer onError.Release()

	wait.Add(1)

	request.EventTarget().AddEventListener(SuccessEvent, onSuccess.Value)
	request.EventTarget().AddEventListener(ErrorEvent, onError.Value)

	wait.Wait()

	return
}
