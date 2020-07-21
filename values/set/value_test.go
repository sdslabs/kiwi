// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package set

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/sdslabs/kiwi"
)

func TestSet_Insert(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestSet()
	toInsert := "random"

	// to insert the set into the store
	testInsertSetHelper(store, testVals, t)

	v, err := store.Do(testKey, Insert, toInsert)
	if err != nil {
		t.Errorf("could not insert %q to set: %v", toInsert, err)
	}

	str, ok := v.([]string)
	if !ok {
		t.Errorf("Insert did not return []string rather got %T", v)
	}

	if str[0] != toInsert {
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

func TestSet_Remove(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestSet()
	toRemove := "c"

	// to insert the set into the store
	testInsertSetHelper(store, testVals, t)

	v, err := store.Do(testKey, Remove, toRemove)
	if err != nil {
		t.Errorf("could not remove elements from the list: %v", err)
	}

	str, ok := v.([]string)
	if !ok {
		t.Errorf("expected []string from Remove; got %T", v)
	}

	if str[0] != toRemove {
		t.Errorf("element removed is not %q; got %q", toRemove, str[0])
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

func TestSet_Has(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestSet()
	toFind := "b"

	// to insert the set into the store
	testInsertSetHelper(store, testVals, t)

	v, err := store.Do(testKey, Has, toFind)
	if err != nil {
		t.Errorf("cannot find %q in the set: %v", toFind, err)
	}

	fidx, ok := v.(bool)
	if !ok {
		t.Errorf("expected boolean; got %T while finding", v)
	}

	if fidx == false {
		t.Errorf("testVals has %q but returned false", toFind)
	}

	// finding something that does not exist
	v, err = store.Do(testKey, Has, "randomStringThatShouldNotExist")
	if err != nil {
		t.Errorf("cannot find invalid value in the set: %v", err)
	}

	fidx, ok = v.(bool)
	if !ok {
		t.Errorf("expected boolean; got %T while finding", v)
	}

	if fidx == true {
		t.Errorf("finding invalid value; expected false")
	}

	// try finding with invalid number of params
	_, err = store.Do(testKey, Has)
	if err == nil {
		t.Errorf("expected error; got nil while finding with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while finding with 0 parameters; got %v", err)
	}

	// trying to find with invalid parameter type
	_, err = store.Do(testKey, Has, 123)
	if err == nil {
		t.Errorf("expected error; got nil while finding with string as int")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while finding with string as int; got %v", err)
	}
}

func TestSet_Len(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestSet()

	// to insert the set into the store
	testInsertSetHelper(store, testVals, t)

	v, err := store.Do(testKey, Len)
	if err != nil {
		t.Errorf("could not get set's length from store: %v", err)
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
		t.Errorf("expected error; got nil while finding with invalid number of params (1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while finding with 1 parameters; got %v", err)
	}
}

func TestSet_Get(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestSet()

	// to insert the set into the store
	testInsertSetHelper(store, testVals, t)

	v, err := store.Do(testKey, Get)
	if err != nil {
		t.Errorf("could not get set's elements from store: %v", err)
	}

	set, ok := v.([]string)
	if !ok {
		t.Errorf("expected Get to return []string; got %T", v)
	}

	if len(set) != len(testVals) {
		t.Errorf("expected length of elems: %d; got %d", len(testVals), len(set))
	}

	for i := 0; i < len(set); i++ {
		_, ok = testVals[set[i]]
		if ok == false {
			t.Errorf("element %q is not present in the set", set[i])
		}
	}

	// try getting length with invalid number of params
	_, err = store.Do(testKey, Get, "noParamsShouldBePassed")
	if err == nil {
		t.Errorf("expected error; got nil while finding with invalid number of params (1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while finding with 1 parameters; got %v", err)
	}
}

func TestSet_JSON(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestSet()
	testInsertSetHelper(store, testVals, t)

	obj, err := store.ToJSON(testKey)
	if err != nil {
		t.Errorf("ToJSON returned unexpected error: %v", err)
	}

	// add a new key and initiate that key FromJSON and check if the value equals
	// by invoking the "GET" action.
	newKey := "xyz"
	err = store.AddKey(newKey, Type)
	if err != nil {
		t.Fatalf("cannot add new key to the store: %v", err)
	}

	err = store.FromJSON(newKey, obj)
	if err != nil {
		t.Errorf("FromJSON returned unexpected error: %v", err)
	}

	v, err := store.Do(testKey, Get)
	if err != nil {
		t.Errorf("cannot GET from the store: %v", err)
	}

	str, ok := v.([]string)
	if !ok {
		t.Errorf("GET did not return a string")
	}

	if err = setEqual(str, testVals); err != nil {
		t.Errorf("expected string FromJSON: %q; got %q", testVals, str)
	}

	// NB: In sets we cannot test the full representation of JSON value,
	// since ordering of elements can change in the slice representation
	// and the map representation (Go's maps are unordered).
	//
	// To test the representation in JSON, we can create a test value with
	// only one element, hence change in ordering won't affect the outcome.
	testElem := "a"
	expectedJSON, err := json.Marshal([]string{testElem})
	if err != nil {
		t.Fatalf("cannot marshal slice to test: %v", err)
	}

	newKey = "def"
	err = store.AddKey(newKey, Type)
	if err != nil {
		t.Fatalf("cannot add new key to the store: %v", err)
	}

	_, err = store.Do(newKey, Insert, testElem)
	if err != nil {
		t.Errorf("cannot insert element into set: %v", err)
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
const testKey = "testSet"

// newTestStore creates a new store for testing.
func newTestStore() (*kiwi.Store, error) {
	schema := kiwi.Schema{testKey: Type}
	return kiwi.NewStoreFromSchema(schema)
}

// setEqual returns error if both the sets are not equal.
func setEqual(set []string, expected map[string]struct{}) error {
	if len(set) != len(expected) {
		return fmt.Errorf("expected length: %d; got length: %d", len(expected), len(set))
	}

	var ok bool

	for i := 0; i < len(set); i++ {
		_, ok = expected[set[i]]
		if ok == false {
			return fmt.Errorf("value %q is not in set", set[i])
		}
	}

	return nil
}

// newTestSet gets a new set that can be used for testing.
func newTestSet() map[string]struct{} {
	var m = make(map[string]struct{})
	m["a"] = struct{}{}
	m["b"] = struct{}{}
	m["c"] = struct{}{}

	return m
}

// setToIFace converts map[string]struct{} to []interface{}.
func setToIFace(set map[string]struct{}) []interface{} {
	v := make([]interface{}, len(set))
	i := 0
	for key := range set {
		v[i] = key
		i++
	}

	return v
}

// testInsertSetHelper tests if insertion works. This can later be reused with other tests.
func testInsertSetHelper(store *kiwi.Store, testVals map[string]struct{}, t *testing.T) {
	v, err := store.Do(testKey, Insert, setToIFace(testVals)...)
	if err != nil {
		t.Errorf("error doing: %v", err)
	}

	set, ok := v.([]string)
	if !ok {
		t.Errorf("Insert action did not return a %T; rather got %T", set, v)
	}

	err = setEqual(set, testVals)
	if err != nil {
		t.Errorf("Insert did not return equal sets: %v", err)
	}
}
