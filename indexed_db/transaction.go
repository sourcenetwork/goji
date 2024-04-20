//go:build js

package indexed_db

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

const (
	TransactionModeReadOnly       = "readonly"
	TransactionModeReadWrite      = "readwrite"
	TransactionModeReadWriteFlush = "readwriteflush"
	TransactionDurabilityDefault  = "default"
	TransactionDurabilityStrict   = "strict"
	TransactionDurabilityRelaxed  = "relaxed"
)

// TransactionValue is an instance of IDBTransaction.
type TransactionValue goji.EventTargetValue

// DB returns the IDBTransaction db property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/db
func (t TransactionValue) DB() DatabaseValue {
	res := js.Value(t).Get("db")
	return DatabaseValue{goji.EventTargetValue(res)}
}

// Durability returns the IDBTransaction durability property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/durability
func (t TransactionValue) Durability() string {
	return js.Value(t).Get("durability").String()
}

// Error returns the IDBTransaction error property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/error
func (t TransactionValue) Error() js.Value {
	return js.Value(t).Get("error")
}

// String returns the IDBTransaction mode property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/mode
func (t TransactionValue) Mode() string {
	return js.Value(t).Get("mode").String()
}

// ObjectStoreNames returns the IDBTransaction objectStoreNames property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/ObjectStoreNames
func (t TransactionValue) ObjectStoreNames() js.Value {
	return js.Value(t).Get("objectStoreNames")
}

// Abort wraps the IDBTransaction abort instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/abort
func (t TransactionValue) Abort() {
	js.Value(t).Call("abort")
}

// Commit wraps the IDBTransaction commit instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/commit
func (t TransactionValue) Commit() {
	js.Value(t).Call("commit")
}

// ObjectStore wraps the IDBTransaction objectStore instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBTransaction/objectStore
func (t TransactionValue) ObjectStore(name string) ObjectStoreValue {
	res := js.Value(t).Call("objectStore", name)
	return ObjectStoreValue(res)
}
