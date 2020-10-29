// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

package benchmark

import (
	"testing"

	"github.com/tidwall/buntdb"
)

const (
	buntdbTestKey = "buntdbTestKey"
	buntdbTestVal = "buntdbTestVal"
)

func BenchmarkBuntDB_Update(b *testing.B) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		b.Fatalf("couldn't open db: %v", err)
	}
	defer db.Close()

	if err := db.Update(func(tx *buntdb.Tx) error {
		for i := 0; i < b.N; i++ {
			if _, _, err := tx.Set(buntdbTestKey, buntdbTestVal, nil); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		b.Fatalf("cannot update: %v", err)
	}
}

func BenchmarkBuntDB_View(b *testing.B) {
	db, err := buntdb.Open(":memory:")
	if err != nil {
		b.Fatalf("couldn't open db: %v", err)
	}
	defer db.Close()

	if err := db.View(func(tx *buntdb.Tx) error {
		for i := 0; i < b.N; i++ {
			_, err := tx.Get(buntdbTestKey)
			if err != nil && err != buntdb.ErrNotFound {
				return err
			}
		}
		return nil
	}); err != nil {
		b.Fatalf("cannot update: %v", err)
	}
}
