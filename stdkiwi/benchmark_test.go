package stdkiwi

import (
	"testing"

	"github.com/sdslabs/kiwi"
	"github.com/sdslabs/kiwi/values/str"
)

func BenchmarkUpdate(b *testing.B) {
	store, err := NewStoreFromSchema(kiwi.Schema{
		testKey: str.Type,
	})
	if err != nil {
		b.Fatalf("cannot create store: %v", err)
	}
	testStr := store.Str(testKey)

	for i := 0; i < b.N; i++ {
		if err := testStr.Update("data"); err != nil {
			b.Fatalf("couldn't update data: %v", err)
		}
	}
}

func BenchmarkGet(b *testing.B) {
	store, err := NewStoreFromSchema(kiwi.Schema{
		testKey: str.Type,
	})
	if err != nil {
		b.Fatalf("cannot create store: %v", err)
	}
	testStr := store.Str(testKey)

	for i := 0; i < b.N; i++ {
		if _, err := testStr.Get(); err != nil {
			b.Fatalf("couldn't get data: %v", err)
		}
	}
}
