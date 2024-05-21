//go:build js

package indexed_db

import (
	"syscall/js"
	"testing"

	"github.com/sourcenetwork/goji"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestObjectStoreAdd(t *testing.T) {
	var req RequestValue[DatabaseValue]
	upgradeNeeded := goji.EventListener(func(event goji.EventValue) {
		req.Result().CreateObjectStore("authors")
	})
	defer upgradeNeeded.Release()

	req = Open(t.Name(), 1)
	req.EventTarget().AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	db, err := Await(req)
	require.NoError(t, err)
	defer db.Close()

	key := js.ValueOf(1)
	val := js.ValueOf(map[string]any{"name": "bob"})

	transaction := db.Transaction(js.ValueOf("authors"), TransactionModeReadWrite)
	defer transaction.Abort()

	store := transaction.ObjectStore("authors")

	actualKey, err := Await(store.Add(val, key))
	require.NoError(t, err)
	assert.Equal(t, key, actualKey)

	actualVal, err := Await(store.Get(key))
	require.NoError(t, err)
	assert.Equal(t, "bob", actualVal.Get("name").String())
}
