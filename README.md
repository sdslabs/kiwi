# kiwi

> A minimalistic in-memory key value store.

Each key is thread safe as it is protected by its own mutex, though different keys can be accessed by various threads.

To get started, create a store with the NewStore function and add keys to it using AddKey.
Each key is associated with a value which has a specific type.
These types are extendible and can be created by implementing the Value interface.

Store can also be initialised with a schema, which is basically a map of keys and value types.

```go
// If you have a pre-defined schema
schema := kiwi.Schema{
    "key1": str.Type, // string key
    // ...
}

store, err := kiwi.NewStoreFromSchema(schema)
if err != nil {
    // handle error
}

// If you want to add/update keys dynamically
store.AddKey("key2", str.Type)
store.UpdateKey("key1", list.Type)

// Invoke an action, say updating a string
_, err := store.Do("key2", str.Update, "abc")
if err != nil {
    // handle error
}

// or if you need to get the string
s, err := store.Do("key2", str.Get)
if err != nil {
    // handle error
}

myString := s.(string)
// use myString
```

Package `github.com/sdslabs/kiwi/stdkiwi` implements type safe methods with standard types registered.
These types include:
- [x] str (`github.com/sdslabs/kiwi/values/str`): String value
- [x] list (`github.com/sdslabs/kiwi/values/list`): A list of strings
- [x] set (`github.com/sdslabs/kiwi/values/set`): An exclusive set of strings
- [x] hash (`github.com/sdslabs/kiwi/values/hash`): A string-string hashmap
- [ ] zset (TODO): A set with each element having scores

If you only require the aforementioned value types, use the `stdkiwi` package.
The above example using the `stdkiwi` package is as follows:

```go
// If you have a pre-defined schema
schema := kiwi.Schema{
    "key1": str.Type, // string key
    // ...
}

store, err := stdkiwi.NewStoreFromSchema(schema)
if err != nil {
    // handle error
}


// If you want to add/update keys dynamically
store.AddKey("key2", str.Typ)
store.UpdateKey("key1", list.Typ)

// Invoke an action, say updating a string
err = store.Str("key2").Update("abc")
if err != nil {
    // handle error
}

mystring := store.Str("key2").Get()
// use mystring

// If the schema is known while starting the app, to avoid unnecessary errors during app runtime,
// use Guard method for the values when initializing the store.
str.List("key1").Guard() // panics on error

err = str.Str("key2").GuardE() // returns error
if err != nil {
	// handle error
}
```

To implement a custom value, all that is required is to implement the `kiwi.Value` interface and register
it with the kiwi store using `kiwi.RegisterValue`. Look at the `github.com/sdslabs/kiwi/values` for examples.
