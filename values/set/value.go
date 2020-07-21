// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package set

import (
	"encoding/json"
	"fmt"

	"github.com/sdslabs/kiwi"
)

func init() {
	kiwi.RegisterValue(func() kiwi.Value { return &Value{} })
}

// Type of set value.
const Type kiwi.ValueType = "set"

// Value can store a set of strings.
//
// It implements the kiwi.Value interface.
type Value map[string]struct{}

// Various errors for set value type.
var (
	ErrInvalidParamLen  = fmt.Errorf("not enough parameters")
	ErrInvalidParamType = fmt.Errorf("invalid paramater type")
)

// newParamLenErr creates an error where parameter length is wrong.
func newParamLenErr(u, l int) error {
	return fmt.Errorf("%w: got %d; requires %d", ErrInvalidParamLen, u, l)
}

// newParamTypeErr creates an error where parameter type is wrong.
func newParamTypeErr(p, e interface{}) error {
	typ := fmt.Sprintf("%T", e)
	return fmt.Errorf("%w: %#v not a(n) %q", ErrInvalidParamType, p, typ)
}

const (
	// Insert inserts the element(s) in the set.
	//
	// Returns an array of added elements.
	Insert kiwi.Action = "INSERT"

	// Remove removes the element(s) from the set.
	//
	// Returns an array of removed elements.
	Remove kiwi.Action = "REMOVE"

	// Has checks if set has the element.
	//
	// Returns boolean value.
	Has kiwi.Action = "HAS"

	// Len gets the length of the set.
	//
	// Returns an integer.
	Len kiwi.Action = "LEN"

	// Get gets all the elements of set.
	//
	// Returns an array of elements.
	Get kiwi.Action = "GET"
)

// Type returns v's type, i.e., "set".
func (v *Value) Type() kiwi.ValueType {
	return Type
}

// DoMap returns the map of v's actions with it's do functions.
func (v *Value) DoMap() map[kiwi.Action]kiwi.DoFunc {
	return map[kiwi.Action]kiwi.DoFunc{
		Insert: v.insert,
		Remove: v.remove,
		Has:    v.has,
		Len:    v.len,
		Get:    v.get,
	}
}

// ToJSON returns the raw byte array of value
func (v *Value) ToJSON() (json.RawMessage, error) {
	vals := make([]string, len(*v))

	i := 0
	for k := range *v {
		vals[i] = k
		i++
	}

	c, err := json.Marshal(vals)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(c), nil
}

// FromJSON populates the value with the data from RawMessage
func (v *Value) FromJSON(rawmessage json.RawMessage) error {
	vals := []string{}
	if err := json.Unmarshal(rawmessage, &vals); err != nil {
		return err
	}

	for _, c := range vals {
		(*v)[c] = struct{}{}
	}

	return nil
}

// insert implements the INSERT action.
func (v *Value) insert(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		return nil, newParamLenErr(0, 1)
	}

	var (
		toInsert string
		ok       bool
	)

	out := make([]string, len(params))
	for i := 0; i < len(params); i++ {
		toInsert, ok = params[i].(string)
		if !ok {
			return nil, newParamTypeErr(params[i], toInsert)
		}
		(*v)[toInsert] = struct{}{}
		out[i] = toInsert
	}
	return out, nil
}

// remove implements the REMOVE action.
func (v *Value) remove(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		return nil, newParamLenErr(0, 1)
	}

	var (
		toRemove string
		ok       bool
	)

	out := make([]string, len(params))
	for i := 0; i < len(params); i++ {
		toRemove, ok = params[i].(string)
		if !ok {
			return nil, newParamTypeErr(params[i], toRemove)
		}
		delete((*v), toRemove)
		out[i] = toRemove
	}
	return out, nil
}

// has implements the HAS action.
func (v *Value) has(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, newParamLenErr(len(params), 1)
	}

	toCheck, ok := params[0].(string)
	if !ok {
		return nil, newParamTypeErr(params[0], toCheck)
	}

	_, ok = (*v)[toCheck]
	return ok, nil
}

// len implements the LEN action.
func (v *Value) len(params ...interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, newParamLenErr(len(params), 0)
	}

	return len(*v), nil
}

// get implements the GET action.
func (v *Value) get(params ...interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, newParamLenErr(len(params), 0)
	}

	i := 0
	out := make([]string, len(*v))
	for elem := range *v {
		out[i] = elem
		i++
	}
	return out, nil
}

// Interface guard.
var _ kiwi.Value = (*Value)(nil)
