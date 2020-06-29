// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import "github.com/sdslabs/kiwi/values/list"

// List implements methods for list value type.
type List struct {
	store *Store
	key   string
}

// Guard guards the keys with values of str type.
func (l *List) Guard() {
	if err := l.GuardE(); err != nil {
		panic(err)
	}
}

// GuardE is same as Guard but does not panic, instead returns the error.
func (l *List) GuardE() error { return l.store.guardValueE(list.Type, l.key) }

// Get gets the string in the list at "index".
func (l *List) Get(index int) (string, error) {
	v, err := l.store.Do(l.key, list.Get, index)
	if err != nil {
		return "", err
	}

	getstr, ok := v.(string)
	if !ok {
		return "", newTypeErr(getstr, v)
	}

	return getstr, nil
}

// Set sets the string in the list at "index".
func (l *List) Set(index int, set string) error {
	if _, err := l.store.Do(l.key, list.Set, index, set); err != nil {
		return err
	}

	return nil
}

// Slice slices the list from start to end (end excluded).
func (l *List) Slice(start, end int) ([]string, error) {
	v, err := l.store.Do(l.key, list.Slice, start, end)
	if err != nil {
		return nil, err
	}

	slice, ok := v.([]string)
	if !ok {
		return nil, newTypeErr(slice, v)
	}

	return slice, nil
}

// Len gets the length of the list.
func (l *List) Len() (int, error) {
	v, err := l.store.Do(l.key, list.Len)
	if err != nil {
		return 0, err
	}

	length, ok := v.(int)
	if !ok {
		return 0, newTypeErr(length, v)
	}

	return length, nil
}

// Append appends the strings to the end of list.
func (l *List) Append(strings ...string) error {
	if len(strings) == 0 {
		return nil
	}

	// slice of strings to slice of interfaces
	ifaces := make([]interface{}, len(strings))
	for i := range strings {
		ifaces[i] = strings[i]
	}

	if _, err := l.store.Do(l.key, list.Append, ifaces...); err != nil {
		return err
	}

	return nil
}

// Pop pops off the last "n" elements from the end of the list.
func (l *List) Pop(n int) error {
	if n == 0 {
		return nil
	}

	if _, err := l.store.Do(l.key, list.Pop, n); err != nil {
		return err
	}

	return nil
}

// Remove removes the elem with the given index.
func (l *List) Remove(index int) error {
	if _, err := l.store.Do(l.key, list.Remove, index); err != nil {
		return err
	}

	return nil
}

// RemoveS removes the elem with the given value.
func (l *List) RemoveS(value string) error {
	if _, err := l.store.Do(l.key, list.Remove, value); err != nil {
		return err
	}

	return nil
}

// Find finds the index of the "value". Returns -1 if it does not exist.
func (l *List) Find(value string) (int, error) {
	v, err := l.store.Do(l.key, list.Find, value)
	if err != nil {
		return 0, err
	}

	idx, ok := v.(int)
	if !ok {
		return 0, newTypeErr(idx, v)
	}

	return idx, nil
}

// Interface guard.
var _ Value = (*List)(nil)
