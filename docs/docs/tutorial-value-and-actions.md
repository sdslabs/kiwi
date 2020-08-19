# Value and actions

So far we have created a store and added a key inside it which has a value
of type string (or `"str"`). Now we'll update the value corresponding with
our key.

## `Do`ing

In Kiwi, anything that is to be "done" with the value is invoked through
the `Do` method of the store. The thing that invokes a particular behaviour
is called an **action**. Each value type supports multiple actions.

String value type supports two actions:
1. **GET:** Gets the value of the string.
2. **UPDATE:** Updates the value of the string.

We need to update the name of our school, so we will invoke the UPDATE action.

```go
_, err := store.Do("school_name", "UPDATE", "My School Name")
if err != nil {
  panic(err)
}
```

::: warning Exploit everything
Using strings as actions can lead to unexpected errors. Each value type in
`github.com/sdslabs/kiwi/values` declares constants that can be used instead
of writing strings for actions (and even type names).

So the above code will look like:

```go
_, err := store.Do("school_name", str.Update, "My School Name")
```
:::

## Verifying

To check if the key was updated, we can use the GET action.

```go
v, err := store.Do("school_name", str.Get)
if err != nil {
  panic(err)
}

mySchoolName := v.(string)

fmt.Println(mySchoolName) // My School Name
```
***

## Final program

```go
package main

import (
  "fmt"

  "github.com/sdslabs/kiwi"
  "github.com/sdslabs/kiwi/values/str"
)

func main() {
  store := kiwi.NewStore()

  if err := store.AddKey("school_name", str.Type); err != nil {
    panic(err)
  }

  _, err := store.Do("school_name", str.Update, "My School Name")
  if err != nil {
    panic(err)
  }

  v, err := store.Do("school_name", str.Get)
  if err != nil {
    panic(err)
  }

  mySchoolName := v.(string)

  fmt.Println(mySchoolName) // My School Name
}
```
