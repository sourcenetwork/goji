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
func (i IndexValue) Count(key js.Value) RequestValue {
	res := js.Value(i).Call("count", key)
	return RequestValue(res)
}

// Get wraps the IDBIndex get instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/get
func (i IndexValue) Get(key js.Value) RequestValue {
	res := js.Value(i).Call("get", key)
	return RequestValue(res)
}

// GetAll wraps the IDBIndex getAll instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/getAll
func (i IndexValue) GetAll(query, count js.Value) RequestValue {
	res := js.Value(i).Call("getAll", query, count)
	return RequestValue(res)
}

// GetAllKeys wraps the IDBIndex getAllKeys instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/getAllKeys
func (i IndexValue) GetAllKeys(query, count js.Value) RequestValue {
	res := js.Value(i).Call("getAllKeys", query, count)
	return RequestValue(res)
}

// GetKey wraps the IDBIndex getKey instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/getKey
func (i IndexValue) GetKey(key js.Value) RequestValue {
	res := js.Value(i).Call("getKey", key)
	return RequestValue(res)
}

// OpenCursor wraps the IDBIndex openCursor instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/openCursor
func (i IndexValue) OpenCursor(key, direction js.Value) RequestValue {
	res := js.Value(i).Call("openCursor", key, direction)
	return RequestValue(res)
}

// OpenKeyCursor wraps the IDBIndex openKeyCursor instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBIndex/openKeyCursor
func (i IndexValue) OpenKeyCursor(key, direction js.Value) RequestValue {
	res := js.Value(i).Call("openKeyCursor", key, direction)
	return RequestValue(res)
}
