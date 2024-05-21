//go:build js

package indexed_db

import (
	"syscall/js"
)

// DatabaseValue is an instance of IDBDatabase
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase
type DatabaseValue js.Value

// Name returns the IDBDatabase name property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/name
func (d DatabaseValue) Name() string {
	return js.Value(d).Get("name").String()
}

// ObjectStoreNames returns the IDBDatabase objectStoreNames property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/objectStoreNames
func (d DatabaseValue) ObjectStoreNames() js.Value {
	return js.Value(d).Get("objectStoreNames")
}

// Version returns the IDBDatabase version property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/version
func (d DatabaseValue) Version() int {
	return js.Value(d).Get("version").Int()
}

// Close wraps the IDBDatabase close instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/close
func (d DatabaseValue) Close() {
	js.Value(d).Call("close")
}

// CreateObjectStore wraps the IDBDatabase createObjectStore instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/createObjectStore
func (d DatabaseValue) CreateObjectStore(name string, opts ...objectStoreOption) ObjectStoreValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(d).Call("createObjectStore", name, options)
		return ObjectStoreValue(res)

	default:
		res := js.Value(d).Call("createObjectStore", name)
		return ObjectStoreValue(res)
	}
}

// DeleteObjectStore wraps the IDBDatabase deleteObjectStore instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/deleteObjectStore
func (d DatabaseValue) DeleteObjectStore(name string) {
	js.Value(d).Call("deleteObjectStore", name)
}

// Transaction wraps the IDBDatabase transaction instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/transaction
func (d DatabaseValue) Transaction(storeNames any, mode string, opts ...transactionOption) TransactionValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(d).Call("transaction", storeNames, mode, options)
		return TransactionValue(res)

	default:
		res := js.Value(d).Call("transaction", storeNames, mode)
		return TransactionValue(res)
	}
}
