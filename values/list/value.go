// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package list

import (
	"encoding/json"
	"fmt"

	"github.com/sdslabs/kiwi"
)

func init() {
	kiwi.RegisterValue(func() kiwi.Value { return new(Value) })
}

// Type of list value.
const Type kiwi.ValueType = "list"

// Value can store an array of strings.
//
// It implements the kiwi.Value interface.
type Value []string

// Various errors for list value type.
var (
	ErrInvalidIndex     = fmt.Errorf("cannot access invalid index")
	ErrInvalidParamLen  = fmt.Errorf("not enough parameters")
	ErrInvalidParamType = fmt.Errorf("invalid paramater type")
)

// newIndexErr creates an error where an out of bounds index is accessed.
func newIndexErr(i, l int) error {
	return fmt.Errorf("%w: %d in slice of length=%d", ErrInvalidIndex, i, l)
}

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
	// Get gets the string at the particular index.
	// If no index is provided it gets the last element.
	//
	// Returns a string.
	Get kiwi.Action = "GET"

	// Set sets the value of string at the index.
	// If no index is given and just the string, it updates the last element.
	//
	// Returns the sett-ed string.
	Set kiwi.Action = "SET"

	// Slice gets slice of the array from begin to end ( excluded ).
	//
	// Returns a slice of strings.
	Slice kiwi.Action = "SLICE"

	// Len gets the length of the list.
	//
	// Returns an integer.
	Len kiwi.Action = "LEN"

	// Append adds elements to the end of list.
	//
	// Returns a slice of strings which are appended.
	Append kiwi.Action = "APPEND"

	// Pop removes the last "n" elements from the string.
	// If "n" is not provided, pops last element.
	//
	// Returns a slice of removed strings.
	Pop kiwi.Action = "POP"

	// Remove removes the elements specified from the list.
	// These can be indexes or the elements themselves.
	//
	// Returns the removed string.
	Remove kiwi.Action = "REMOVE"

	// Find gets the index of element if it exists else -1.
	//
	// Returns an integer.
	Find kiwi.Action = "FIND"
)

// Type returns v's type, i.e., "list".
func (v *Value) Type() kiwi.ValueType {
	return Type
}

// DoMap returns the map of v's actions with it's do functions.
func (v *Value) DoMap() map[kiwi.Action]kiwi.DoFunc {
	return map[kiwi.Action]kiwi.DoFunc{
		Get:    v.get,
		Set:    v.set,
		Slice:  v.slice,
		Len:    v.length,
		Append: v.pushBack,
		Pop:    v.pop,
		Remove: v.remove,
		Find:   v.find,
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
	return json.Unmarshal(rawmessage, v)
}

// get implements the GET action.
func (v *Value) get(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		return (*v)[len(*v)-1], nil
	}

	idx, ok := params[0].(int)
	if !ok {
		return nil, newParamTypeErr(params[0], idx)
	}

	if idx < 0 || idx >= len(*v) {
		return nil, newIndexErr(idx, len(*v))
	}

	return (*v)[idx], nil
}

// set implements the SET action.
func (v *Value) set(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		return nil, newParamLenErr(0, 1)
	}

	var (
		idx      = len(*v) - 1
		toUpdate string
		ok       bool
	)

	if len(params) == 1 {
		toUpdate, ok = params[0].(string)
		if !ok {
			return nil, newParamTypeErr(params[0], toUpdate)
		}
	} else {
		idx, ok = params[0].(int)
		if !ok {
			return nil, newParamTypeErr(params[0], idx)
		}

		toUpdate, ok = params[1].(string)
		if !ok {
			return nil, newParamTypeErr(params[1], toUpdate)
		}
	}

	if idx < 0 || idx >= len(*v) {
		return nil, newIndexErr(idx, len(*v))
	}

	(*v)[idx] = toUpdate
	return toUpdate, nil
}

// slice implements the SLICE action.
func (v *Value) slice(params ...interface{}) (interface{}, error) {
	start, end := 0, len(*v)
	var ok bool

	switch len(params) {
	case 0:
	case 1:
		end, ok = params[0].(int)
		if !ok {
			return nil, newParamTypeErr(params[0], end)
		}

		if end < 0 || end > len(*v) {
			return nil, newIndexErr(end, len(*v))
		}

	default: // > 1
		start, ok = params[0].(int)
		if !ok {
			return nil, newParamTypeErr(params[0], start)
		}

		end, ok = params[1].(int)
		if !ok {
			return nil, newParamTypeErr(params[1], end)
		}

		if start > end || start < 0 || end > len(*v) {
			return nil, fmt.Errorf("%w: (%d, %d) in slice of length %d", ErrInvalidIndex, start, end, len(*v))
		}
	}

	newSlice := make([]string, end-start)
	for i := start; i < end; i++ {
		newSlice[i-start] = (*v)[i]
	}

	return newSlice, nil
}

// length implements the LEN action.
func (v *Value) length(params ...interface{}) (interface{}, error) {
	return len(*v), nil
}

// pushBack implements the APPEND action.
func (v *Value) pushBack(params ...interface{}) (interface{}, error) {
	toAppend := make([]string, len(params))

	for i := range params {
		str, ok := params[i].(string)
		if !ok {
			return nil, newParamTypeErr(params[i], str)
		}

		toAppend[i] = str
	}

	*v = append(*v, toAppend...)
	return toAppend, nil
}

// pop implements the POP action.
func (v *Value) pop(params ...interface{}) (interface{}, error) {
	n := 1

	if len(params) > 0 {
		var ok bool
		n, ok = params[0].(int)
		if !ok {
			return nil, newParamTypeErr(params[0], n)
		}
	}

	if n < 0 || n > len(*v) {
		return nil, newIndexErr(n, len(*v))
	}

	removed := make([]string, n)
	for i := 0; i < n; i++ {
		removed[i] = (*v)[len(*v)-n+i]
	}

	*v = (*v)[:len(*v)-n]
	return removed, nil
}

// remove implements the REMOVE action.
func (v *Value) remove(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		return nil, newParamLenErr(0, 1)
	}

	var toRemove string

	str, ok := params[0].(string)
	if ok {
		toRemove = str

		for i := 0; i < len(*v); i++ {
			if toRemove != (*v)[i] {
				continue
			}

			*v = append((*v)[:i], (*v)[i+1:]...)
			break
		}
	} else {
		idx, ok := params[0].(int)
		if !ok {
			return nil, newParamTypeErr(params[0], idx)
		}

		if idx < 0 || idx >= len(*v) {
			return nil, newIndexErr(idx, len(*v))
		}

		toRemove = (*v)[idx]

		*v = append((*v)[:idx], (*v)[idx+1:]...)
	}

	return toRemove, nil
}

// find implements the FIND action.
func (v *Value) find(params ...interface{}) (interface{}, error) {
	if len(params) == 0 {
		return nil, newParamLenErr(0, 1)
	}

	str, ok := params[0].(string)
	if !ok {
		return nil, newParamTypeErr(params[0], str)
	}

	for i := 0; i < len(*v); i++ {
		if (*v)[i] == str {
			return i, nil
		}
	}

	return -1, nil
}

// Interface guard.
var _ kiwi.Value = (*Value)(nil)
