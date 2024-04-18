//go:build js

package web_assembly

import (
	_ "embed"
	"syscall/js"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sourcenetwork/goji"
)

//go:embed test.wasm
var wasmBytes []byte

func TestCompile(t *testing.T) {
	src := goji.Uint8ArrayFromBytes(wasmBytes)

	valid := Validate(js.Value(src))
	require.True(t, valid)

	prom := Compile(js.Value(src))
	res, err := goji.Await(prom)
	require.NoError(t, err)
	require.Len(t, res, 1)
}

func TestInstantiate(t *testing.T) {
	src := goji.Uint8ArrayFromBytes(wasmBytes)
	importObject := js.ValueOf(map[string]any{})

	valid := Validate(js.Value(src))
	require.True(t, valid)

	prom := Instantiate(js.Value(src), importObject)
	res, err := goji.Await(prom)
	require.NoError(t, err)
	require.Len(t, res, 1)
}
