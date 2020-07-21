// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package zset

import (
	"encoding/json"
	"fmt"

	"github.com/sdslabs/kiwi"

	"github.com/wangjia184/sortedset"
)

func init() {
	kiwi.RegisterValue(func() kiwi.Value {
		v := Value{}
		v.SortedSet = *(sortedset.New())
		return &v
	})
}

// Type of zset value.
const Type kiwi.ValueType = "zset"

// Value can store a set of elements having scores.
//
// It implements the kiwi.Value interface.
type Value struct{ sortedset.SortedSet }

// Various errors for zset value type.
var (
	ErrInvalidParamLen   = fmt.Errorf("not enough parameters")
	ErrInvalidParamType  = fmt.Errorf("invalid paramater type")
	ErrInvalidParamValue = fmt.Errorf("invalid parameter value")
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

// newParamValueErr creates an error where parameter value is wrong.
func newParamValueErr(v interface{}) error {
	typ := fmt.Sprintf("%T", v)
	return fmt.Errorf("%w: cannot process function on %q", ErrInvalidParamValue, typ)
}

const (
	// Insert inserts the element(s) in the zset.
	// Default score is 0.
	// If element is already present, updates the score to 0.
	//
	// Returns an array of added elements.
	Insert kiwi.Action = "INSERT"

	// Remove removes the element(s) from the zset.
	//
	// Returns an array of removed elements.
	Remove kiwi.Action = "REMOVE"

	// Increment increments the score by given value.
	//
	// Returns updated score.
	Increment kiwi.Action = "INCREMENT"

	// Len gets the length of the zset.
	//
	// Returns an integer.
	Len kiwi.Action = "LEN"

	// Get gets the score of given element.
	//
	// Returns an integer.
	Get kiwi.Action = "GET"

	// PeekMax gets the element with maximum score.
	//
	// Returns string value.
	// Returns nil if zset is empty.
	PeekMax kiwi.Action = "PEEKMAX"

	// PeekMin gets the element with minimum score.
	//
	// Returns string value.
	// Returns nil if zset is empty.
	PeekMin kiwi.Action = "PEEKMIN"
)

// Type returns v's type, i.e., "zset".
func (v *Value) Type() kiwi.ValueType {
	return Type
}

// DoMap returns the map of v's actions with it's do functions.
func (v *Value) DoMap() map[kiwi.Action]kiwi.DoFunc {
	return map[kiwi.Action]kiwi.DoFunc{
		Insert:    v.insert,
		Remove:    v.remove,
		Increment: v.increment,
		Len:       v.len,
		Get:       v.get,
		PeekMax:   v.peekmax,
		PeekMin:   v.peekmin,
	}
}

// ToJSON returns the raw byte array of v's data
func (v *Value) ToJSON() (json.RawMessage, error) {
	nodes := v.GetByRankRange(1, -1, false)

	vals := map[string]int{}
	for _, node := range nodes {
		vals[node.Key()] = int(node.Score())
	}

	c, err := json.Marshal(vals)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(c), nil
}

// FromJSON populates the s with the data from RawMessage
func (v *Value) FromJSON(rawmessage json.RawMessage) error {
	vals := map[string]int{}
	if err := json.Unmarshal(rawmessage, &vals); err != nil {
		return err
	}

	for key, score := range vals {
		v.AddOrUpdate(key, sortedset.SCORE(score), nil)
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
		(*v).AddOrUpdate(toInsert, 0, nil)
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

		temp := (*v).Remove(toRemove)
		if temp == nil {
			return nil, newParamValueErr(toRemove)
		}
		out[i] = toRemove
	}
	return out, nil
}

// increment implements the INCREMENT action.
func (v *Value) increment(params ...interface{}) (interface{}, error) {
	if len(params) != 2 {
		return nil, newParamLenErr(len(params), 2)
	}

	key, ok := params[0].(string)
	if !ok {
		return nil, newParamTypeErr(params[0], key)
	}
	sc, ok := params[1].(int)
	if !ok {
		return nil, newParamTypeErr(params[1], sc)
	}

	temp := (*v).GetByKey(key)
	if temp == nil {
		return nil, newParamValueErr(key)
	}

	(*v).AddOrUpdate(key, sortedset.SCORE(sc)+temp.Score(), nil)
	return sc + int(temp.Score()), nil
}

// len implements the LEN action.
func (v *Value) len(params ...interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, newParamLenErr(len(params), 0)
	}

	return (*v).GetCount(), nil
}

// get implements the GET action.
func (v *Value) get(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, newParamLenErr(len(params), 1)
	}

	toCheck, ok := params[0].(string)
	if !ok {
		return nil, newParamTypeErr(params[0], toCheck)
	}

	temp := (*v).GetByKey(toCheck)
	if temp == nil {
		return nil, newParamValueErr(toCheck)
	}
	return int(temp.Score()), nil
}

// peekmax implements the PEEKMAX action.
func (v *Value) peekmax(params ...interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, newParamLenErr(len(params), 0)
	}

	temp := (*v).PeekMax()
	if temp == nil {
		return nil, nil
	}
	return temp.Key(), nil
}

// peekmin implements the PEEKMIN action.
func (v *Value) peekmin(params ...interface{}) (interface{}, error) {
	if len(params) != 0 {
		return nil, newParamLenErr(len(params), 0)
	}

	temp := (*v).PeekMin()
	if temp == nil {
		return nil, nil
	}
	return temp.Key(), nil
}

// Interface guard.
var _ kiwi.Value = (*Value)(nil)
