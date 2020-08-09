// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package stdkiwi

import "github.com/sdslabs/kiwi/values/zset"

// Zset implements methods for zset value type.
type Zset struct {
	store *Store
	key   string
}

// Guard guards the keys with values of str type.
func (z *Zset) Guard() {
	if err := z.GuardE(); err != nil {
		panic(err)
	}
}

// GuardE is same as Guard but does not panic, instead returns the error.
func (z *Zset) GuardE() error { return z.store.guardValueE(zset.Type, z.key) }

// Insert inserts the elements to the zset.
func (z *Zset) Insert(elements ...string) error {
	if len(elements) == 0 {
		return nil
	}

	ifaces := make([]interface{}, len(elements))
	for i := range elements {
		ifaces[i] = elements[i]
	}

	if _, err := z.store.Do(z.key, zset.Insert, ifaces...); err != nil {
		return err
	}

	return nil
}

// Remove removes the elements from the zset.
func (z *Zset) Remove(elements ...string) error {
	if len(elements) == 0 {
		return nil
	}

	ifaces := make([]interface{}, len(elements))
	for i := range elements {
		ifaces[i] = elements[i]
	}

	if _, err := z.store.Do(z.key, zset.Remove, ifaces...); err != nil {
		return err
	}

	return nil
}

// Increment increment the score of element of the zset.
func (z *Zset) Increment(element string, score int) error {
	if _, err := z.store.Do(z.key, zset.Increment, element, score); err != nil {
		return err
	}

	return nil
}

// Get gets the score of element from the zset.
func (z *Zset) Get(element string) (int, error) {
	v, err := z.store.Do(z.key, zset.Get, element)
	if err != nil {
		return -1, err
	}

	val, ok := v.(int)
	if !ok {
		return -1, newTypeErr(val, v)
	}

	return val, nil
}

// Len gets the length of the zset.
func (z *Zset) Len() (int, error) {
	v, err := z.store.Do(z.key, zset.Len)
	if err != nil {
		return 0, err
	}

	length, ok := v.(int)
	if !ok {
		return 0, newTypeErr(length, v)
	}

	return length, nil
}

// PeekMax gets the element with highest score from the zset.
func (z *Zset) PeekMax() (string, error) {
	v, err := z.store.Do(z.key, zset.PeekMax)
	if err != nil {
		return "", err
	}

	val, ok := v.(string)
	if !ok {
		return "", newTypeErr(val, v)
	}

	return val, nil
}

// PeekMin gets the element with minimum score from the zset.
func (z *Zset) PeekMin() (string, error) {
	v, err := z.store.Do(z.key, zset.PeekMin)
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
