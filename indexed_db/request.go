//go:build js

package indexed_db

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

// RequestValue is an instance of IDBRequest.
type RequestValue goji.EventTargetValue

// Error returns the IDBRequest error property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/error
func (r RequestValue) Error() js.Value {
	return js.Value(r).Get("error")
}

// Result returns the IDBRequest result property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/result
func (r RequestValue) Result() js.Value {
	return js.Value(r).Get("result")
}

// Source returns the IDBRequest source property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/source
func (r RequestValue) Source() js.Value {
	return js.Value(r).Get("source")
}

// ReadyState returns the IDBRequest readyState property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/readyState
func (r RequestValue) ReadyState() string {
	return js.Value(r).Get("readyState").String()
}

// Transaction returns the IDBRequest transaction property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBRequest/transaction
func (r RequestValue) Transaction() TransactionValue {
	res := js.Value(r).Get("transaction")
	return TransactionValue(res)
}
