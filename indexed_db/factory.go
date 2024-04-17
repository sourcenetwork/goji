//go:build js

package indexed_db

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

func init() {
	indexedDB = js.Global().Get("indexedDB")
}

var indexedDB js.Value

// Open wraps the IDBFactory open method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/open
func Open(name string, version uint) RequestValue {
	res := indexedDB.Call("open", name, version)
	return RequestValue(res)
}

// DeleteDatabase wraps the IDBFactory deleteDatabase method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/deleteDatabase
func DeleteDatabase(name string, options js.Value) RequestValue {
	res := indexedDB.Call("deleteDatabase", name, options)
	return RequestValue(res)
}

// Cmp wraps the IDBFactory cmp method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/cmp
func Cmp(first, second js.Value) int {
	return indexedDB.Call("cmp", first, second).Int()
}

// Databases wraps the IDBFactory databases method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/databases
func Databases() goji.PromiseValue {
	res := indexedDB.Call("databases")
	return goji.PromiseValue(res)
}
