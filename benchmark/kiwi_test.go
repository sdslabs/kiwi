// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package benchmark

import (
	"testing"

	"github.com/sdslabs/kiwi/stdkiwi"

	"github.com/sdslabs/kiwi"
	"github.com/sdslabs/kiwi/values/str"
)

const (
	kiwiTestKey = "kiwiTestKey"
	kiwiTestVal = "kiwiTestVal"
)

func BenchmarkKiwi_Update(b *testing.B) {
	store, err := stdkiwi.NewStoreFromSchema(kiwi.Schema{
		kiwiTestKey: str.Type,
	})
	if err != nil {
		b.Fatalf("cannot create store: %v", err)
	}
	testStr := store.Str(kiwiTestKey)

	for i := 0; i < b.N; i++ {
		if err := testStr.Update(kiwiTestVal); err != nil {
			b.Fatalf("couldn't update data: %v", err)
		}
	}
}

func BenchmarkKiwi_Get(b *testing.B) {
	store, err := stdkiwi.NewStoreFromSchema(kiwi.Schema{
		kiwiTestKey: str.Type,
	})
	if err != nil {
		b.Fatalf("cannot create store: %v", err)
	}
	testStr := store.Str(kiwiTestKey)

	for i := 0; i < b.N; i++ {
		if _, err := testStr.Get(); err != nil {
			b.Fatalf("couldn't get data: %v", err)
		}
	}
}
