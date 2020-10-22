// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package zhash

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/sdslabs/kiwi"
)

type keyValuePair struct {
	key string
	value string
}

func TestZhash_Insert(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()
	toInsert := "random"

	// to insert the zhash into the store
	testInsertZhashHelper(store, testVals, t)

	v, err := store.Do(testKey, Insert, toInsert)
	if err != nil {
		t.Errorf("could not insert %q to zhash: %v", toInsert, err)
	}

	str, ok := v.(string)
	if !ok {
		t.Errorf("Insert did not return []string rather got %T", v)
	}

	if str != toInsert {
		t.Errorf("expected %q; got %q", toInsert, str[0])
	}

	// trying to insert with invalid parameter type
	_, err = store.Do(testKey, Insert, 123)
	if err == nil {
		t.Errorf("expected error; got nil while inserting with integer")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while inserting with integer; got %v", err)
	}

	// trying with invalid number of parameters
	_, err = store.Do(testKey, Insert)
	if err == nil {
		t.Errorf("expected error; got nil while inserting with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while inserting with 0 parameters; got %v", err)
	}
}

func TestZhash_Increment(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()

	// to insert the zhash into the store
	testInsertZhashHelper(store, testVals, t)

	v, err := store.Do(testKey, Increment, "a", 10)
	if err != nil {
		t.Errorf("could not increment %q to zhash: %v", "a", err)
	}

	sc, ok := v.(int)
	if !ok {
		t.Errorf("Increment did not return int rather got %T", v)
	}

	if sc != 10 {
		t.Errorf("expected %d; got %d", 10, sc)
	}

	// trying to increment with negative value
	v, err = store.Do(testKey, Increment, "a", -5)
	if err != nil {
		t.Errorf("could not increment %q to zhash: %v", "a", err)
	}

	sc, ok = v.(int)
	if !ok {
		t.Errorf("Increment did not return int rather got %T", v)
	}

	if sc != 5 {
		t.Errorf("expected %d; got %d", 5, sc)
	}

	// trying to increment with invalid parameter type
	_, err = store.Do(testKey, Increment, 123, 10)
	if err == nil {
		t.Errorf("expected error; got nil while incrementing with integer")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while incrementing with integer; got %v", err)
	}

	// trying with invalid number of parameters
	_, err = store.Do(testKey, Increment)
	if err == nil {
		t.Errorf("expected error; got nil while inserting with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while incrementing with 0 parameters; got %v", err)
	}

	// trying to increment invalid value (which is not present)
	_, err = store.Do(testKey, Increment, "z", 10)
	if err == nil {
		t.Errorf("expected error; got nil while incrementing with invalid value")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamValue {
		t.Errorf("expected ErrInvalidParamValue while incrementing with param which is not present; got %v", err)
	}
}

func TestZhash_Remove(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()
	toRemove := "a"

	// to insert the zhash into the store
	testInsertZhashHelper(store, testVals, t)

	v, err := store.Do(testKey, Remove, toRemove)
	if err != nil {
		t.Errorf("could not remove element(s) from the zhash: %v", err)
	}

	str, ok := v.([]string)
	if !ok {
		t.Errorf("expected []string from Remove; got %T", v)
	}

	if str[0] != toRemove {
		t.Errorf("element removed is not %q; got %q", toRemove, str[0])
	}

	// trying to remove invalid value (which is not present)
	_, err = store.Do(testKey, Remove, toRemove)
	if err == nil {
		t.Errorf("expected error; got nil while removing with invalid value")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamValue {
		t.Errorf("expected ErrInvalidParamValue while removing with param which is not present; got %v", err)
	}

	// trying to remove with invalid parameter type
	_, err = store.Do(testKey, Remove, true)
	if err == nil {
		t.Errorf("expected error; got nil while removing with param as bool")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while removing with param as bool; got %v", err)
	}

	// trying with invalid number of parameters
	_, err = store.Do(testKey, Remove)
	if err == nil {
		t.Errorf("expected error; got nil while removing with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while removing with 0 parameters; got %v", err)
	}
}

func TestZhash_Len(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()

	// to insert the zhash into the store
	testInsertZhashHelper(store, testVals, t)

	v, err := store.Do(testKey, Len)
	if err != nil {
		t.Errorf("could not get zhash's length from store: %v", err)
	}

	l, ok := v.(int)
	if !ok {
		t.Errorf("expected Len to return int; got %T", v)
	}

	if l != len(testVals) {
		t.Errorf("expected length: %d; got %d", len(testVals), l)
	}

	// try getting length with invalid number of params
	_, err = store.Do(testKey, Len, "noParamsShouldBePassed")
	if err == nil {
		t.Errorf("expected error; got nil while getting length with invalid number of params (1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while getting length with 1 parameters; got %v", err)
	}
}

func TestZhash_Get(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()
	toGet := "a"

	// to insert the zhash into the store
	testInsertZhashHelper(store, testVals, t)

	v, err := store.Do(testKey, Get, toGet)
	if err != nil {
		t.Errorf("could not get element's score from the zhash: %v", err)
	}

	sc, ok := v.(Item)
	if !ok {
		t.Errorf("expected int from Get; got %T", v)
	}

	if sc.Score != 0 {
		t.Errorf("expected score: %d; got: %d", 0, sc.Score)
	}

	// trying to get score of invalid value (which is not present)
	_, err = store.Do(testKey, Get, "x")
	if err == nil {
		t.Errorf("expected error; got nil while trying to get with invalid value")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamValue {
		t.Errorf("expected ErrInvalidParamValue while trying to get invalid element; got %v", err)
	}

	// trying to get with invalid parameter type
	_, err = store.Do(testKey, Get, true)
	if err == nil {
		t.Errorf("expected error; got nil while trying to get with param as bool")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while trying to get with param as bool; got %v", err)
	}

	// trying with invalid number of parameters
	_, err = store.Do(testKey, Get)
	if err == nil {
		t.Errorf("expected error; got nil while getting with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while getting with 0 parameters; got %v", err)
	}
}

func TestZhash_PeekMax(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()

	// to insert the zhash into the store
	testInsertZhashHelper(store, testVals, t)

	// Incrementing scores by some non-zero values
	_, err = store.Do(testKey, Increment, "a", 10)
	if err != nil {
		t.Errorf("could not increment %q to zhash: %v", "a", err)
	}

	_, err = store.Do(testKey, Increment, "c", -10)
	if err != nil {
		t.Errorf("could not increment %q to zhash: %v", "c", err)
	}

	v, err := store.Do(testKey, PeekMax)
	if err != nil {
		t.Errorf("could not peekmax from store: %v", err)
	}

	str, ok := v.(string)
	if !ok {
		t.Errorf("expected PeekMax to return string; got %T", v)
	}

	if str != "a" {
		t.Errorf("expected element: %q; got %q", "a", str)
	}

	// try peekmax with invalid number of params
	_, err = store.Do(testKey, PeekMax, "noParamsShouldBePassed")
	if err == nil {
		t.Errorf("expected error; got nil while using peekmin with invalid number of params (1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while using peekmin with 1 parameters; got %v", err)
	}
}

func TestZhash_PeekMin(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()

	// to insert the zhash into the store
	testInsertZhashHelper(store, testVals, t)

	// Incrementing scores by some non-zero values
	_, err = store.Do(testKey, Increment, "a", 10)
	if err != nil {
		t.Errorf("could not increment %q to zhash: %v", "a", err)
	}

	_, err = store.Do(testKey, Increment, "c", -10)
	if err != nil {
		t.Errorf("could not increment %q to zhash: %v", "c", err)
	}

	v, err := store.Do(testKey, PeekMin)
	if err != nil {
		t.Errorf("could not peekmin from store: %v", err)
	}

	str, ok := v.(string)
	if !ok {
		t.Errorf("expected PeekMin to return string; got %T", v)
	}

	if str != "c" {
		t.Errorf("expected element: %q; got %q", "c", str)
	}

	// try peekmin with invalid number of params
	_, err = store.Do(testKey, PeekMin, "noParamsShouldBePassed")
	if err == nil {
		t.Errorf("expected error; got nil while using peekmin with invalid number of params (1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while using peekmin with 1 parameters; got %v", err)
	}
}

func TestZHash_JSON(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestElems()
	testInsertZhashHelper(store, testVals, t)

	obj, err := store.ToJSON(testKey)
	if err != nil {
		t.Errorf("ToJSON returned unexpected error: %v", err)
	}

	// add a new key and initiate that key FromJSON and check if the value equals
	// by invoking the "PEEKMAX" and "PEEKMIN" actions.
	newKey := "xyz"
	err = store.AddKey(newKey, Type)
	if err != nil {
		t.Fatalf("cannot add new key to the store: %v", err)
	}

	err = store.FromJSON(newKey, obj)
	if err != nil {
		t.Errorf("FromJSON returned unexpected error: %v", err)
	}

	_, err = store.Do(testKey, Increment, testVals[0].key, 100)
	if err != nil {
		t.Errorf("cannot INCREMENT from the store: %v", err)
	}

	_, err = store.Do(testKey, Increment, testVals[len(testVals)-1].key, -100)
	if err != nil {
		t.Errorf("cannot INCREMENT from the store: %v", err)
	}

	v, err := store.Do(testKey, PeekMax)
	if err != nil {
		t.Errorf("cannot PEEKMAX from store: %v", err)
	}

	if testVals[0].key != v.(string) {
		t.Errorf("expected %q with max score; got %q", testVals[0].key, v.(string))
	}

	v, err = store.Do(testKey, PeekMin)
	if err != nil {
		t.Errorf("cannot PEEKMIN from store: %v", err)
	}

	if testVals[len(testVals)-1].key != v.(string) {
		t.Errorf("expected %q with max score; got %q", testVals[0].key, v.(string))
	}

	// NB: In zhashs we cannot test the full representation of JSON value,
	// since ordering of elements can change in the JSON representation
	// and the map representation (Go's maps are unordered).
	//
	// To test the representation in JSON, we can create a test value with
	// only one key, hence change in ordering won't affect the outcome.
	testK, testS, testV := "a", 102, "b"
	expectedJSON, err := json.Marshal(map[string]Item{testK: Item{testV, testS}})
	if err != nil {
		t.Fatalf("cannot marshal map to test: %v", err)
	}

	newKey = "def"
	err = store.AddKey(newKey, Type)
	if err != nil {
		t.Fatalf("cannot add new key to the store: %v", err)
	}

	_, err = store.Do(newKey, Insert, testK, testV)
	if err != nil {
		t.Errorf("cannot insert element into zhash: %v", err)
	}

	_, err = store.Do(newKey, Increment, testK, testS)
	if err != nil {
		t.Errorf("cannot increment score of element into zhash: %v", err)
	}

	obj, err = store.ToJSON(newKey)
	if err != nil {
		t.Errorf("ToJSON returned unexpected error: %v", err)
	}

	if !bytes.Equal(obj, expectedJSON) {
		t.Errorf("expected JSON:\n%s; got:\n%s", string(expectedJSON), string(obj))
	}
}

// testKey to test the value.
const testKey = "testZhash"

// newTestStore creates a new store for testing.
func newTestStore() (*kiwi.Store, error) {
	schema := kiwi.Schema{testKey: Type}
	return kiwi.NewStoreFromSchema(schema)
}

// newTestElems gets a new array of key value pairs that can be used for testing.
func newTestElems() []keyValuePair {
	return []keyValuePair{{"a", "b"}, {"c", "d"}}
}

// zhashToIFace converts []string to []interface{}.
func zhashToIFace(elems []string) []interface{} {
	v := make([]interface{}, len(elems))
	for i := range elems {
		v[i] = elems[i]
	}

	return v
}

// arrayEqual returns error if both the lists are not equal.
func arrayEqual(elems, expected []string) error {
	if len(elems) != len(expected) {
		return fmt.Errorf("expected length: %d; got length: %d", len(expected), len(elems))
	}

	for i := range expected {
		if elems[i] != expected[i] {
			return fmt.Errorf("expected index %d = %q; got %q", i, expected[i], elems[i])
		}
	}

	return nil
}

// testInsertZhashHelper tests if insertion works. This can later be reused with other tests.
func testInsertZhashHelper(store *kiwi.Store, testVals []keyValuePair, t *testing.T) {
	for i := range testVals {
		v, err := store.Do(testKey, Insert, testVals[i].key, testVals[i].value)
		if err != nil {
			t.Errorf("error doing: %v", err)
		}
		key, ok := v.(string)
		if !ok {
			t.Errorf("Insert action did not return a %T; rather got %T", key, v)
		}

		if key != testVals[i].key {
			t.Errorf("Insert did not return equal sets; expected %v; rather got %v", testVals[i].key, key)
		}
	}
}
