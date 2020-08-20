# Values

A value is associated with a key in a Kiwi store. Each value has a static
type. For example, the store contains a key called `name` which should
contain the name of the application. To store this we need a string value.
Hence the key `name` will be associated with a value of type that stores
string data.

## Pluggable types

Kiwi's type system is extendable, i.e., different value types can be plugged
in as per requirement. Even the standard value types defined in the
[values](https://github.com/sdslabs/kiwi/values) package are implemented as
third party types.

So if you want to store a hashmap, the package
[github.com/sdslabs/kiwi/values/hash](https://pkg.go.dev/github.com/sdslabs/kiwi/values/hash)
defines a Kiwi value type that can store data as a hash map. This can be plugged
in simply by importing the package:

```go
import _ "github.com/sdslabs/kiwi/values/hash"

// ...

if err := store.AddKey("key_name", "hash"); err != nil {
  // handle error
}
```

## Actions

A value can be interacted with using different actions. An action is nothing
but a string telling the store about what to do with the value associated
with the key you want to interact with.

Each value type defines it's own actions. Say, you want to insert into the
hash map, so you would `Do` an `"INSERT"` action.

```go
_, err := store.Do("key_name", "INSERT", "key", "val")
if err != nil {
  // handle error
}
```

::: tip Tip
Each value type defined in the Kiwi package exposes actions and type names as
constants. So rather than blank importing, you can import the complete package
and use them such as:

```go
import "github.com/sdslabs/kiwi/values/hash"

// ...

if err := store.AddKey("key_name", hash.Type); err != nil { // hash.Type = "hash"
  // handle error
}

_, err := store.Do("key_name", hash.Insert, "key", "val") // hash.Insert = "INSERT"
if err != nil {
  // handle error
}
```
:::

## JSON

Each value is JSON compatible, i.e., they can be loaded from or converted
into JSON. This property is what makes a store importable and exportable
from and into JSON respectively.

```go
import "github.com/sdslabs/kiwi"

// ...

jsonRawData, err := store.ToJson("my_key")
if err != nil {
  // handle error
}

if err := store.FromJSON("my_key", jsonRawData); err != nil {
  // handle error
}
```
