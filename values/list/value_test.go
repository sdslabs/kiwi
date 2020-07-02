// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package list

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"

	"github.com/sdslabs/kiwi"
)

func TestList_Append(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	testAppendSliceHelper(store, testVals, t)

	_, err = store.Do(testKey, Append, 123)
	if err == nil {
		t.Errorf("expected error; got nil while appending integer to list")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while appending integer; got %v", err)
	}
}

func TestList_Slice(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	// checks slice for no params (getting all the elements of array)
	testAppendSliceHelper(store, testVals, t)

	// trying to slice with invalid indexes
	_, err = store.Do(testKey, Slice, -1)
	if err == nil {
		t.Errorf("expected error; got nil while slicing invalid index (0, -1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while slicing invalid index (0, -1); got %v", err)
	}

	_, err = store.Do(testKey, Slice, 1, 11)
	if err == nil {
		t.Errorf("expected error; got nil while slicing invalid index (1, 11)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while slicing invalid index (1, 11); got %v", err)
	}

	_, err = store.Do(testKey, Slice, -1, 7)
	if err == nil {
		t.Errorf("expected error; got nil while slicing invalid index (-1, 7)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while slicing invalid index (-1, 7); got %v", err)
	}

	// trying to slice with invalid parameter type
	_, err = store.Do(testKey, Slice, "1")
	if err == nil {
		t.Errorf("expected error; got nil while slicing with index as string")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while slicing with index as string; got %v", err)
	}
}

func TestList_Get(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	// to append the slice into the store
	testAppendSliceHelper(store, testVals, t)

	idx := rand.Intn(len(testVals))

	// get idx from store
	v, err := store.Do(testKey, Get, idx)
	if err != nil {
		t.Errorf("could not get list[%d] from store: %v", idx, err)
	}

	str, ok := v.(string)
	if !ok {
		t.Errorf("Get did not return string rather got %T", v)
	}

	if str != testVals[idx] {
		t.Errorf("Get[%d]: expected %q; got %q", idx, testVals[idx], str)
	}

	// get last elem when no param is given
	idx = len(testVals) - 1

	v, err = store.Do(testKey, Get)
	if err != nil {
		t.Errorf("could not get list's last elem from store: %v", err)
	}

	str, ok = v.(string)
	if !ok {
		t.Errorf("Get did not return string rather got %T", v)
	}

	if str != testVals[idx] {
		t.Errorf("Get[]: expected %q; got %q", testVals[idx], str)
	}

	// test invalid index
	_, err = store.Do(testKey, Get, -1)
	if err == nil {
		t.Errorf("expected error; got nil while getting invalid index (-1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while getting invalid index (-1); got %v", err)
	}

	_, err = store.Do(testKey, Get, 11)
	if err == nil {
		t.Errorf("expected error; got nil while getting invalid index (11)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while getting invalid index (11); got %v", err)
	}

	// trying to get with invalid parameter type
	_, err = store.Do(testKey, Get, "1")
	if err == nil {
		t.Errorf("expected error; got nil while getting with index as string")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while getting with index as string; got %v", err)
	}
}

func TestList_Set(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	// to append the slice into the store
	testAppendSliceHelper(store, testVals, t)

	idx := rand.Intn(len(testVals))
	toSet := "random"

	// set idx from store to "random"
	v, err := store.Do(testKey, Set, idx, toSet)
	if err != nil {
		t.Errorf("could not set list[%d] from store: %v", idx, err)
	}

	str, ok := v.(string)
	if !ok {
		t.Errorf("Set did not return string rather got %T", v)
	}

	if str != toSet {
		t.Errorf("Set[%d, %q]: expected %q; got %q", idx, toSet, testVals[idx], str)
	}

	// same change in testVals
	testVals[idx] = toSet

	// get last elem when no param is given
	idx = len(testVals) - 1
	toSet = "newRandom"

	v, err = store.Do(testKey, Set, toSet)
	if err != nil {
		t.Errorf("could not set list's last elem from store: %v", err)
	}

	str, ok = v.(string)
	if !ok {
		t.Errorf("Set did not return string rather got %T", v)
	}

	if str != toSet {
		t.Errorf("Set[%q]: expected %q; got %q", toSet, testVals[idx], str)
	}

	testVals[idx] = toSet

	// verify the complete list by slicing
	err = verifyList(store, testVals)
	if err != nil {
		t.Errorf("expected slice: %v; got error: %v", testVals, err)
	}

	// test invalid index
	_, err = store.Do(testKey, Set, -1, toSet)
	if err == nil {
		t.Errorf("expected error; got nil while getting invalid index (-1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while getting invalid index (-1); got %v", err)
	}

	_, err = store.Do(testKey, Set, 11, toSet)
	if err == nil {
		t.Errorf("expected error; got nil while getting invalid index (11)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while getting invalid index (11); got %v", err)
	}

	// trying to get with invalid parameter type
	_, err = store.Do(testKey, Set, true, false)
	if err == nil {
		t.Errorf("expected error; got nil while setting with index as bool")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while setting with index as bool; got %v", err)
	}

	// trying with invalid number of parameters
	_, err = store.Do(testKey, Set)
	if err == nil {
		t.Errorf("expected error; got nil while setting with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while setting with 0 parameters; got %v", err)
	}
}

func TestList_Len(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	// to append the slice into the store
	testAppendSliceHelper(store, testVals, t)

	v, err := store.Do(testKey, Len)
	if err != nil {
		t.Errorf("could not get list's length from store: %v", err)
	}

	l, ok := v.(int)
	if !ok {
		t.Errorf("expected Len to return int; got %T", v)
	}

	if l != len(testVals) {
		t.Errorf("expected length: %d; got %d", len(testVals), l)
	}
}

func TestList_Pop(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	// to append the slice into the store
	testAppendSliceHelper(store, testVals, t)

	// pop the last element off the list
	v, err := store.Do(testKey, Pop)
	if err != nil {
		t.Errorf("couldn't pop off the last element from the list")
	}

	slice, ok := v.([]string)
	if !ok {
		t.Errorf("expected Pop to return []string; got %T", v)
	}

	if len(slice) != 1 {
		t.Errorf("expected length of popped elems: 1; got %d", len(slice))
	}

	if slice[0] != testVals[len(testVals)-1] {
		t.Errorf("Pop[] returned invalid element: expected: %q; got: %q",
			testVals[len(testVals)-1], slice[0])
	}

	// update test vals accordingly
	testVals = testVals[:len(testVals)-1]

	n := 3
	v, err = store.Do(testKey, Pop, n)
	if err != nil {
		t.Errorf("couldn't pop off the last element from the list")
	}

	slice, ok = v.([]string)
	if !ok {
		t.Errorf("expected Pop to return []string; got %T", v)
	}

	if len(slice) != n {
		t.Errorf("expected length of popped elems: %d; got %d", n, len(slice))
	}

	for i := range slice {
		if slice[i] != testVals[len(testVals)-n+i] {
			t.Errorf("Pop[%d] returned inavlid element: expected: %q; got %q",
				n, testVals[len(testVals)-n+i], slice[i])
		}
	}

	testVals = testVals[:len(testVals)-n]

	// verify the complete list by slicing
	err = verifyList(store, testVals)
	if err != nil {
		t.Errorf("expected slice: %v; got error: %v", testVals, err)
	}

	// test invalid index
	_, err = store.Do(testKey, Pop, -1)
	if err == nil {
		t.Errorf("expected error; got nil while popping invalid number (-1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while popping invalid number (-1); got %v", err)
	}

	_, err = store.Do(testKey, Pop, 11)
	if err == nil {
		t.Errorf("expected error; got nil while popping invalid number (11)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while popping invalid number (11); got %v", err)
	}

	// trying to get with invalid parameter type
	_, err = store.Do(testKey, Pop, "1")
	if err == nil {
		t.Errorf("expected error; got nil while popping with number as string")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while popping with number as string; got %v", err)
	}
}

func TestList_Remove(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	// to append the slice into the store
	testAppendSliceHelper(store, testVals, t)

	// remove 2nd and 7th element from the list
	v, err := store.Do(testKey, Remove, 1)
	if err != nil {
		t.Errorf("could not remove elements from the list: %v", err)
	}

	str, ok := v.(string)
	if !ok {
		t.Errorf("expected []string from Remove; got %T", v)
	}

	if str != testVals[1] {
		t.Errorf("first element removed is not testVals[1] (%q); got %q", testVals[1], str)
	}

	v, err = store.Do(testKey, Remove, testVals[6])
	if err != nil {
		t.Errorf("could not remove elements from the list: %v", err)
	}

	str, ok = v.(string)
	if !ok {
		t.Errorf("expected []string from Remove; got %T", v)
	}

	if str != testVals[6] {
		t.Errorf("second element removed is not testVals[6] (%q); got %q", testVals[6], str)
	}

	testVals = append(testVals[:1], testVals[2:]...)
	testVals = append(testVals[:5], testVals[6:]...) // adjusting for already removed element

	// verify the complete list by slicing
	err = verifyList(store, testVals)
	if err != nil {
		t.Errorf("expected slice: %v; got error: %v", testVals, err)
	}

	// trying to remove invalid index
	_, err = store.Do(testKey, Remove, -1)
	if err == nil {
		t.Errorf("expected error; got nil while removing invalid index (-1)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while removing invalid index (-1); got %v", err)
	}

	_, err = store.Do(testKey, Remove, 11)
	if err == nil {
		t.Errorf("expected error; got nil while removing invalid number (11)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidIndex {
		t.Errorf("expected ErrInvalidIndex while removing invalid number (11); got %v", err)
	}

	// trying to remove with invalid parameter type
	_, err = store.Do(testKey, Remove, true)
	if err == nil {
		t.Errorf("expected error; got nil while removing with index or string as bool")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while removing with index or string as bool; got %v", err)
	}
}

func TestList_Find(t *testing.T) {
	store, err := newTestStore()
	if err != nil {
		t.Fatalf("couldn't create store: %v", err)
	}

	testVals := newTestList()

	// to append the slice into the store
	testAppendSliceHelper(store, testVals, t)

	idx := rand.Intn(len(testVals))

	v, err := store.Do(testKey, Find, testVals[idx])
	if err != nil {
		t.Errorf("cannot find testVals[%d] in the slice: %v", idx, err)
	}

	fidx, ok := v.(int)
	if !ok {
		t.Errorf("expected int; got %T while finding", v)
	}

	if fidx != idx {
		t.Errorf("finding testVals[%d] returned %d", idx, fidx)
	}

	// finding something that does not exist
	v, err = store.Do(testKey, Find, "randomStringThatShouldNotExist")
	if err != nil {
		t.Errorf("cannot find invalid value in the slice: %v", err)
	}

	fidx, ok = v.(int)
	if !ok {
		t.Errorf("expected int; got %T while finding", v)
	}

	if fidx != -1 {
		t.Errorf("finding invalid value returned %d; expected -1", fidx)
	}

	// try finding with invalid number of params
	_, err = store.Do(testKey, Find)
	if err == nil {
		t.Errorf("expected error; got nil while finding with invalid number of params (0)")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamLen {
		t.Errorf("expected ErrInvalidParamLen while finding with 0 parameters; got %v", err)
	}

	// trying to find with invalid parameter type
	_, err = store.Do(testKey, Find, 123)
	if err == nil {
		t.Errorf("expected error; got nil while finding with string as int")
	}

	if er := errors.Unwrap(err); er != ErrInvalidParamType {
		t.Errorf("expected ErrInvalidParamType while finding with string as int; got %v", err)
	}
}

// testKey to test the value.
const testKey = "testList"

// newTestStore creates a new store for testing.
func newTestStore() (*kiwi.Store, error) {
	schema := kiwi.Schema{testKey: Type}
	return kiwi.NewStoreFromSchema(schema)
}

// listEqual returns error if both the lists are not equal.
func listEqual(slice, expected []string) error {
	if len(slice) != len(expected) {
		return fmt.Errorf("expected length: %d; got length: %d", len(expected), len(slice))
	}

	for i := range expected {
		if slice[i] != expected[i] {
			return fmt.Errorf("expected index %d = %q; got %q", i, expected[i], slice[i])
		}
	}

	return nil
}

// newTestList gets a new list that can be used for testing.
func newTestList() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
}

// listToIFace converts []string to []interface{}.
func listToIFace(slice []string) []interface{} {
	v := make([]interface{}, len(slice))
	for i := range slice {
		v[i] = slice[i]
	}

	return v
}

// verifyList tests if the lsit in store and expected [] string are same.
func verifyList(store *kiwi.Store, expected []string) error {
	v, err := store.Do(testKey, Slice)
	if err != nil {
		return err
	}

	slice, ok := v.([]string)
	if !ok {
		return fmt.Errorf("Slice did not return %T", slice)
	}

	return listEqual(slice, expected)
}

// testAppendSliceHelper tests if appending works. This can later be reused with other tests.
func testAppendSliceHelper(store *kiwi.Store, testVals []string, t *testing.T) {
	v, err := store.Do(testKey, Append, listToIFace(testVals)...)
	if err != nil {
		t.Errorf("error doing: %v", err)
	}

	slice, ok := v.([]string)
	if !ok {
		t.Errorf("Append action did not return a %T; rather got %T", slice, v)
	}

	err = listEqual(slice, testVals)
	if err != nil {
		t.Errorf("Append did not return equal lists: %v", err)
	}

	err = verifyList(store, testVals)
	if err != nil {
		t.Errorf("error verifying with Slice: %v", err)
	}
}
