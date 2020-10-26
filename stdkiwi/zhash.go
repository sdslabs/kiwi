// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import (
	"github.com/sdslabs/kiwi/values/zhash"
)

// Zhash implements methods for zhash value type.
type Zhash struct {
	store *Store
	key   string
}

// Guard guards the keys with values of str type.
func (z *Zhash) Guard() {
	if err := z.GuardE(); err != nil {
		panic(err)
	}
}

// GuardE is same as Guard but does not panic, instead returns the error.
func (z *Zhash) GuardE() error { return z.store.guardValueE(zhash.Type, z.key) }

// Insert inserts the elements to the zhash.
func (z *Zhash) Insert(key, value string) error {
	if _, err := z.store.Do(z.key, zhash.Insert, key, value); err != nil {
		return err
	}
	return nil
}

// Set sets the value of an pre-existing key of the zhash.
func (z *Zhash) Set(key, value string) error {
	if _, err := z.store.Do(z.key, zhash.Set, key, value); err != nil {
		return err
	}
	return nil
}

// Remove removes the elements from the zhash.
func (z *Zhash) Remove(elements ...string) error {
	if len(elements) == 0 {
		return nil
	}

	ifaces := make([]interface{}, len(elements))
	for i := range elements {
		ifaces[i] = elements[i]
	}

	if _, err := z.store.Do(z.key, zhash.Remove, ifaces...); err != nil {
		return err
	}

	return nil
}

// Increment increment the score of element of the zhash.
func (z *Zhash) Increment(element string, score int) error {
	if _, err := z.store.Do(z.key, zhash.Increment, element, score); err != nil {
		return err
	}

	return nil
}

// Get gets the value and the score of element from the zhash.
func (z *Zhash) Get(element string) (zhash.Item, error) {
	v, err := z.store.Do(z.key, zhash.Get, element)
	if err != nil {
		return zhash.Item{}, err
	}

	val, ok := v.(zhash.Item)
	if !ok {
		return zhash.Item{}, newTypeErr(val, v)
	}

	return val, nil
}

// Len gets the length of the zhash.
func (z *Zhash) Len() (int, error) {
	v, err := z.store.Do(z.key, zhash.Len)
	if err != nil {
		return 0, err
	}

	length, ok := v.(int)
	if !ok {
		return 0, newTypeErr(length, v)
	}

	return length, nil
}

// PeekMax gets the element with highest score from the zhash.
func (z *Zhash) PeekMax() (string, error) {
	v, err := z.store.Do(z.key, zhash.PeekMax)
	if err != nil {
		return "", err
	}

	val, ok := v.(string)
	if !ok {
		return "", newTypeErr(val, v)
	}

	return val, nil
}

// PeekMin gets the element with minimum score from the zhash.
func (z *Zhash) PeekMin() (string, error) {
	v, err := z.store.Do(z.key, zhash.PeekMin)
	if err != nil {
		return "", err
	}

	val, ok := v.(string)
	if !ok {
		return "", newTypeErr(val, v)
	}

	return val, nil
}

// Interface guard.
var _ Value = (*Set)(nil)
