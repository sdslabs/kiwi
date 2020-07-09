// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import "github.com/sdslabs/kiwi/values/set"

// Set implements methods for set value type.
type Set struct {
	store *Store
	key   string
}

// Guard guards the keys with values of str type.
func (s *Set) Guard() {
	if err := s.GuardE(); err != nil {
		panic(err)
	}
}

// GuardE is same as Guard but does not panic, instead returns the error.
func (s *Set) GuardE() error { return s.store.guardValueE(set.Type, s.key) }

// Insert inserts the elements to the set.
func (s *Set) Insert(elements ...string) error {
	if len(elements) == 0 {
		return nil
	}

	ifaces := make([]interface{}, len(elements))
	for i := range elements {
		ifaces[i] = elements[i]
	}

	if _, err := s.store.Do(s.key, set.Insert, ifaces...); err != nil {
		return err
	}

	return nil
}

// Remove removes the elements from the set.
func (s *Set) Remove(elements ...string) error {
	if len(elements) == 0 {
		return nil
	}

	ifaces := make([]interface{}, len(elements))
	for i := range elements {
		ifaces[i] = elements[i]
	}

	if _, err := s.store.Do(s.key, set.Remove, ifaces...); err != nil {
		return err
	}

	return nil
}

// Has checks if element is present in the set.
func (s *Set) Has(element string) (bool, error) {
	v, err := s.store.Do(s.key, set.Has, element)
	if err != nil {
		return false, err
	}

	val, ok := v.(bool)
	if !ok {
		return false, newTypeErr(val, v)
	}

	return val, nil
}

// Len gets the length of the set.
func (s *Set) Len() (int, error) {
	v, err := s.store.Do(s.key, set.Len)
	if err != nil {
		return 0, err
	}

	length, ok := v.(int)
	if !ok {
		return 0, newTypeErr(length, v)
	}

	return length, nil
}

// Get gets all the elements of the set.
func (s *Set) Get() ([]string, error) {
	v, err := s.store.Do(s.key, set.Get)
	if err != nil {
		return nil, err
	}

	elems, ok := v.([]string)
	if !ok {
		return nil, newTypeErr(elems, v)
	}

	return elems, nil
}

// Interface guard.
var _ Value = (*Set)(nil)
