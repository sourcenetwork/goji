//go:build js

package indexed_db

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

const (
	// The blocked handler is executed when an open
	// connection to a database is blocking a
	// versionchange transaction on the same database.
	BlockedEvent = "blocked"
	// The upgradeneeded event is fired when an attempt
	// was made to open a database with a version number
	// higher than its current version.
	UpgradeNeededEvent = "upgradeneeded"
)

func init() {
	indexedDB = js.Global().Get("indexedDB")
}

var indexedDB js.Value

// Open wraps the IDBFactory open method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/open
func Open(name string, version uint) RequestValue[DatabaseValue] {
	res := indexedDB.Call("open", name, version)
	return RequestValue[DatabaseValue](res)
}

// DeleteDatabase wraps the IDBFactory deleteDatabase method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/deleteDatabase
func DeleteDatabase(name string) RequestValue[js.Value] {
	res := indexedDB.Call("deleteDatabase", name)
	return RequestValue[js.Value](res)
}

// Cmp wraps the IDBFactory cmp method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/cmp
func Cmp(first any, second any) int {
	return indexedDB.Call("cmp", first, second).Int()
}

// Databases wraps the IDBFactory databases method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBFactory/databases
func Databases() goji.PromiseValue {
	res := indexedDB.Call("databases")
	return goji.PromiseValue(res)
}
