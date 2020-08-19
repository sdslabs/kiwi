# Standard Kiwi

We have the school name but we want to keep a list of students. Not only
that but we need to store each student's ID. Each student may or may not
have unique names but they definitely will have unique IDs. Let us store
students IDs corresponding to their names in a hash map.

## Old practices

### Add a new key to store students

```go
// Imports
import "github.com/sdslabs/kiwi/values/hash"

// Main
if err := store.AddKey("students", hash.Type); err != nil {
  panic(err)
}
```

::: tip Note
When schema of the store is known, we can create the store by defining a
schema:

```go
schema := kiwi.Schema{
  "school_name": str.Type,
  "students":    hash.Type,
}

store, err := kiwi.NewStoreFromSchema(schema)
if err != nil {
  panic(err)
}
```
:::

### Add and delete students

```go
_, err := store.Do("students", hash.Insert, 123, "SDSLabs")
if err != nil {
  panic(err)
}
```

::: danger Error
Hash is a string-string map. When we call the above function, the **INSERT**
action expects a string key and string value, but instead it gets an integer
which results in an error.
:::

## Solution

One work-around for this can be defining type safe wrappers and use them
instead of `Do`ing all the time. Well... we did it for you. The
[stdkiwi](https://pkg.go.dev/github.com/sdslabs/kiwi/stdkiwi) package
defines such types and methods so you can invoke actions without encountering
unnecessary errors. Let's see how we can update our code to use this package.

### New stdkiwi store

Firstly we need to update our store. Instead of using `kiwi.Store`, we need
to create a `stdkiwi.Store`. It's pretty simple -- just replace `kiwi.NewStore`
or `kiwi.NewStoreFromSchema` with `stdkiwi.NewStore` or `stdkiwi.NewStoreFromSchema`
respectively.

```go
// Imports
import "github.com/sdslabs/kiwi/stdkiwi"

// Main
store := stdkiwi.NewStore()
```

Don't worry. Changing into `stdkiwi.Store` won't break any of the previous
code. This store is defined as such that it is compatible with the `kiwi.Store`.

::: tip Note
You can also clean-up your imports. `stdkiwi` imports all the standard value
types, i.e., the ones defined in [github.com/sdslabs/kiwi/values](https://github.com/sdslabs/kiwi/tree/main/values).
:::

### Invoking actions

Now that we have our special store, we can use the types and methods it
provides us to safely execute our actions.

```go
students := store.Hash("students") // assumes "students" key is of hash type

if err := students.Insert("123", "SDSLabs"); err != nil {
  panic(err)
}
```

***

## Final program

```go
package main

import (
  "github.com/sdslabs/kiwi"
  "github.com/sdslabs/kiwi/stdkiwi"
  "github.com/sdslabs/kiwi/values/str"
  "github.com/sdslabs/kiwi/values/hash"
)

func main() {
  schema := kiwi.Schema{
    "school_name": str.Type,
    "students":    hash.Type,
  }
  store, err := stdkiwi.NewStoreFromSchema(schema)
  if err != nil {
    panic(err)
  }

  _, err = store.Do("school_name", str.Update, "My School Name")
  if err != nil {
    panic(err)
  }

  students := store.Hash("students") // assumes "students" key is of hash type

  if err := students.Insert("123", "SDSLabs"); err != nil {
    panic(err)
  }

  if err := students.Insert("007", "Kiwi"); err != nil {
    panic(err)
  }
}
```
