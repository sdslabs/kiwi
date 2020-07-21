// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package hash

import (
	"encoding/json"
	"fmt"

	"github.com/sdslabs/kiwi"
)

func init() {
	kiwi.RegisterValue(func() kiwi.Value { return &Value{} })
}

// Type of set value.
const Type kiwi.ValueType = "hash"

// Value can store a string-string hash map.
//
// It implements the kiwi.Value interface.
type Value map[string]string

// Various errors for hash value type.
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
	// Insert inserts the key-value pair in hash map.
	// Updates the value if key is already present.
	//
	// Returns the added key string
	Insert kiwi.Action = "INSERT"

	// Remove removes the string(s) from the set.
	//
	// Returns an array of removed key strings.
	Remove kiwi.Action = "REMOVE"

	// Has checks if hash-map has the key.
	//
	// Returns boolean value.
	Has kiwi.Action = "HAS"

	// Len gets the length of the hash map.
	//
	// Returns an integer.
	Len kiwi.Action = "LEN"

	// Get gets the value of the given key(s).
	//
	// Returns an array of values.
	Get kiwi.Action = "GET"

	// Keys gets all the keys from the hash-map.
	//
	// Returns an array of keys.
	Keys kiwi.Action = "KEYS"

	// Map gets a copy of hash-map.
	//
	// Returns a hash-map.
	Map kiwi.Action = "MAP"
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
		Keys:   v.keys,
		Map:    v.copymap,
	}
}

// ToJSON returns the raw byte array of s's data
func (v *Value) ToJSON() (json.RawMessage, error) {
	c, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(c), nil
}

// FromJSON populates the s with the data from RawMessage
func (v *Value) FromJSON(rawmessage json.RawMessage) error {
	return json.Unmarshal(rawmessage, &v)
}

// insert implements the INSERT action.
func (v *Value) insert(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, newParamLenErr(len(params), 2)
	}

	var (
		ok    bool
		key   string
		value string
	)

	key, ok = params[0].(string)
	if !ok {
		return nil, newParamTypeErr(params[0], key)
	}

	value, ok = params[1].(string)
	if !ok {
		return nil, newParamTypeErr(params[1], value)
	}

	(*v)[key] = value
	return key, nil
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
	if len(params) == 0 {
		return nil, newParamLenErr(0, 1)
	}

	var (
		toGet string
		ok    bool
	)

	out := make([]string, len(params))
	for i := 0; i < len(params); i++ {
		toGet, ok = params[i].(string)
		if !ok {
			return nil, newParamTypeErr(params[i], toGet)
		}
		out[i] = (*v)[toGet]
	}
	return out, nil
}

// key implements the KEY action.
func (v *Value) keys(params ...interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, newParamLenErr(len(params), 0)
	}

	i := 0
	out := make([]string, len(*v))
	for key := range *v {
		out[i] = key
		i++
	}
	return out, nil
}

// copymap implements the MAP action.
func (v *Value) copymap(params ...interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, newParamLenErr(len(params), 0)
	}

	out := make(map[string]string)
	for key, value := range *v {
		out[key] = value
	}
	return out, nil
}

// Interface guard.
var _ kiwi.Value = (*Value)(nil)
