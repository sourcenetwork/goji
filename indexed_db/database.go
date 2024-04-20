//go:build js

package indexed_db

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

// DatabaseValue is an instance of IDBDatabase
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase
type DatabaseValue struct {
	goji.EventTargetValue
}

// Name returns the IDBDatabase name property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/name
func (d DatabaseValue) Name() string {
	return js.Value(d.EventTargetValue).Get("name").String()
}

// ObjectStoreNames returns the IDBDatabase objectStoreNames property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/objectStoreNames
func (d DatabaseValue) ObjectStoreNames() js.Value {
	return js.Value(d.EventTargetValue).Get("objectStoreNames")
}

// Version returns the IDBDatabase version property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/version
func (d DatabaseValue) Version() int {
	return js.Value(d.EventTargetValue).Get("version").Int()
}

// Close wraps the IDBDatabase close instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/close
func (d DatabaseValue) Close() {
	js.Value(d.EventTargetValue).Call("close")
}

// CreateObjectStore wraps the IDBDatabase createObjectStore instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/createObjectStore
func (d DatabaseValue) CreateObjectStore(name string, options js.Value) ObjectStoreValue {
	res := js.Value(d.EventTargetValue).Call("createObjectStore", name, options)
	return ObjectStoreValue(res)
}

// DeleteObjectStore wraps the IDBDatabase deleteObjectStore instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/deleteObjectStore
func (d DatabaseValue) DeleteObjectStore(name string) {
	js.Value(d.EventTargetValue).Call("deleteObjectStore", name)
}

// Transaction wraps the IDBDatabase transaction instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/transaction
func (d DatabaseValue) Transaction(storeNames js.Value, mode string, options js.Value) TransactionValue {
	res := js.Value(d.EventTargetValue).Call("transaction", storeNames, mode, options)
	return TransactionValue(res)
}
