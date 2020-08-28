# Your first store

In this tutorial we will learn how to use Kiwi by creating an app to keep
a track of students in a school.

::: tip Note
This tutorial gives an overview of how to use Kiwi and does not focus on
creating a multi-threaded application. Nevertheless, there is no difference
in using Kiwi when dealing with multiple threads.
:::

## Creating project

```sh
mkdir kiwi-example
cd kiwi-example
```

Install Kiwi:

```sh
go get -u github.com/sdslabs/kiwi
```

Create the main file (entry-point of our app): `main.go`:

```go
package main

func main() {
}
```

## New store

All keys are accessed through the store. That's how we make sure every action
is safe to access. So let's create our first store:

```go
store := kiwi.NewStore()
```

That's it. Now you can add keys, update their values or use them. Let's add
a key that stores the name of the school.

## Adding key

Each value associated with a key in Kiwi has a type. In this case we need a
string to store the name of the school. But before we create a string key
we need to tell Kiwi about the string value type. Fortunately, it's just one
line. Simply import the `github/sdslabs/kiwi/values/str` package:

```go
import _ "github.com/sdslabs/kiwi/values/str"
```

And now we can add a string key in our program:

```go
if err := store.AddKey("school_name", "str"); err != nil {
  panic(err)
}
```

***

## Final program

```go
package main

import (
  "github.com/sdslabs/kiwi"
  _ "github.com/sdslabs/kiwi/values/str"
)

func main() {
  store := kiwi.NewStore()

  if err := store.AddKey("school_name", "str"); err != nil {
    panic(err)
  }
}
```
