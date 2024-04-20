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

// RequestValue is an instance of IDBRequest.
type RequestValue struct {
	goji.EventTargetValue
}

// Error returns the IDBRequest error property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/error
func (r RequestValue) Error() js.Value {
	return js.Value(r.EventTargetValue).Get("error")
}

// Result returns the IDBRequest result property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/result
func (r RequestValue) Result() js.Value {
	return js.Value(r.EventTargetValue).Get("result")
}

// Source returns the IDBRequest source property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/source
func (r RequestValue) Source() js.Value {
	return js.Value(r.EventTargetValue).Get("source")
}

// ReadyState returns the IDBRequest readyState property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/readyState
func (r RequestValue) ReadyState() string {
	return js.Value(r.EventTargetValue).Get("readyState").String()
}

// Transaction returns the IDBRequest transaction property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/transaction
func (r RequestValue) Transaction() TransactionValue {
	res := js.Value(r.EventTargetValue).Get("transaction")
	return TransactionValue(res)
}

// Await is a helper that waits for a request and returns the result and error.
func Await(request RequestValue) (res js.Value, err error) {
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

	request.AddEventListener(SuccessEvent, onSuccess.Value)
	request.AddEventListener(ErrorEvent, onError.Value)

	wait.Wait()

	return
}
