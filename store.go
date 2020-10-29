// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package kiwi

import (
	"encoding/json"
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
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.keyExists(key) == nil
}

// AddKey adds a new key to the store. It throws an error if the key already exists.
func (s *Store) AddKey(key string, typ ValueType) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.keyNotExist(key); err != nil {
		return err
	}

	v, err := newValue(typ)
	if err != nil {
		return err
	}

	s.kv[key] = valWrapper{val: v, mu: &sync.RWMutex{}}
	return nil
}

// UpdateKey updates the key if it exists. Throws an error if it doesn't.
func (s *Store) UpdateKey(key string, typ ValueType) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.keyExists(key); err != nil {
		return err
	}

	v, err := newValue(typ)
	if err != nil {
		return err
	}

	old := s.kv[key]
	old.mu.Lock()
	defer old.mu.Unlock()

	s.kv[key] = valWrapper{val: v, mu: &sync.RWMutex{}}
	return nil
}

// DeleteKey deletes the key if it exists. Throws an error if it doesn't.
func (s *Store) DeleteKey(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.keyExists(key); err != nil {
		return err
	}

	old := s.kv[key]
	old.mu.Lock()
	defer old.mu.Unlock()

	delete(s.kv, key)
	return nil
}

// GetValueType returns the type of value corresponding to the key.
func (s *Store) GetValueType(key string) (ValueType, error) {
	s.mu.RLock()
	if err := s.keyExists(key); err != nil {
		s.mu.RUnlock()
		return "", err
	}
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

	return s.getSchema()
}

// getSchema returns the schema.
func (s *Store) getSchema() Schema {
	schema := make(Schema)
	for k := range s.kv {
		schema[k] = s.kv[k].val.Type()
	}

	return schema
}

// Do executes the action for the value associated with the value.
func (s *Store) Do(key string, action Action, params ...interface{}) (interface{}, error) {
	s.mu.RLock()
	if err := s.keyExists(key); err != nil {
		s.mu.RUnlock()
		return nil, err
	}
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

// ToJSON converts the data associated with the value into JSON format.
func (s *Store) ToJSON(key string) (json.RawMessage, error) {
	s.mu.RLock()
	if err := s.keyExists(key); err != nil {
		s.mu.RUnlock()
		return nil, err
	}
	v := s.kv[key]
	s.mu.RUnlock()

	v.mu.RLock()
	defer v.mu.RUnlock()

	return s.toJSON(&v)
}

// toJSON converts the value's data to JSON.
func (s *Store) toJSON(v *valWrapper) (json.RawMessage, error) {
	res, err := v.val.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("error in ToJSON: %v", err)
	}

	return res, nil
}

// FromJSON takes the raw JSON form of data and loads it into the value.
func (s *Store) FromJSON(key string, rawmessage json.RawMessage) error {
	s.mu.RLock()
	if err := s.keyExists(key); err != nil {
		s.mu.RUnlock()
		return err
	}
	v := s.kv[key]
	s.mu.RUnlock()

	v.mu.Lock()
	defer v.mu.Unlock()

	return s.fromJSON(&v, rawmessage)
}

// fromJSON converts the raw JSON to the value.
func (s *Store) fromJSON(v *valWrapper, rawmessage json.RawMessage) error {
	if err := v.val.FromJSON(rawmessage); err != nil {
		return fmt.Errorf("error in FromJSON: %v", err)
	}

	return nil
}

// ValJSON is the JSON object for each value.
type ValJSON struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// StoreJSON is the type which includes all the data for the store.
type StoreJSON = map[string]ValJSON

// Export returns JSON data for the store.
//
// The data is in the format (StoreJSON):
//
// 	{
// 		"key_1": {
// 			"type": "str",
// 			"data": "hello"
// 		},
// 		"key_2": {
// 			"type": "hash",
// 			"data": {
// 				"a": "b",
// 				"c": "d"
// 			}
// 		}
// 	}
//
func (s *Store) Export() (json.RawMessage, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	schema := s.getSchema()
	sjson := make(StoreJSON, len(schema))

	for k, v := range schema {
		vw := s.kv[k]

		vw.mu.RLock()
		data, err := s.toJSON(&vw)
		vw.mu.Unlock()
		if err != nil {
			return nil, fmt.Errorf("error exporting for %q key: %v", k, err)
		}
		sjson[k] = ValJSON{
			Type: string(v),
			Data: data,
		}
	}

	c, err := json.Marshal(sjson)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// ImportOpts are the options that can be used to configure how to import
// data into the store from raw JSON.
type ImportOpts struct {
	// AddKeys specifies whether to add keys that do not exist.
	AddKeys bool

	// UpdateTypes specifies if the type of key in JSON doesn't match the one
	// with already-defined key, should the type of key be updated or not.
	UpdateTypes bool

	// ErrOnInvalidKey specifies whether to throw error if the key in the JSON
	// does not exist in the actual schema of the Store.
	// This option is considered only when `AddKeys` is false.
	ErrOnInvalidKey bool
}

// Import loads store from the data.
//
// The default behavior is that the store takes the data from the JSON and
// if an unknown key exists, i.e., a key that is not already added to the
// store, it silently skips the value associated with it. This can be
// configured using the ImportOpts.
//
// The data is in the format (StoreJSON):
//
// 	{
// 		"key_1": {
// 			"type": "str",
// 			"data": "hello"
// 		},
// 		"key_2": {
// 			"type": "hash",
// 			"data": {
// 				"a": "b",
// 				"c": "d"
// 			}
// 		}
// 	}
//
func (s *Store) Import(rawmessage json.RawMessage, opts ImportOpts) error {
	var sjson StoreJSON
	if err := json.Unmarshal(rawmessage, &sjson); err != nil {
		return err
	}

	// Lock the store for whole of the process now.
	s.mu.Lock()
	defer s.mu.Unlock()

	for k := range sjson {
		if err := s.keyExists(k); err != nil {
			if !opts.AddKeys && !opts.ErrOnInvalidKey {
				continue
			}
			if !opts.AddKeys && opts.ErrOnInvalidKey {
				return err
			}
			if opts.AddKeys {
				if er := s.AddKey(k, ValueType(sjson[k].Type)); er != nil {
					return er
				}
			}
		}

		// now that the key exists
		v := s.kv[k]

		if err := func() error {
			v.mu.Lock()
			defer v.mu.Unlock()

			if sjson[k].Type != string(v.val.Type()) {
				if !opts.UpdateTypes {
					return fmt.Errorf("value type in JSON and store schema do not match")
				}
				if er := s.UpdateKey(k, ValueType(sjson[k].Type)); er != nil {
					return er
				}
			}

			if err := s.FromJSON(k, sjson[k].Data); err != nil {
				return err
			}

			return nil
		}(); err != nil {
			return err
		}
	}

	return nil
}

// keyExists checks if the key exists in the store. Throws an error if it doesn't.
func (s *Store) keyExists(key string) error {
	if _, ok := s.kv[key]; !ok {
		return newKeyErr(ErrKeyNotExist, key)
	}

	return nil
}

// keyExists checks if the key does not exist in the store. Throws an error if it does.
func (s *Store) keyNotExist(key string) error {
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
