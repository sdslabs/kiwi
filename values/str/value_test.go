// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package str

import (
	"testing"

	"github.com/sdslabs/kiwi"
)

func TestValue(t *testing.T) {
	key := "abc"

	schema := kiwi.Schema{key: Type}
	store, err := kiwi.NewStoreFromSchema(schema)
	if err != nil {
		t.Fatalf("error while creating store: %v", err)
	}

	orig := "def"

	ret, err := store.Do(key, Update, orig)
	if err != nil {
		t.Errorf("error while updating key: %v", err)
	}
	str, ok := ret.(string)
	if !ok {
		t.Errorf("update did not return string: got a %T", ret)
	}
	if str != orig {
		t.Errorf("update did not return correct string: got %s", str)
	}

	ret, err = store.Do(key, Get)
	if err != nil {
		t.Errorf("error while getting key: %v", err)
	}
	str, ok = ret.(string)
	if !ok {
		t.Errorf("get did not return string: got a %T", ret)
	}
	if str != orig {
		t.Errorf("get did not return correct string: got %s", str)
	}

	_, err = store.Do(key, Update)
	if err == nil {
		t.Errorf("did not get expected error when updating with invalid (0 num) params")
	}
}
