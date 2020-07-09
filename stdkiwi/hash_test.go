// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi/values/hash"
)

func TestHash(t *testing.T) {
	store := newTestStore(t, hash.Type)
	s := store.Hash(testKey)

	// check that it does not panic
	s.Guard()

	// and the same should work with GuardE as well
	if err := s.GuardE(); err != nil {
		t.Errorf("GuardE threw an error for default testKey: %v", err)
	}

	// Insert 2 key-value pairs
	if err := s.Insert("a", "x"); err != nil {
		t.Errorf("could not Insert: %v", err)
	}

	if err := s.Insert("b", "y"); err != nil {
		t.Errorf("could not Insert: %v", err)
	}

	length, err := s.Len()
	if err != nil {
		t.Errorf("could not Len: %v", err)
	}

	if length != 2 {
		t.Errorf("expected length of set: 2; got %d", length)
	}

	toGet := []string{"a", "b"}
	values, err := s.Get(toGet...)
	if err != nil {
		t.Errorf("could not Get: %v", err)
	}

	if len(values) != length {
		t.Errorf("length expected: %d; got %d", length, len(values))
	}

	if values[0] != "x" || values[1] != "y" {
		t.Errorf("Get did not return correct value(s)")
	}

	ok, err := s.Has("b")
	if err != nil {
		t.Errorf("could not Has: %v", err)
	}

	if ok != true {
		t.Errorf("expected true; got false")
	}

	keys, err := s.Key()
	if err != nil {
		t.Errorf("could not Key: %v", err)
	}

	if len(keys) != length {
		t.Errorf("length expected: %d; got %d", length, len(keys))
	}

	if keys[0] != "a" || keys[1] != "b" {
		t.Errorf("Key did not return correct key(s)")
	}

	// remove two elements
	toRemove := []string{"a", "b"}
	err = s.Remove(toRemove...)
	if err != nil {
		t.Errorf("could not Remove: %v", err)
	}

	// now the map is empty
	cmap, err := s.Map()
	if err != nil {
		t.Errorf("could not Map: %v", err)
	}

	if len(cmap) != 0 {
		t.Errorf("length of map expected: 0; got %d", len(cmap))
	}

	// check guard for invalid key
	err = store.Str("randomKey").GuardE()
	if err == nil {
		t.Errorf("expected error while GuardE for invalid key; got nil")
	}
}
