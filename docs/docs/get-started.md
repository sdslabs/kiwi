# Get Started

::: warning Prerequisite
Kiwi requires [Go](https://golang.org/) >= 1.14
:::

## Installing

Kiwi can be integrated with your application just like any other go library.

```sh
go get -u github.com/sdslabs/kiwi
```

Now you can import kiwi any where in your code.

```go
import "github.com/sdslabs/kiwi"
```

## Quick start

To start using kiwi, you need to create a store. A store can be created using
two methods:

1. Using the `stdkiwi` package: [github.com/sdslabs/kiwi/stdkiwi](https://pkg.github.com/sdslabs/kiwi/stdkiwi)
2. Using the core `kiwi` package: [github.com/sdslabs/kiwi](https://pkg.go.dev/github.com/sdslabs/kiwi)

### Using `stdkiwi`

```go
package main

import (
  "fmt"

  "github.com/sdslabs/kiwi/stdkiwi"
)

func main() {
  // Create a new store.
  store := stdkiwi.NewStore()

  // Add key called "my_string" of "str" (string) type.
  if err := store.AddKey("my_string", "str"); err != nil {
    panic(err)
  }

  // To avoid using "my_string", declare a variable of str type.
  myString := store.Str("my_string")

  // Update "my_string" to `"Hello, World!"`.
  if err := myString.Update("Hello, World!"); err != nil {
    panic(err)
  }

  // Get value corresponding to "my_string".
  str, err := myString.Get()
  if err != nil {
    panic(err)
  }

  fmt.Println(str) // Hello, World!
}
```

### Using core package

The same program using the core package can be written as:

```go
package main

import (
  "fmt"

  "github.com/sdslabs/kiwi"

  // Importing this package enables kiwi to create string values
  "github.com/sdslabs/kiwi/values/str"
)

func main() {
  store := kiwi.NewStore()

  if err := store.AddKey("my_string", str.Type); err != nil { // str.Type = "str"
    panic(err)
  }

  _, err := store.Do("my_string", str.Update, "Hello, World!") // str.Update = "UPDATE"
  if err != nil {
    panic(err)
  }

  str, err := store.Do("my_string", str.Get) // str.Get = "GET"
  if err != nil {
    panic(err)
  }

  fmt.Println(str.(string))
}
```

::: warning Use stdkiwi
Using `stdkiwi` is preferred method if you don't need any other data types
other than the standard value types. This is because `stdkiwi` defines many
types corresponding to the standard value types, like `stdkiwi.Str`. Actions
corresponding to these types are defined as type safe methods.
:::
