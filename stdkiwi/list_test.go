// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi/values/list"
)

func TestList(t *testing.T) {
	store := newTestStore(t, list.Type)
	s := store.List(testKey)

	// check that it does not panic
	s.Guard()

	// and the same should work with GuardE as well
	if err := s.GuardE(); err != nil {
		t.Errorf("GuardE threw an error for default testKey: %v", err)
	}

	vals := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	if err := s.Append(vals...); err != nil {
		t.Errorf("could not Append: %v", err)
	}

	length, err := s.Len()
	if err != nil {
		t.Errorf("could not Len: %v", err)
	}

	if length != len(vals) {
		t.Errorf("expected length of slice: %d; got %d", len(vals), length)
	}

	slice, err := s.Slice(0, length)
	if err != nil {
		t.Errorf("could not Slice: %v", err)
	}

	if len(slice) != length {
		// weird flex but okay
		t.Errorf("slice length expected: %d; got %d", length, len(slice))
	}

	vals[2] = "C"

	err = s.Set(2, vals[2])
	if err != nil {
		t.Errorf("could not Set: %v", err)
	}

	getstr, err := s.Get(2)
	if err != nil {
		t.Errorf("could not Get: %v", err)
	}

	if getstr != vals[2] {
		t.Errorf("expected string from Get: %s; got %s", vals[2], getstr)
	}

	idx, err := s.Find(vals[2])
	if err != nil {
		t.Errorf("could not Find: %v", err)
	}

	if idx != 2 {
		t.Errorf("expected index of %q = 2; got %d", vals[2], idx)
	}

	// pop last 3 elements from the list
	err = s.Pop(3)
	if err != nil {
		t.Errorf("could not Pop: %v", err)
	}

	// remove first two elements
	err = s.Remove(0)
	if err != nil {
		t.Errorf("could not Remove: %v", err)
	}

	err = s.RemoveS("e")
	if err != nil {
		t.Errorf("could not RemoveS: %v", err)
	}

	// now what we should be left with is only "C"
	length, err = s.Len()
	if err != nil {
		t.Errorf("could not Len: %s", err)
	}

	if length != 3 {
		t.Errorf("expected length: 1; got %d", length)
	}

	slice, err = s.Slice(0, length)
	if err != nil {
		t.Errorf("could not Slice: %v", err)
	}

	if slice[0] != "b" || slice[1] != "C" || slice[2] != "d" {
		t.Errorf("expected list: %v; got %v", vals[1:4], slice)
	}

	// check guard for invalid key
	err = store.Str("randomKey").GuardE()
	if err == nil {
		t.Errorf("expected error while GuardE for invalid key; got nil")
	}
}
