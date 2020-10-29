# Stdkiwi

Store's `Do` method takes parameters (other than the first two) as
`interface{}` which means you can pass anything into it. This isn't very
safe as this can cause many unwanted errors by either passing wrong number
of params or the wrong type. Package
[github.com/sdslabs/kiwi/stdkiwi](https://pkg.go.dev/github.com/sdslabs/kiwi/stdkiwi)
exists to solve the same issue.

## Port from `kiwi.Store`

`stdkiwi.Store` is fully compatible with `kiwi.Store`. The latter can be
replaced by the former without breaking anything. Even for creating new store
both the packages have same functions:

```go
import (
  "github.com/sdslabs/kiwi"
  "github.com/sdslabs/kiwi/stdkiwi"
)

// ...

schema := kiwi.Schema{
  "key_one": "str",
  "key_two": "hash",
}

store, err := stdkiwi.NewStore(schema)
if err != nil {
  // handle error
}

if err := store.AddKey("key_three", "set"); err != nil {
  // handle error
}

jsondata, err := store.Export()
if err != nil {
  // handle error
}
```

## Type safe methods

Package `stdkiwi` registers all the standard value types and defines various
types which can be used to invoke or `Do` actions with type safe methods.

Each of the value type has a different Go type defined in the `stdkiwi`
package and a new value of that type can be created with the corresponding
method. For example:

```go
import (
  "fmt"

  "github.com/sdslabs/kiwi/stdkiwi"
)

// ...

keyTwo := store.Hash("key_two") // assumes "key_two" is a "hash"

// All actions for a set are defined as functions
// Let's say we want to insert into the set
if err := keyTwo.Insert("key", "val"); err != nil {
  // handle error
}

// Or checking if a key exists in the hash,
// return value is not an interface{} now either
exists, err := keyTwo.Has("key")
if err != nil {
  // handle error
}

fmt.Println(exists) // true
```

The value types registered with stdkiwi are:

| Type  | Package                                                                                         | Type    | Method  |
| ----- | ----------------------------------------------------------------------------------------------- | ------- | ------- |
| str   | [github.com/sdslabs/kiwi/values/str](https://pkg.go.dev/github.com/sdslabs/kiwi/values/str)     | `Str`   | `Str`   |
| list  | [github.com/sdslabs/kiwi/values/list](https://pkg.go.dev/github.com/sdslabs/kiwi/values/list)   | `List`  | `List`  |
| set   | [github.com/sdslabs/kiwi/values/set](https://pkg.go.dev/github.com/sdslabs/kiwi/values/set)     | `Set`   | `Set`   |
| hash  | [github.com/sdslabs/kiwi/values/hash](https://pkg.go.dev/github.com/sdslabs/kiwi/values/hash)   | `Hash`  | `Hash`  |
| zset  | [github.com/sdslabs/kiwi/values/zset](https://pkg.go.dev/github.com/sdslabs/kiwi/values/zset)   | `Zset`  | `Zset`  |
| zhash | [github.com/sdslabs/kiwi/values/zhash](https://pkg.go.dev/github.com/sdslabs/kiwi/values/zhash) | `Zhash` | `Zhash` |


## Guards

Calling methods to create the special type basically assumes that type of
value. Like calling `store.List("key_one")` assumes that `"key_one"` is a
list where infact it can be any other type. As an additional layer of safety
these types implement `stdkiwi.Value` interface which implements  the `Guard`
method.

Let's say `"key_one"` is of `"str"` type and `store.List("key_one")` would
result in error when any action is invoked. To protect from this, you can
call the `Guard` method which will panic in case of such event.

```go
import "github.com/sdslabs/kiwi/stdkiwi"

// ...

keyOne := store.List("key_one")

keyOne.Guard() // Panics!
```

::: tip Tip
It is a good practice to call guard after creating the store when new
instance of the application is created. If a store is created dynamically,
use `GuardE` which is another version of `Guard` which returns the error
instead of `panic`ing.
:::
