// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package str

import (
	"encoding/json"
	"fmt"

	"github.com/sdslabs/kiwi"
)

func init() {
	kiwi.RegisterValue(func() kiwi.Value { return new(Value) })
}

// Type of str value.
const Type kiwi.ValueType = "str"

// Value can store a string.
//
// It implements the kiwi.Value interface.
type Value string

const (
	// Get gets the string value.
	//
	// Returns a string.
	Get kiwi.Action = "GET"

	// Update updates the value of the string.
	//
	// Returns the updated string.
	Update kiwi.Action = "UPDATE"
)

// Type returns v's type, i.e., "str".
func (v *Value) Type() kiwi.ValueType {
	return Type
}

// DoMap returns the map of v's actions with it's do functions.
func (v *Value) DoMap() map[kiwi.Action]kiwi.DoFunc {
	return map[kiwi.Action]kiwi.DoFunc{
		Get: func(params ...interface{}) (interface{}, error) {
			// don't need any params
			return string(*v), nil
		},
		Update: func(params ...interface{}) (interface{}, error) {
			// requires one string parameter
			if len(params) < 1 {
				return nil, fmt.Errorf("str.Value.Update requires 1 argument")
			}

			newStr, ok := params[0].(string)
			if !ok {
				return nil, fmt.Errorf("str.Value.Update takes only string argument")
			}

			*v = Value(newStr)
			return newStr, nil
		},
	}
}

// ToJSON returns the raw byte array of v's data
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

// Interface guard.
var _ kiwi.Value = (*Value)(nil)
