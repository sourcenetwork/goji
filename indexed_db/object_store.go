//go:build js

package indexed_db

import (
	"syscall/js"
)

// ObjectStoreValue is an instance of IDBObjectStore.
type ObjectStoreValue js.Value

// AutoIncrement returns the IDBObjectStore autoIncrement property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/autoIncrement
func (o ObjectStoreValue) AutoIncrement() bool {
	return js.Value(o).Get("autoIncrement").Bool()
}

// IndexNames returns the IDBObjectStore indexNames property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/indexNames
func (o ObjectStoreValue) IndexNames() js.Value {
	return js.Value(o).Get("indexNames")
}

// KeyPath returns the IDBObjectStore keyPath property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/keyPath
func (o ObjectStoreValue) KeyPath() js.Value {
	return js.Value(o).Get("keyPath")
}

// Name returns the IDBObjectStore name property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/name
func (o ObjectStoreValue) Name() string {
	return js.Value(o).Get("name").String()
}

// Transaction returns the IDBObjectStore transaction property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/transaction
func (o ObjectStoreValue) Transaction() TransactionValue {
	res := js.Value(o).Get("transaction")
	return TransactionValue(res)
}

// Add wraps the IDBObjectStore add instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/add
func (o ObjectStoreValue) Add(args ...any) RequestValue[js.Value] {
	res := js.Value(o).Call("add", args...)
	return RequestValue[js.Value](res)
}

// Clear wraps the IDBObjectStore clear instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/clear
func (o ObjectStoreValue) Clear() RequestValue[js.Value] {
	res := js.Value(o).Call("clear")
	return RequestValue[js.Value](res)
}

// Count wraps the IDBObjectStore count instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/count
func (o ObjectStoreValue) Count(query any) RequestValue[js.Value] {
	res := js.Value(o).Call("count", query)
	return RequestValue[js.Value](res)
}

// CreateIndex wraps the IDBObjectStore createIndex instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/createIndex
func (o ObjectStoreValue) CreateIndex(indexName string, keyPath any, opts ...indexOption) IndexValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(o).Call("createIndex", indexName, keyPath, options)
		return IndexValue(res)

	default:
		res := js.Value(o).Call("createIndex", indexName, keyPath)
		return IndexValue(res)
	}
}

// Delete wraps the IDBObjectStore delete instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/delete
func (o ObjectStoreValue) Delete(key any) RequestValue[js.Value] {
	res := js.Value(o).Call("delete", key)
	return RequestValue[js.Value](res)
}

// DeleteIndex wraps the IDBObjectStore deleteIndex instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/deleteIndex
func (o ObjectStoreValue) DeleteIndex(key any) {
	js.Value(o).Call("deleteIndex", key)
}

// Get wraps the IDBObjectStore get instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/get
func (o ObjectStoreValue) Get(key any) RequestValue[js.Value] {
	res := js.Value(o).Call("get", key)
	return RequestValue[js.Value](res)
}

// GetAll wraps the IDBObjectStore getAll instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/getAll
func (o ObjectStoreValue) GetAll(args ...any) RequestValue[js.Value] {
	res := js.Value(o).Call("getAll", args...)
	return RequestValue[js.Value](res)
}

// GetAllKeys wraps the IDBObjectStore getAllKeys instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/getAllKeys
func (o ObjectStoreValue) GetAllKeys(args ...any) RequestValue[js.Value] {
	res := js.Value(o).Call("getAllKeys", args...)
	return RequestValue[js.Value](res)
}

// GetKey wraps the IDBObjectStore getKey instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/getKey
func (o ObjectStoreValue) GetKey(key any) RequestValue[js.Value] {
	res := js.Value(o).Call("getKey", key)
	return RequestValue[js.Value](res)
}

// Index wraps the IDBObjectStore index instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/index
func (o ObjectStoreValue) Index(name string) IndexValue {
	res := js.Value(o).Call("index", name)
	return IndexValue(res)
}

// OpenCursor wraps the IDBObjectStore openCursor instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/openCursor
func (o ObjectStoreValue) OpenCursor(args ...any) RequestValue[CursorWithValue] {
	res := js.Value(o).Call("openCursor", args...)
	return RequestValue[CursorWithValue](res)
}

// OpenKeyCursor wraps the IDBObjectStore openKeyCursor instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/openKeyCursor
func (o ObjectStoreValue) OpenKeyCursor(args ...any) RequestValue[CursorValue] {
	res := js.Value(o).Call("openKeyCursor", args...)
	return RequestValue[CursorValue](res)
}

// Put wraps the IDBObjectStore put instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/put
func (o ObjectStoreValue) Put(args ...any) RequestValue[js.Value] {
	res := js.Value(o).Call("put", args...)
	return RequestValue[js.Value](res)
}

// ObjectStoreOptions is used to set object store options.
var ObjectStoreOptions = &objectStoreOptions{}

type objectStoreOptions struct{}

type objectStoreOption func(opts js.Value)

// WithKeyPath sets the key path option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/createObjectStore#keypath
func (objectStoreOptions) WithKeyPath(keyPath any) objectStoreOption {
	return func(opts js.Value) {
		opts.Set("keyPath", keyPath)
	}
}

// WithAutoIncrement sets the auto increment option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBDatabase/createObjectStore#autoincrement
func (objectStoreOptions) WithAutoIncrement(enabled bool) objectStoreOption {
	return func(opts js.Value) {
		opts.Set("autoIncrement", enabled)
	}
}
