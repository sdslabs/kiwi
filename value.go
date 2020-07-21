// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package kiwi

import (
	"encoding/json"
	"fmt"
)

// Errors related to values.
var (
	ErrValueRegistered    = fmt.Errorf("value already registered")
	ErrValueNotRegistered = fmt.Errorf("value not registered")
)

// newValueErr creates a new error with related value type.
func newValueErr(err error, typ ValueType) error {
	return fmt.Errorf("%w: %v", err, typ)
}

// Errors related to actions.
var (
	ErrInvalidAction = fmt.Errorf("invalid action")
)

// valueMap maintains all the values registered with the library.
var valueMap = map[ValueType]func() Value{}

type (
	// ValueType is the name of type of value.
	ValueType string

	// Action is required to invoke (or "Do") an action.
	Action string

	// DoFunc is a function that can be executed when an action is invoked.
	DoFunc func(params ...interface{}) (interface{}, error)
)

// Value is something that can be associated with a "key".
//
// A value implements its own methods and can be accessed by type assertion.
// To add a value type to the register simply call RegisterValue.
// All the default values are already registered with the package.
type Value interface {
	// Type returns the type of the value.
	Type() ValueType

	// DoMap returns a map which associates actions with a do function.
	DoMap() map[Action]DoFunc

	// ToJSON returns a raw byte array of the data.
	ToJSON() (json.RawMessage, error)

	// FromJSON returns the data in golang from a raw byte array.
	FromJSON(json.RawMessage) error
}

// RegisterValue registers a new value type with the package.
//
// It takes in two params: the type of the value and a function to create a new value.
func RegisterValue(newFn func() Value) {
	typ := newFn().Type()

	if _, ok := valueMap[typ]; ok {
		panic(newValueErr(ErrValueRegistered, typ))
	}

	valueMap[typ] = newFn
}

// ListRegisteredValues lists all the values registered with the package.
func ListRegisteredValues() []string {
	vals := make([]string, len(valueMap))

	i := 0
	for key := range valueMap {
		vals[i] = string(key)
		i++
	}

	return vals
}

// newValue creates a value from it's type.
func newValue(typ ValueType) (Value, error) {
	nf, ok := valueMap[typ]
	if !ok {
		return nil, newValueErr(ErrValueNotRegistered, typ)
	}

	return nf(), nil
}
