//go:build js

package goji

import (
	_ "embed"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test.wasm
var wasmBytes []byte

func TestWebAssemblyCompile(t *testing.T) {
	src := Uint8ArrayFromBytes(wasmBytes)

	valid := WebAssembly.Validate(js.Value(src))
	require.True(t, valid)

	prom := WebAssembly.Compile(js.Value(src))
	res, err := Await(prom)
	require.NoError(t, err)
	require.Len(t, res, 1)
}

func TestWebAssemblyInstantiate(t *testing.T) {
	src := Uint8ArrayFromBytes(wasmBytes)
	importObject := js.ValueOf(map[string]any{})

	valid := WebAssembly.Validate(js.Value(src))
	require.True(t, valid)

	prom := WebAssembly.Instantiate(js.Value(src), importObject)
	res, err := Await(prom)
	require.NoError(t, err)
	require.Len(t, res, 1)
}
