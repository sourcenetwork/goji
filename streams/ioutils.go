//go:build js

package streams

import (
	"io"
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

var _ io.ReadCloser = (*reader)(nil)

type reader struct {
	read ReadableStreamBYOBReaderValue
	done bool
}

// NewReader wraps a ReadableStreamBYOBReaderValue into an io.ReadCloser.
func NewReader(read ReadableStreamBYOBReaderValue) io.ReadCloser {
	return &reader{read: read}
}

func (r *reader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.done {
		return 0, io.EOF
	}
	view := goji.Uint8Array.New(len(b))
	res, err := goji.Await(r.read.Read(js.Value(view)))
	if err != nil {
		return 0, err
	}
	copied := js.CopyBytesToGo(b, res[0].Get("value"))
	r.done = res[0].Get("done").Bool()
	return copied, nil
}

func (r *reader) Close() error {
	_, err := goji.Await(r.read.Cancel("user requested"))
	return err
}

var _ io.WriteCloser = (*writer)(nil)

type writer struct {
	write WritableStreamDefaultWriterValue
}

// NewWriter wraps a WritableStreamDefaultWriterValue into an io.WriteCloser.
func NewWriter(write WritableStreamDefaultWriterValue) io.WriteCloser {
	return &writer{write}
}

func (w *writer) Write(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	view := goji.Uint8ArrayFromBytes(b)
	_, err = goji.Await(w.write.Write(js.Value(view)))
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (w *writer) Close() error {
	_, err := goji.Await(w.write.Close())
	return err
}
