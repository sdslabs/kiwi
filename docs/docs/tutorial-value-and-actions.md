# Value and actions

So far we have created a store and added a key inside it which has a value
of type string (or `"str"`). Now we'll update the value corresponding with
our key.

## Do

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

## Verify

To check if the key was updated, we can use the GET action.

```go
v, err := store.Do("school_name", "GET")
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
  _ "github.com/sdslabs/kiwi/values/str"
)

func main() {
  store := kiwi.NewStore()

  if err := store.AddKey("school_name", "str"); err != nil {
    panic(err)
  }

  _, err := store.Do("school_name", "UPDATE", "My School Name")
  if err != nil {
    panic(err)
  }

  v, err := store.Do("school_name", "GET")
  if err != nil {
    panic(err)
  }

  mySchoolName := v.(string)

  fmt.Println(mySchoolName) // My School Name
}
```
