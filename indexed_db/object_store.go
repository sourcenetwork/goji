//go:build js

package indexed_db

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
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
func (o ObjectStoreValue) Add(value js.Value, key js.Value) RequestValue {
	res := js.Value(o).Call("add", value, key)
	return RequestValue{goji.EventTargetValue(res)}
}

// Clear wraps the IDBObjectStore clear instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/clear
func (o ObjectStoreValue) Clear() RequestValue {
	res := js.Value(o).Call("clear")
	return RequestValue{goji.EventTargetValue(res)}
}

// Count wraps the IDBObjectStore count instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/count
func (o ObjectStoreValue) Count(query js.Value) RequestValue {
	res := js.Value(o).Call("count", query)
	return RequestValue{goji.EventTargetValue(res)}
}

// CreateIndex wraps the IDBObjectStore createIndex instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/createIndex
func (o ObjectStoreValue) CreateIndex(indexName, keyPath, options js.Value) IndexValue {
	res := js.Value(o).Call("createIndex", indexName, keyPath, options)
	return IndexValue(res)
}

// Delete wraps the IDBObjectStore delete instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/delete
func (o ObjectStoreValue) Delete(key js.Value) RequestValue {
	res := js.Value(o).Call("delete", key)
	return RequestValue{goji.EventTargetValue(res)}
}

// DeleteIndex wraps the IDBObjectStore deleteIndex instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/deleteIndex
func (o ObjectStoreValue) DeleteIndex(key js.Value) {
	js.Value(o).Call("deleteIndex", key)
}

// Get wraps the IDBObjectStore get instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/get
func (o ObjectStoreValue) Get(key js.Value) RequestValue {
	res := js.Value(o).Call("get", key)
	return RequestValue{goji.EventTargetValue(res)}
}

// GetAll wraps the IDBObjectStore getAll instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/getAll
func (o ObjectStoreValue) GetAll(key js.Value, count js.Value) RequestValue {
	res := js.Value(o).Call("getAll", key, count)
	return RequestValue{goji.EventTargetValue(res)}
}

// GetAllKeys wraps the IDBObjectStore getAllKeys instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/getAllKeys
func (o ObjectStoreValue) GetAllKeys(key js.Value, count js.Value) RequestValue {
	res := js.Value(o).Call("getAllKeys", key, count)
	return RequestValue{goji.EventTargetValue(res)}
}

// GetKey wraps the IDBObjectStore getKey instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/getKey
func (o ObjectStoreValue) GetKey(key js.Value) RequestValue {
	res := js.Value(o).Call("getKey", key)
	return RequestValue{goji.EventTargetValue(res)}
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
func (o ObjectStoreValue) OpenCursor(query, direction js.Value) RequestValue {
	res := js.Value(o).Call("openCursor", query, direction)
	return RequestValue{goji.EventTargetValue(res)}
}

// OpenKeyCursor wraps the IDBObjectStore openKeyCursor instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/openKeyCursor
func (o ObjectStoreValue) OpenKeyCursor(query, direction js.Value) RequestValue {
	res := js.Value(o).Call("openKeyCursor", query, direction)
	return RequestValue{goji.EventTargetValue(res)}
}

// Put wraps the IDBObjectStore put instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBObjectStore/put
func (o ObjectStoreValue) Put(item, key js.Value) RequestValue {
	res := js.Value(o).Call("put", item, key)
	return RequestValue{goji.EventTargetValue(res)}
}
