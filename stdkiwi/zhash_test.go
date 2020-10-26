// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi/values/zhash"
)

func TestZHash(t *testing.T) {
	store := newTestStore(t, zhash.Type)
	s := store.Zhash(testKey)

	// check that it does not panic
	s.Guard()

	// and the same should work with GuardE as well
	if err := s.GuardE(); err != nil {
		t.Errorf("GuardE threw an error for default testKey: %v", err)
	}

	if err := s.Insert("a", "b"); err != nil {
		t.Errorf("could not Insert: %v", err)
	}
	if err := s.Insert("c", "d"); err != nil {
		t.Errorf("could not Insert: %v", err)
	}

	if err := s.Set("a", "e"); err != nil {
		t.Errorf("could not Update: %v", err)
	}

	length, err := s.Len()
	if err != nil {
		t.Errorf("could not Len: %v", err)
	}

	if length != 2 {
		t.Errorf("expected length of zset: %d; got %d", 2, length)
	}

	// updating score to some non zero values.

	if err = s.Increment("a", 5); err != nil {
		t.Errorf("could not increment: %v", err)
	}

	item, err := s.Get("c")
	if err != nil {
		t.Errorf("could not Get: %v", err)
	}

	if item.Score != 0 {
		t.Errorf("expected score: %d; got: %d", 0, item.Score)
	}

	if item.Value != "d" {
		t.Errorf("expected value: d; got %v", item.Value)
	}

	maxstr, err := s.PeekMax()
	if err != nil {
		t.Errorf("could not PeekMax: %v", err)
	}

	if maxstr != "a" {
		t.Errorf("expected element: %q; got: %q", "a", maxstr)
	}

	minstr, err := s.PeekMin()
	if err != nil {
		t.Errorf("could not PeekMin: %v", err)
	}

	if minstr != "c" {
		t.Errorf("expected element: %q; got: %q", "f", minstr)
	}

	if err := s.Remove("a", "c"); err != nil {
		t.Errorf("Unexpected error when removing elements: %v", err)
	}
}
