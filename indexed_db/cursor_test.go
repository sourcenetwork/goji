//go:build js

package indexed_db

import (
	"sync"
	"syscall/js"
	"testing"

	"github.com/sourcenetwork/goji"
	"github.com/stretchr/testify/require"
)

func TestCursor(t *testing.T) {
	var req RequestValue[DatabaseValue]
	upgradeNeeded := goji.EventListener(func(event goji.EventValue) {
		req.Result().CreateObjectStore("authors", ObjectStoreOptions.WithKeyPath(js.ValueOf("name")))
	})
	defer upgradeNeeded.Release()

	req = Open(t.Name(), 1)
	req.EventTarget().AddEventListener(UpgradeNeededEvent, upgradeNeeded.Value)

	db, err := Await(req)
	require.NoError(t, err)
	defer db.Close()

	transaction := db.Transaction(js.ValueOf("authors"), TransactionModeReadWrite)
	defer transaction.Abort()

	store := transaction.ObjectStore("authors")

	_, err = Await(store.Add(map[string]any{"name": "bob"}))
	require.NoError(t, err)

	_, err = Await(store.Add(map[string]any{"name": "alice"}))
	require.NoError(t, err)

	var cursorWait sync.WaitGroup
	var cursorReq RequestValue[CursorValue]
	cursorSuccess := goji.EventListener(func(event goji.EventValue) {
		value := cursorReq.Result()
		if !js.Value(value).IsNull() {
			value.Continue()
		}
		cursorWait.Done()
	})
	defer cursorSuccess.Release()

	cursorWait.Add(3) // expect 3 iterations

	cursorReq = store.OpenCursor()
	cursorReq.EventTarget().AddEventListener(SuccessEvent, cursorSuccess.Value)

	cursorWait.Wait()
}
