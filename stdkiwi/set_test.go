// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi/values/set"
)

func TestSet(t *testing.T) {
	store := newTestStore(t, set.Type)
	s := store.Set(testKey)

	// check that it does not panic
	s.Guard()

	// and the same should work with GuardE as well
	if err := s.GuardE(); err != nil {
		t.Errorf("GuardE threw an error for default testKey: %v", err)
	}

	vals := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	if err := s.Insert(vals...); err != nil {
		t.Errorf("could not Insert: %v", err)
	}

	length, err := s.Len()
	if err != nil {
		t.Errorf("could not Len: %v", err)
	}

	if length != len(vals) {
		t.Errorf("expected length of set: %d; got %d", len(vals), length)
	}

	elems, err := s.Get()
	if err != nil {
		t.Errorf("could not Get: %v", err)
	}

	if len(elems) != length {
		t.Errorf("length expected: %d; got %d", length, len(elems))
	}

	ok, err := s.Has(vals[2])
	if err != nil {
		t.Errorf("could not Has: %v", err)
	}

	if ok != true {
		t.Errorf("expected true; got false")
	}

	// remove two elements
	toRemove := []string{"a", "b"}
	err = s.Remove(toRemove...)
	if err != nil {
		t.Errorf("could not Remove: %v", err)
	}

	// now length should reduced by 2
	length, err = s.Len()
	if err != nil {
		t.Errorf("could not Len: %s", err)
	}

	if length != 6 {
		t.Errorf("expected length: 6; got %d", length)
	}

	//checking for removed element
	ok, err = s.Has("b")
	if err != nil {
		t.Errorf("could not Has: %v", err)
	}

	if ok != false {
		t.Errorf("expected false; got true")
	}

	// check guard for invalid key
	err = store.Str("randomKey").GuardE()
	if err == nil {
		t.Errorf("expected error while GuardE for invalid key; got nil")
	}
}
