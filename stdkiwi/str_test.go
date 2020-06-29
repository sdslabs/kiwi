// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi/values/str"
)

func TestStr(t *testing.T) {
	store := newTestStore(t, str.Type)
	s := store.Str(testKey)

	// check that it does not panic
	s.Guard()

	// and the same should work with GuardE as well
	if err := s.GuardE(); err != nil {
		t.Errorf("GuardE threw an error for default testKey: %v", err)
	}

	randomWords := "random words"

	if err := s.Update(randomWords); err != nil {
		t.Errorf("could not Update the string: %v", err)
	}

	getstr, err := s.Get()
	if err != nil {
		t.Errorf("could not Get the string: %v", err)
	}

	if getstr != randomWords {
		t.Errorf("expected Get string %q; got %q", randomWords, getstr)
	}

	l, err := s.Len()
	if err != nil {
		t.Errorf("could not Len the string: %v", err)
	}

	if l != len(randomWords) {
		t.Errorf("expected Len %d; got %d", len(randomWords), l)
	}

	err = s.Clear()
	if err != nil {
		t.Errorf("could not Clear the string: %v", err)
	}

	getstr, err = s.Get()
	if err != nil {
		t.Errorf("could not Get the string: %v", err)
	}

	if getstr != "" {
		t.Errorf("expected Get string \"\" (empty string); got %q", getstr)
	}

	// check guard for invalid key
	err = store.Str("randomKey").GuardE()
	if err == nil {
		t.Errorf("expected error while GuardE for invalid key; got nil")
	}
}
