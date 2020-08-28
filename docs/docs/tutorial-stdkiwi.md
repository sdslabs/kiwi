# Standard Kiwi

We have the school name but we want to keep a list of students. Not only
that but we need to store each student's ID. Each student may or may not
have unique names but they definitely will have unique IDs. Let us store
students IDs corresponding to their names in a hash map.

## Old practices

### Add a new key to store students

```go
// Imports
import _ "github.com/sdslabs/kiwi/values/hash"

// Main
if err := store.AddKey("students", "hash"); err != nil {
  panic(err)
}
```

### Add student

```go
_, err := store.Do("students", "INSERT", 123, "SDSLabs")
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
  "github.com/sdslabs/kiwi/stdkiwi"
)

func main() {
  store := stdkiwi.NewStore()

  if err := store.AddKey("school_name", "str"); err != nil {
    panic(err)
  }

  if err := store.AddKey("students", "hash"); err != nil {
    panic(err)
  }

  _, err = store.Do("school_name", "UPDATE", "My School Name")
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
