// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi/values/zset"
)

func TestZSet(t *testing.T) {
	store := newTestStore(t, zset.Type)
	s := store.Zset(testKey)

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
		t.Errorf("expected length of zset: %d; got %d", len(vals), length)
	}

	// updating score to some non zero values.

	if err = s.Increment("a", 5); err != nil {
		t.Errorf("could not increment: %v", err)
	}

	if err = s.Increment("f", -5); err != nil {
		t.Errorf("could not Increment: %v", err)
	}

	score, err := s.Get("b")
	if err != nil {
		t.Errorf("could not Get: %v", err)
	}

	if score != 0 {
		t.Errorf("expected score: %d; got: %d", 0, score)
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

	if minstr != "f" {
		t.Errorf("expected element: %q; got: %q", "f", minstr)
	}
}
