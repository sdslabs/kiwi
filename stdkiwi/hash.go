// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"github.com/sdslabs/kiwi/values/hash"
	"github.com/sdslabs/kiwi/values/set"
)

// Hash implements methods for hash value type.
type Hash struct {
	store *Store
	key   string
}

// Guard guards the keys with values of str type.
func (h *Hash) Guard() {
	if err := h.GuardE(); err != nil {
		panic(err)
	}
}

// GuardE is same as Guard but does not panic, instead returns the error.
func (h *Hash) GuardE() error { return h.store.guardValueE(hash.Type, h.key) }

// Insert inserts the key-value pair in the hashmap.
func (h *Hash) Insert(key, value string) error {
	if _, err := h.store.Do(h.key, hash.Insert, key, value); err != nil {
		return err
	}

	return nil
}

// Remove removes the key-value pair(s) from the hashmap.
func (h *Hash) Remove(elements ...string) error {
	if len(elements) == 0 {
		return nil
	}

	ifaces := make([]interface{}, len(elements))
	for i := range elements {
		ifaces[i] = elements[i]
	}

	if _, err := h.store.Do(h.key, hash.Remove, ifaces...); err != nil {
		return err
	}

	return nil
}

// Has checks if key is present in the hashmap.
func (h *Hash) Has(key string) (bool, error) {
	v, err := h.store.Do(h.key, set.Has, key)
	if err != nil {
		return false, err
	}

	val, ok := v.(bool)
	if !ok {
		return false, newTypeErr(val, v)
	}

	return val, nil
}

// Len gets the length of the hashmap.
func (h *Hash) Len() (int, error) {
	v, err := h.store.Do(h.key, hash.Len)
	if err != nil {
		return 0, err
	}

	length, ok := v.(int)
	if !ok {
		return 0, newTypeErr(length, v)
	}

	return length, nil
}

// Get gets value(s) of the given key(s).
func (h *Hash) Get(keys ...string) ([]string, error) {
	if len(keys) == 0 {
		return nil, nil
	}

	ifaces := make([]interface{}, len(keys))
	for i := range keys {
		ifaces[i] = keys[i]
	}

	v, err := h.store.Do(h.key, hash.Get, ifaces...)
	if err != nil {
		return nil, err
	}

	values, ok := v.([]string)
	if !ok {
		return nil, newTypeErr(values, v)
	}

	return values, nil
}

// Keys gets all the keys of the hashmap.
func (h *Hash) Keys() ([]string, error) {
	v, err := h.store.Do(h.key, hash.Keys)
	if err != nil {
		return nil, err
	}

	keys, ok := v.([]string)
	if !ok {
		return nil, newTypeErr(keys, v)
	}

	return keys, nil
}

// Map gets a copy of the hashmap.
func (h *Hash) Map() (map[string]string, error) {
	v, err := h.store.Do(h.key, hash.Map)
	if err != nil {
		return nil, err
	}

	cmap, ok := v.(map[string]string)
	if !ok {
		return nil, newTypeErr(cmap, v)
	}

	return cmap, nil
}

// Interface guard.
var _ Value = (*Hash)(nil)
