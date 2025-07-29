//go:build js

package streams

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

// WritableStreamValue is an instance of WritableStream.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStream
type WritableStreamValue js.Value

// Locked returns the WritableStream.locked property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStream/locked
func (v WritableStreamValue) Locked() bool {
	return js.Value(v).Get("locked").Bool()
}

// Abort calls the WritableStream.abort method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStream/abort
func (v WritableStreamValue) Abort(reason string) goji.PromiseValue {
	res := js.Value(v).Call("abort", reason)
	return goji.PromiseValue(res)
}

// Close calls the WritableStream.close method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStream/close
func (v WritableStreamValue) Close() goji.PromiseValue {
	res := js.Value(v).Call("close")
	return goji.PromiseValue(res)
}

// GetWriter calls the WritableStream.getWriter method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStream/getWriter
func (v WritableStreamValue) GetWriter() WritableStreamDefaultWriterValue {
	res := js.Value(v).Call("getWriter")
	return WritableStreamDefaultWriterValue(res)
}

// WritableStreamDefaultWriterValue is an instance of WritableStreamDefaultWriter.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter
type WritableStreamDefaultWriterValue js.Value

// Closed returns the WritableStreamDefaultWriter.closed property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter/closed
func (v WritableStreamDefaultWriterValue) Closed() goji.PromiseValue {
	res := js.Value(v).Get("closed")
	return goji.PromiseValue(res)
}

// DesiredSize returns the WritableStreamDefaultWriter.desiredSize property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter/desiredSize
func (v WritableStreamDefaultWriterValue) DesiredSize() int {
	return js.Value(v).Get("desiredSize").Int()
}

// Ready returns the WritableStreamDefaultWriter.ready property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter/ready
func (v WritableStreamDefaultWriterValue) Ready() goji.PromiseValue {
	res := js.Value(v).Get("ready")
	return goji.PromiseValue(res)
}

// Abort calls the WritableStreamDefaultWriter.abort method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter/abort
func (v WritableStreamDefaultWriterValue) Abort(reason string) goji.PromiseValue {
	res := js.Value(v).Call("abort", reason)
	return goji.PromiseValue(res)
}

// Close calls the WritableStreamDefaultWriter.close method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter/close
func (v WritableStreamDefaultWriterValue) Close() goji.PromiseValue {
	res := js.Value(v).Call("close")
	return goji.PromiseValue(res)
}

// ReleaseLock calls the WritableStreamDefaultWriter.releaseLock method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter/releaseLock
func (v WritableStreamDefaultWriterValue) ReleaseLock() {
	js.Value(v).Call("releaseLock")
}

// ReleaseLock calls the WritableStreamDefaultWriter.write method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WritableStreamDefaultWriter/write
func (v WritableStreamDefaultWriterValue) Write(chunk js.Value) goji.PromiseValue {
	res := js.Value(v).Call("write", chunk)
	return goji.PromiseValue(res)
}
