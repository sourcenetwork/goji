//go:build js

package streams

import (
	"context"
	"io"
	"syscall/js"

	"github.com/sourcenetwork/goji"
)

var _ io.ReadCloser = (*Reader)(nil)

// Reader wraps a ReadableStreamBYOBReaderValue into an io.ReadCloser.
type Reader struct {
	read ReadableStreamBYOBReaderValue
	done bool
}

// NewReader returns a new Reader that reads from the provided ReadableStreamBYOBReaderValue.
func NewReader(read ReadableStreamBYOBReaderValue) *Reader {
	return &Reader{read: read}
}

func (r *Reader) Read(b []byte) (n int, err error) {
	return r.ReadContext(context.Background(), b)
}

func (r *Reader) ReadContext(ctx context.Context, b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	if r.done {
		return 0, io.EOF
	}
	view := goji.Uint8Array.New(len(b))
	res, err := goji.AwaitContext(ctx, r.read.Read(js.Value(view)))
	if err != nil {
		return 0, err
	}
	copied := js.CopyBytesToGo(b, res[0].Get("value"))
	r.done = res[0].Get("done").Bool()
	return copied, nil
}

func (r *Reader) Close() error {
	_, err := goji.Await(r.read.Cancel("user requested"))
	return err
}

var _ io.WriteCloser = (*Writer)(nil)

// Writer wraps a WritableStreamDefaultWriterValue into an io.WriteCloser.
type Writer struct {
	write WritableStreamDefaultWriterValue
}

// NewWriter returns a new writer that writes to the provided WritableStreamDefaultWriterValue.
func NewWriter(write WritableStreamDefaultWriterValue) *Writer {
	return &Writer{write}
}

func (w *Writer) Write(b []byte) (n int, err error) {
	return w.WriteContext(context.Background(), b)
}

func (w *Writer) WriteContext(ctx context.Context, b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, nil
	}
	view := goji.Uint8ArrayFromBytes(b)
	_, err = goji.AwaitContext(ctx, w.write.Write(js.Value(view)))
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (w *Writer) Close() error {
	_, err := goji.Await(w.write.Close())
	return err
}

func (w *Writer) Abort() error {
	_, err := goji.Await(w.write.Abort("user requested"))
	return err
}
