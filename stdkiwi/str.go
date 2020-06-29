// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import "github.com/sdslabs/kiwi/values/str"

// Str implements methods for str value type.
type Str struct {
	store *Store
	key   string
}

// Guard guards the keys with values of str type.
func (s *Str) Guard() {
	if err := s.GuardE(); err != nil {
		panic(err)
	}
}

// GuardE is same as Guard but does not panic, instead returns the error.
func (s *Str) GuardE() error { return s.store.guardValueE(str.Type, s.key) }

// Get gets the string stored corresponding to the key.
func (s *Str) Get() (string, error) {
	v, err := s.store.Do(s.key, str.Get)
	if err != nil {
		return "", err
	}

	getstr, ok := v.(string)
	if !ok {
		// This is very unexpected but it's better to handle error than to be sorry
		return "", newTypeErr(getstr, v)
	}

	return getstr, nil
}

// Update updates the string corresponding to the key.
func (s *Str) Update(update string) error {
	if _, err := s.store.Do(s.key, str.Update, update); err != nil {
		return err
	}

	return nil
}

// Clear empties the string corresponding to the key.
func (s *Str) Clear() error {
	return s.Update("")
}

// Len returns the length of the string.
func (s *Str) Len() (int, error) {
	getstr, err := s.Get()
	if err != nil {
		return 0, err
	}

	return len(getstr), nil
}

// Interface guard.
var _ Value = (*Str)(nil)
