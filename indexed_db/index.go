//go:build js

package indexed_db

import (
	"syscall/js"
)

// IndexValue is an IDBIndex instance.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex
type IndexValue js.Value

// KeyPath returns the IDBIndex keyPath property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/keyPath
func (i IndexValue) KeyPath() js.Value {
	return js.Value(i).Get("keyPath")
}

// MultiEntry returns the IDBIndex multiEntry property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/multiEntry
func (i IndexValue) MultiEntry() bool {
	return js.Value(i).Get("multiEntry").Bool()
}

// Name returns the IDBIndex name property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/name
func (i IndexValue) Name() string {
	return js.Value(i).Get("name").String()
}

// ObjectStore returns the IDBIndex objectStore property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/objectStore
func (i IndexValue) ObjectStore() ObjectStoreValue {
	res := js.Value(i).Get("objectStore")
	return ObjectStoreValue(res)
}

// Unique returns the IDBIndex unique property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/unique
func (i IndexValue) Unique() bool {
	return js.Value(i).Get("unique").Bool()
}

// Count wraps the IDBIndex count instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/count
func (i IndexValue) Count(key any) RequestValue[js.Value] {
	res := js.Value(i).Call("count", key)
	return RequestValue[js.Value](res)
}

// Get wraps the IDBIndex get instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/get
func (i IndexValue) Get(key any) RequestValue[js.Value] {
	res := js.Value(i).Call("get", key)
	return RequestValue[js.Value](res)
}

// GetAll wraps the IDBIndex getAll instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/getAll
func (i IndexValue) GetAll(args ...any) RequestValue[js.Value] {
	res := js.Value(i).Call("getAll", args...)
	return RequestValue[js.Value](res)
}

// GetAllKeys wraps the IDBIndex getAllKeys instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/getAllKeys
func (i IndexValue) GetAllKeys(args ...any) RequestValue[js.Value] {
	res := js.Value(i).Call("getAllKeys", args...)
	return RequestValue[js.Value](res)
}

// GetKey wraps the IDBIndex getKey instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/getKey
func (i IndexValue) GetKey(key any) RequestValue[js.Value] {
	res := js.Value(i).Call("getKey", key)
	return RequestValue[js.Value](res)
}

// OpenCursor wraps the IDBIndex openCursor instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/openCursor
func (i IndexValue) OpenCursor(args ...any) RequestValue[CursorValue] {
	res := js.Value(i).Call("openCursor", args...)
	return RequestValue[CursorValue](res)
}

// OpenKeyCursor wraps the IDBIndex openKeyCursor instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/openKeyCursor
func (i IndexValue) OpenKeyCursor(args ...any) RequestValue[CursorValue] {
	res := js.Value(i).Call("openKeyCursor", args...)
	return RequestValue[CursorValue](res)
}

// IndexOptions is used to set index options.
var IndexOptions = &indexOptions{}

type indexOptions struct{}

type indexOption func(opts js.Value)

// WithUnique sets the index unique option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/createIndex#unique
func WithUnique(enabled bool) indexOption {
	return func(opts js.Value) {
		opts.Set("unique", enabled)
	}
}

// WithMultiEntry sets the index multiEntry option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/createIndex#multientry
func WithMultiEntry(enabled bool) indexOption {
	return func(opts js.Value) {
		opts.Set("multiEntry", enabled)
	}
}
