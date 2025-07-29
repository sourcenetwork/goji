//go:build js

package streams

import (
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

// ReadableStreamValue is an instance of a ReadableStream.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream
type ReadableStreamValue js.Value

// Locked returns the ReadableStream.locked property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/locked
func (v ReadableStreamValue) Locked() bool {
	return js.Value(v).Get("locked").Bool()
}

// Locked calls the ReadableStream.cancel method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/cancel
func (v ReadableStreamValue) Cancel(reason string) goji.PromiseValue {
	res := js.Value(v).Call("cancel", reason)
	return goji.PromiseValue(res)
}

// GetDefaultReader returns a ReadableStreamDefaultReaderValue.
func (v ReadableStreamValue) GetDefaultReader() ReadableStreamDefaultReaderValue {
	res := js.Value(v).Call("getReader")
	return ReadableStreamDefaultReaderValue(res)
}

// GetBYOBReader returns a ReadableStreamBYOBReaderValue.
func (v ReadableStreamValue) GetBYOBReader() ReadableStreamBYOBReaderValue {
	res := js.Value(v).Call("getReader", map[string]any{"mode": "byob"})
	return ReadableStreamBYOBReaderValue(res)
}

// PipeThrough calls the ReadableStream.pipeThrough method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/pipeThrough
func (v ReadableStreamValue) PipeThrough(transformStream js.Value, opts ...readableStreamPipeOption) ReadableStreamValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(v).Call("pipeThrough", transformStream, options)
		return ReadableStreamValue(res)

	default:
		res := js.Value(v).Call("pipeThrough", transformStream)
		return ReadableStreamValue(res)
	}
}

// PipeTo calls the ReadableStream.pipeTo method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/pipeTo
func (v ReadableStreamValue) PipeTo(destination js.Value, opts ...readableStreamPipeOption) goji.PromiseValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(v).Call("pipeTo", destination, options)
		return goji.PromiseValue(res)

	default:
		res := js.Value(v).Call("pipeTo", destination)
		return goji.PromiseValue(res)
	}
}

// Tee calls the ReadableStream.tee method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/tee
func (v ReadableStreamValue) Tee() (ReadableStreamValue, ReadableStreamValue) {
	res := js.Value(v).Call("tee")
	one := ReadableStreamValue(res.Index(0))
	two := ReadableStreamValue(res.Index(1))
	return one, two
}

// ReadableStreamPipeOptions sets options for ReadableStream.pipeThrough
// and ReadableStream.pipeTo calls.
var ReadableStreamPipeOptions = &readableStreamPipeOptions{}

type readableStreamPipeOptions struct{}

type readableStreamPipeOption func(opts js.Value)

// WithPreventClose sets the preventClose option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/pipeThrough#preventclose
func (e readableStreamPipeOptions) WithPreventClose(value bool) readableStreamPipeOption {
	return func(opts js.Value) {
		opts.Set("preventClose", value)
	}
}

// WithPreventAbort sets the preventAbort option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/pipeThrough#preventabort
func (e readableStreamPipeOptions) WithPreventAbort(value bool) readableStreamPipeOption {
	return func(opts js.Value) {
		opts.Set("preventAbort", value)
	}
}

// WithPreventCancel sets the preventCancel option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/pipeThrough#preventcancel
func (e readableStreamPipeOptions) WithPreventCancel(value bool) readableStreamPipeOption {
	return func(opts js.Value) {
		opts.Set("preventCancel", value)
	}
}

// WithSignal sets the signal option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStream/pipeThrough#signal
func (e readableStreamPipeOptions) WithSignal(value js.Value) readableStreamPipeOption {
	return func(opts js.Value) {
		opts.Set("signal", value)
	}
}

// ReadableStreamDefaultReaderValue is an instance of a ReadableStreamDefaultReader.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamDefaultReader
type ReadableStreamDefaultReaderValue js.Value

// Closed returns the ReadableStreamDefaultReader.closed property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamDefaultReader/closed
func (v ReadableStreamDefaultReaderValue) Closed() goji.PromiseValue {
	res := js.Value(v).Get("closed")
	return goji.PromiseValue(res)
}

// Cancel calls the ReadableStreamDefaultReader.cancel method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamDefaultReader/cancel
func (v ReadableStreamDefaultReaderValue) Cancel(reason string) goji.PromiseValue {
	res := js.Value(v).Call("cancel", reason)
	return goji.PromiseValue(res)
}

// Cancel calls the ReadableStreamDefaultReader.read method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamDefaultReader/read
func (v ReadableStreamDefaultReaderValue) Read() goji.PromiseValue {
	res := js.Value(v).Call("read")
	return goji.PromiseValue(res)
}

// ReleaseLock calls the ReadableStreamDefaultReader.releaseLock method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamDefaultReader/releaseLock
func (v ReadableStreamDefaultReaderValue) ReleaseLock() {
	js.Value(v).Call("releaseLock")
}

// ReadableStreamBYOBReaderValue is an instance of a ReadableStreamBYOBReader.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamBYOBReader
type ReadableStreamBYOBReaderValue js.Value

// Closed returns the ReadableStreamBYOBReader.closed property.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamBYOBReader/closed
func (v ReadableStreamBYOBReaderValue) Closed() goji.PromiseValue {
	res := js.Value(v).Get("closed")
	return goji.PromiseValue(res)
}

// Closed calls the ReadableStreamBYOBReader.cancel method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamBYOBReader/cancel
func (v ReadableStreamBYOBReaderValue) Cancel(reason string) goji.PromiseValue {
	res := js.Value(v).Call("cancel", reason)
	return goji.PromiseValue(res)
}

// Read calls the ReadableStreamBYOBReader.read method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamBYOBReader/read
func (v ReadableStreamBYOBReaderValue) Read(view js.Value, opts ...byobReaderReadOption) goji.PromiseValue {
	switch {
	case len(opts) > 0:
		options := js.ValueOf(map[string]any{})
		for _, opt := range opts {
			opt(options)
		}
		res := js.Value(v).Call("read", view, options)
		return goji.PromiseValue(res)

	default:
		res := js.Value(v).Call("read", view)
		return goji.PromiseValue(res)
	}
}

// ReleaseLock calls the ReadableStreamBYOBReader.releaseLock method.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamBYOBReader/releaseLock
func (v ReadableStreamBYOBReaderValue) ReleaseLock() {
	js.Value(v).Call("releaseLock")
}

// BYOBReaderReadOptions sets options for ReadableStreamBYOBReader.read calls.
var BYOBReaderReadOptions = &byobReaderReadOptions{}

type byobReaderReadOptions struct{}

type byobReaderReadOption func(opts js.Value)

// WithMin sets the min option.
//
// https://developer.mozilla.org/en-US/docs/Web/API/ReadableStreamBYOBReader/read#min
func (e byobReaderReadOptions) WithMin(min int) byobReaderReadOption {
	return func(opts js.Value) {
		opts.Set("min", min)
	}
}
