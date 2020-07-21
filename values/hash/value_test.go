// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package hash

import (
	"bytes"
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/sdslabs/kiwi"
)

func TestHash_Insert(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()
	key := "d"
	value := "w"

	// to insert the hashmap into the store
	testInsertHashHelper(store, testVals, t)

	v, err := store.Do(testKey, Insert, key, value)
	if err != nil {
		t.Errorf("could not set hashmap[%q] from store: %v", key, err)
	}

	str, ok := v.(string)
	if !ok {
		t.Errorf("Insert did not return string rather got %T", v)
	}

	if str != key {
		t.Errorf("Insert[%q, %q]: expected %q; got %q", key, value, key, str)
	}

	// same change in testVals
	testVals[key] = value

	// trying to insert with invalid parameter type
	_, err = store.Do(testKey, Insert, true, false)
	if err == nil {
		t.Errorf("expected error; got nil while setting with key and value as bool")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while setting with key and value as bool; got %v", err)
	}

	// trying with invalid number of parameters
	_, err = store.Do(testKey, Insert)
	if err == nil {
		t.Errorf("expected error; got nil while setting with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while setting with 0 parameters; got %v", err)
	}
}

func TestHash_Remove(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()
	toRemove := "b"

	// to insert the hashmap into the store
	testInsertHashHelper(store, testVals, t)

	v, err := store.Do(testKey, Remove, toRemove)
	if err != nil {
		t.Errorf("could not remove elements from the hashmap: %v", err)
	}

	str, ok := v.([]string)
	if !ok {
		t.Errorf("expected []string from Remove; got %T", v)
	}

	if str[0] != toRemove {
		t.Errorf("key removed is not %q; got %q", toRemove, str[0])
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

func TestHash_Has(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()
	toFind := "a"

	// to insert the hashmap into the store
	testInsertHashHelper(store, testVals, t)

	v, err := store.Do(testKey, Has, toFind)
	if err != nil {
		t.Errorf("cannot find %q key in the map: %v", toFind, err)
	}

	find, ok := v.(bool)
	if !ok {
		t.Errorf("expected boolean; got %T while finding", v)
	}

	if find == false {
		t.Errorf("testVals has %q but returned false", toFind)
	}

	// finding something that does not exist
	v, err = store.Do(testKey, Has, "randomStringThatDoNotExist")
	if err != nil {
		t.Errorf("cannot find invalid value in the map: %v", err)
	}

	find, ok = v.(bool)
	if !ok {
		t.Errorf("expected boolean; got %T while finding", v)
	}

	if find == true {
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

func TestHash_Len(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()

	// to insert the hashmap into the store
	testInsertHashHelper(store, testVals, t)

	v, err := store.Do(testKey, Len)
	if err != nil {
		t.Errorf("could not get hashmap's length from store: %v", err)
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

func TestHash_Get(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()
	toGet := "a"

	// to insert the hashmap into the store
	testInsertHashHelper(store, testVals, t)

	v, err := store.Do(testKey, Get, toGet)
	if err != nil {
		t.Errorf("could not get %q value from store: %v", toGet, err)
	}

	str, ok := v.([]string)
	if !ok {
		t.Errorf("expected Get to return []string; got %T", v)
	}

	if str[0] != testVals[toGet] {
		t.Errorf("value expected is %q; got %q", testVals[toGet], str[0])
	}

	// try getting value with invalid number of params
	_, err = store.Do(testKey, Get)
	if err == nil {
		t.Errorf("expected error; got nil while finding with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while finding with 0 parameters; got %v", err)
	}

	// trying to get value with invalid parameter type
	_, err = store.Do(testKey, Get, 123)
	if err == nil {
		t.Errorf("expected error; got nil while finding with string as int")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while finding with string as int; got %v", err)
	}
}

func TestHash_Key(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()

	// to insert the hashmap into the store
	testInsertHashHelper(store, testVals, t)

	v, err := store.Do(testKey, Keys)
	if err != nil {
		t.Errorf("could not get hashmap's keys from store: %v", err)
	}

	out, ok := v.([]string)
	if !ok {
		t.Errorf("expected key to return []string; got %T", v)
	}

	if len(out) != len(testVals) {
		t.Errorf("expected length: %d; got %d", len(testVals), len(out))
	}

	for i := 0; i < len(out); i++ {
		_, ok = testVals[out[i]]
		if ok == false {
			t.Errorf("key %q is not present in the map", out[i])
		}
	}

	// try getting keys with invalid number of params
	_, err = store.Do(testKey, Keys, "noParamsShouldBePassed")
	if err == nil {
		t.Errorf("expected error; got nil while finding with invalid number of params (1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while finding with 1 parameters; got %v", err)
	}
}

func TestHash_Map(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()

	// to insert the hashmap into the store
	testInsertHashHelper(store, testVals, t)

	v, err := store.Do(testKey, Map)
	if err != nil {
		t.Errorf("could not get hashmap's copy from store: %v", err)
	}

	m, ok := v.(map[string]string)
	if !ok {
		t.Errorf("expected Map to return map[string]string; got %T", v)
	}

	eq := reflect.DeepEqual(m, testVals)
	if !eq {
		t.Errorf("returned map is not equal to actual map")
	}

	// try getting length with invalid number of params
	_, err = store.Do(testKey, Map, "noParamsShouldBePassed")
	if err == nil {
		t.Errorf("expected error; got nil while finding with invalid number of params (1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while finding with 1 parameters; got %v", err)
	}
}

func TestHash_JSON(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestHash()
	testInsertHashHelper(store, testVals, t)

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

	v, err := store.Do(testKey, Map)
	if err != nil {
		t.Errorf("cannot GET from the store: %v", err)
	}

	str, ok := v.(map[string]string)
	if !ok {
		t.Errorf("GET did not return a map[string]string")
	}

	if !reflect.DeepEqual(str, testVals) {
		t.Errorf("expected string FromJSON: %q; got %q", testVals, str)
	}

	// NB: In hashes we cannot test the full representation of JSON value,
	// since ordering of elements can change in the JSON representation
	// and the map representation (Go's maps are unordered).
	//
	// To test the representation in JSON, we can create a test value with
	// only one key-value, hence change in ordering won't affect the outcome.
	testK, testV := "a", "b"
	expectedJSON, err := json.Marshal(map[string]string{testK: testV})
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
		t.Errorf("cannot insert element into hash: %v", err)
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
const testKey = "testHash"

// newTestStore creates a new store for testing.
func newTestStore() (*kiwi.Store, error) {
	schema := kiwi.Schema{testKey: Type}
	return kiwi.NewStoreFromSchema(schema)
}

// newTestHash gets a new hashmap that can be used for testing.
func newTestHash() map[string]string {
	var m = make(map[string]string)
	m["a"] = "x"
	m["b"] = "y"
	m["c"] = "z"

	return m
}

// testInsertHashHelper tests if insertion works. This can later be reused with other tests.
func testInsertHashHelper(store *kiwi.Store, testVals map[string]string, t *testing.T) {
	var (
		v   interface{}
		err error
	)

	for key := range testVals {
		v, err = store.Do(testKey, Insert, key, testVals[key])
		if err != nil {
			t.Errorf("error doing: %v", err)
		}

		retKey, ok := v.(string)
		if !ok {
			t.Errorf("Insert action did not return a %T; rather got %T", retKey, v)
		}

		if retKey != key {
			t.Errorf("Insert did not return correct key")
		}
	}
}
