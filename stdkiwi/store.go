// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"fmt"

	"github.com/sdslabs/kiwi"
)

// Store wraps the kiwi.Store and implements various API methods for standard values.
type Store struct{ *kiwi.Store }

// NewStore creates a new std store.
func NewStore() *Store {
	return &Store{kiwi.NewStore()}
}

// NewStoreFromSchema creates a new std store from schema.
func NewStoreFromSchema(schema kiwi.Schema) (*Store, error) {
	store, err := kiwi.NewStoreFromSchema(schema)
	if err != nil {
		return nil, err
	}

	return &Store{store}, nil
}

//
// Other methods of kiwi.Store are directly accessible
//

// guardValueE returns error if the key does not correspond to the value type.
//
// Functions wrapping guardValue should be executed in the init of the program to
// safely access the methods corresponding to the key value without chance of handling
// errors during the runtime of the program.
// This also throws error if the key does not exist; Quite helpful when schema is defined.
//
// Use "Guard" for when schema defined so runtime errors are avoided else use "GuardE".
func (s *Store) guardValueE(val kiwi.ValueType, key string) error {
	typ, err := s.GetValueType(key)
	if err != nil {
		return err
	}

	if val != typ {
		return fmt.Errorf("expected value type for key(%q) is %q; got %q", key, val, typ)
	}

	return nil
}

// Str returns an "Str" with the key set as "key".
//
// This does not verify if the key and value type pair is correct.
// If this does not work, "Do" will eventually throw an error.
// To avoid this use "GuardE" method or "Guard" if schema is pre-defined.
func (s *Store) Str(key string) *Str {
	return &Str{
		store: s,
		key:   key,
	}
}

// List returns a "List" with the key set as "key".
//
// This does not verify if the key and value type pair is correct.
// If this does not work, "Do" will eventually throw an error.
// To avoid this use "GuardE" method or "Guard" if schema is pre-defined.
func (s *Store) List(key string) *List {
	return &List{
		store: s,
		key:   key,
	}
}

// Set returns a "Set" with the key set as "key".
//
// This does not verify if the key and value type pair is correct.
// If this does not work, "Do" will eventually throw an error.
// To avoid this use "GuardE" method or "Guard" if schema is pre-defined.
func (s *Store) Set(key string) *Set {
	return &Set{
		store: s,
		key:   key,
	}
}

// Hash returns a "Hashmap" with the key set as "key".
//
// This does not verify if the key and value type pair is correct.
// If this does not work, "Do" will eventually throw an error.
// To avoid this use "GuardE" method or "Guard" if schema is pre-defined.
func (s *Store) Hash(key string) *Hash {
	return &Hash{
		store: s,
		key:   key,
	}
}

// Zset returns a "Zset" with the key set as "key".
//
// This does not verify if the key and value type pair is correct.
// If this does not work, "Do" will eventually throw an error.
// To avoid this use "GuardE" method or "Guard" if schema is pre-defined.
func (s *Store) Zset(key string) *Zset {
	return &Zset{
		store: s,
		key:   key,
	}
}

// Value can be used to access the methods for standard value types.
type Value interface {
	// Guard should panic if the key does not correspond to the correct type.
	Guard()

	// GuardE is same as guard but does not panic, instead throws an error.
	GuardE() error
}

//
// Helper functions
//

// newTypeErr creates a new error for when there's unexpectly types unmatched for returns.
func newTypeErr(expected, actual interface{}) error {
	return fmt.Errorf("unexpected error: return value expected %T; got %T", expected, actual)
}
