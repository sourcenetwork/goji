//go:build js

package indexed_db

import (
	"syscall/js"
)

const (
	// The cursor shows all records, including duplicates.
	// It starts at the lower bound of the key range and moves
	// upwards (monotonically increasing in the order of keys).
	CursorDirectionNext = "next"
	// The cursor shows all records, excluding duplicates.
	// If multiple records exist with the same key, only the
	// first one iterated is retrieved. It starts at the lower
	// bound of the key range and moves upwards.
	CursorDirectionNextUnique = "nextunique"
	// The cursor shows all records, including duplicates.
	// It starts at the upper bound of the key range and moves
	// downwards (monotonically decreasing in the order of keys).
	CursorDirectionPrev = "prev"
	// The cursor shows all records, excluding duplicates.
	// If multiple records exist with the same key, only the
	// first one iterated is retrieved. It starts at the upper
	// bound of the key range and moves downwards.
	CursorDirectionPrevUnique = "prevunique"
)

// CursorValue is an instance of IDBCursor.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor
type CursorValue js.Value

// Direction returns the IDBCursor direction property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/direction
func (c CursorValue) Direction() string {
	return js.Value(c).Get("direction").String()
}

// Key returns the IDBCursor key property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/key
func (c CursorValue) Key() js.Value {
	return js.Value(c).Get("key")
}

// PrimaryKey returns the IDBCursor primaryKey property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/primaryKey
func (c CursorValue) PrimaryKey() js.Value {
	return js.Value(c).Get("primaryKey")
}

// Request returns the IDBCursor request property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/request
func (c CursorValue) Request() RequestValue[js.Value] {
	res := js.Value(c).Get("request")
	return RequestValue[js.Value](res)
}

// Source returns the IDBCursor request property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/source
func (c CursorValue) Source() js.Value {
	return js.Value(c).Get("source")
}

// Advance wraps the IDBCursor advance instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/advance
func (c CursorValue) Advance(count uint) {
	js.Value(c).Call("advance", count)
}

// Continue wraps the IDBCursor continue instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/continue
func (c CursorValue) Continue(args ...any) {
	js.Value(c).Call("continue", args...)
}

// ContinuePrimaryKey wraps the IDBCursor continuePrimaryKey instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/continuePrimaryKey
func (c CursorValue) ContinuePrimaryKey(key any, primaryKey any) {
	js.Value(c).Call("continuePrimaryKey", key, primaryKey)
}

// Delete wraps the IDBCursor delete instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/delete
func (c CursorValue) Delete() RequestValue[js.Value] {
	res := js.Value(c).Call("delete")
	return RequestValue[js.Value](res)
}

// Update wraps the IDBCursor update instance method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursor/update
func (c CursorValue) Update(value any) RequestValue[js.Value] {
	res := js.Value(c).Call("update", value)
	return RequestValue[js.Value](res)
}

// Value returns the IDBCursorWithValue value property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/IDBCursorWithValue/value
func (c CursorValue) Value() js.Value {
	return js.Value(c).Get("value")
}
