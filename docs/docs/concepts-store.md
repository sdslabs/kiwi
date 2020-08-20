# Store

Store is the entry point for interacting with the values. A store protects
values when they are being used via multiple threads.

## Create new store

The simplest way to create a new store is to use `NewStore` function.
This creates an empty store and add/remove keys from it.

```go
import "github.com/sdslabs/kiwi"

// ...

store := kiwi.NewStore()
```

## Add, update and remove keys

A key can be added to the store using `AddKey` method. If the key already
exists, `AddKey` will return a non-nil error.

`UpdateKey` updates the value type of the key. If the key does not exist,
it returns a non-nil error. Calling update deletes the old value irrespective
of if the value type is same or not.

A key can be removed using `DeleteKey` method. This also returns a non-nil
error when the key does not exist.

```go
import (
  "github.com/sdslabs/kiwi"

  // Value types
  "github.com/sdslabs/kiwi/values/hash"
  "github.com/sdslabs/kiwi/values/str"
)

// ...

if err := store.AddKey("my_key", str.Type); err != nil { // str.Type = "str"
  // handle error
}

if err := store.UpdateKey("my_key", hash.Type); err != nil { // hash.Type = "hash"
  // handle error
}

if err := store.DeleteKey("my_key"); err != nil {
  // handle error
}
```

## Schema

Most of the times a schema is known for the store. A store can be created
with this schema itself.

`Schema` type is a map of keys v/s their value types. To create a store,
`NewStoreFromSchema` can be used:

```go
import (
  "github.com/sdslabs/kiwi"

  // Value types
  "github.com/sdslabs/kiwi/values/hash"
  "github.com/sdslabs/kiwi/values/str"
)

// ...

schema := kiwi.Schema{
  "my_string": str.Type,
  "my_hash":   hash.Type,
}

store, err := kiwi.NewStoreFromSchema(schema)
if err != nil {
  // handle error
}
```

The above example is equivalent to creating a new empty store and calling
`AddKey` twice.

For any store, you can also get the schema using `GetSchema` method.

# Interact with values

To mutate or read data associated with the keys, use `Do` method of the store.
It takes two parameters, one is the key name and the other is the kind of
action to be invoked.

Let's say we need to check if a string exists in a set. Value type `"set"` (in
the package `github.com/sdslabs/kiwi/values/set`) defines an action called
**HAS** to check if the value exists.

```go
import (
  "fmt"

  "github.com/sdslabs/kiwi"

  // Value types
  "github.com/sdslabs/kiwi/values/set"
)

// ...

v, err := store.Do("my_key", set.Has, "some_string") // set.Has = "HAS"
if err != nil {
  // occurance of error here means `Do` failed due to invalid code
}

fmt.Println(v.(bool)) // HAS action returns a boolean
```

`Do` returns two values. First being the interface that carries any data
associated with the action. Like in the aforementioned code, HAS returns a
boolean. The second is the error which can occur in cases when:

1. The key is invalid
2. The action is not defined for the value type
3. Type or number of parameters is wrong

Another kind of interaction can be to get the data from values in JSON format
or load the value from JSON. This can be done using the `ToJSON` and `FromJSON`
methods of the store.

More on values in the [next concept](./concepts-values.md).

## Import and export

Data from the store can be exported into JSON or imported from JSON using
the `Export` and `Import` methods respectively.

The JSON format required is as follows:

```json
{
  "key_one": {
    "type": "str",
    "data": "my string"
  },
  "key_two": {
    "type": "hash",
    "data": {
      "my": "kiwi",
      "hash": "map"
    }
  }
}
```

In the above JSON, the store has two keys: `"key_one"` and `"key_two"`. The
values associated with these keys have types `"str"` and `"hash"` respectively.
`"data"` contains the data correspondind to the keys.

Type `StoreJSON` can be marshalled into the above format.
