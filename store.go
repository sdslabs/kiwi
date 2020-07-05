// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package kiwi

import (
	"fmt"
	"strings"
	"sync"
)

// Various errors related to keys.
var (
	ErrKeyExists   = fmt.Errorf("key already exists")
	ErrKeyNotExist = fmt.Errorf("key does not exist")
)

// newKeyErr creates a new err with related key.
func newKeyErr(err error, key string) error {
	return fmt.Errorf("%w: %s", err, key)
}

// Store is the main element that contains and manages all the key value pairs.
type Store struct {
	kv map[string]valWrapper
	mu sync.RWMutex
}

// NewStore creates an empty store without any key value pairs initialized.
func NewStore() *Store {
	return &Store{
		kv: make(map[string]valWrapper),
		mu: sync.RWMutex{},
	}
}

// NewStoreFromSchema creates a new store from the provided schema.
func NewStoreFromSchema(schema Schema) (*Store, error) {
	s := NewStore()

	for key, val := range schema {
		if err := s.AddKey(key, val); err != nil {
			return nil, err
		}
	}

	return s, nil
}

// KeyExists tells if the key exists or not.
func (s *Store) KeyExists(key string) bool {
	return s.keyExists(key) == nil
}

// AddKey adds a new key to the store. It throws an error if the key already exists.
func (s *Store) AddKey(key string, typ ValueType) error {
	if err := s.keyNotExist(key); err != nil {
		return err
	}

	v, err := newValue(typ)
	if err != nil {
		return err
	}

	s.mu.Lock()
	s.kv[key] = valWrapper{val: v, mu: &sync.RWMutex{}}
	s.mu.Unlock()
	return nil
}

// UpdateKey updates the key if it exists. Throws an error if it doesn't.
func (s *Store) UpdateKey(key string, typ ValueType) error {
	if err := s.keyExists(key); err != nil {
		return err
	}

	v, err := newValue(typ)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	old := s.kv[key]
	old.mu.Lock()
	defer old.mu.Unlock()

	s.kv[key] = valWrapper{val: v, mu: &sync.RWMutex{}}

	return nil
}

// DeleteKey deletes the key if it exists. Throws an error if it doesn't.
func (s *Store) DeleteKey(key string) error {
	if err := s.keyExists(key); err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	old := s.kv[key]
	old.mu.Lock()
	defer old.mu.Unlock()

	delete(s.kv, key)

	return nil
}

// GetValueType returns the type of value corresponding to the key.
func (s *Store) GetValueType(key string) (ValueType, error) {
	if err := s.keyExists(key); err != nil {
		return "", err
	}

	s.mu.RLock()
	v := s.kv[key]
	s.mu.RUnlock()

	v.mu.RLock()
	defer v.mu.RUnlock()

	return v.val.Type(), nil
}

// GetSchema returns the schema of the store.
func (s *Store) GetSchema() Schema {
	s.mu.RLock()
	defer s.mu.RUnlock()

	schema := make(Schema)
	for k := range s.kv {
		schema[k] = s.kv[k].val.Type()
	}

	return schema
}

// Do executes the action for the value associated with the value.
func (s *Store) Do(key string, action Action, params ...interface{}) (interface{}, error) {
	if err := s.keyExists(key); err != nil {
		return nil, err
	}

	s.mu.RLock()
	v := s.kv[key]
	s.mu.RUnlock()

	v.mu.Lock()
	defer v.mu.Unlock()

	doFunc, ok := v.val.DoMap()[action]
	if !ok {
		return nil, fmt.Errorf("%w: %v", ErrInvalidAction, action)
	}

	res, err := doFunc(params...)
	return res, err
}

// keyExists checks if the key exists in the store. Throws an error if it doesn't.
func (s *Store) keyExists(key string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.kv[key]; !ok {
		return newKeyErr(ErrKeyNotExist, key)
	}

	return nil
}

// keyExists checks if the key does not exist in the store. Throws an error if it does.
func (s *Store) keyNotExist(key string) error {
	// no need to lock since the function already does handle that.
	if err := s.keyExists(key); err == nil {
		return newKeyErr(ErrKeyExists, key)
	}

	return nil
}

// Schema contains the value types corresponding to their keys.
type Schema map[string]ValueType

// String implements the fmt.Stringer interface.
func (s Schema) String() string {
	strs := make([]string, len(s)+2)

	i := 1
	for key := range s {
		strs[i] = fmt.Sprintf("\t%s: %s", key, s[key])
		i++
	}

	strs[0], strs[len(s)+1] = "{", "}"
	return strings.Join(strs, "\n")
}

// valWrapper contains the value as well as it's mutex.
type valWrapper struct {
	mu  *sync.RWMutex
	val Value
}

// Interface guard.
var _ fmt.Stringer = (*Schema)(nil)
