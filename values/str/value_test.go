// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package str

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	expectedJSON := json.RawMessage(fmt.Sprintf(`"%s"`, orig)) // "string"

	obj, err := store.ToJSON(key)
	if err != nil {
		t.Errorf("ToJSON returned unexpected error: %v", err)
	}
	if !bytes.Equal(obj, expectedJSON) {
		t.Errorf("expected JSON:\n%s; got:\n%s", string(expectedJSON), string(obj))
	}

	// add a new key and initiate that key FromJSON and check if the value equals
	// by invoking the "GET" action.
	newKey := "xyz"
	err = store.AddKey(newKey, Type)
	if err != nil {
		t.Fatalf("cannot add new key to the store: %v", err)
	}

	err = store.FromJSON(newKey, obj)
	if err != nil {
		t.Errorf("FromJSON returned unexpected error: %v", err)
	}
	v, err := store.Do(key, Get)
	if err != nil {
		t.Errorf("cannot GET from the store: %v", err)
	}
	str, ok = v.(string)
	if !ok {
		t.Errorf("GET did not return a string")
	}
	if str != orig {
		t.Errorf("expected string FromJSON: %q; got %q", orig, str)
	}
}
