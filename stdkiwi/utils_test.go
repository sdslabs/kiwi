// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi"
)

// testKey is the key to be used while creating the store.
const testKey = "testKey"

// newTestStore creates a new store for testing the stdkiwi methods.
func newTestStore(t *testing.T, valType kiwi.ValueType) *Store {
	store, err := NewStoreFromSchema(kiwi.Schema{testKey: valType})
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	return store
}
