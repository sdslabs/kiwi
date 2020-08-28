# Add new value type

A value is something that can be associated with a key in a Kiwi store.
A value implements its own methods and can be accessed by type assertion.
To define a new third part value you need to implement the `kiwi.Value` interface
which consists of `Type`, `DoMap`, `ToJSON` and `FromJSON`.

## Type

Each value has a static type:
1. Has a unique name (since all types are recognized by their type name).
2. Usually named after data types that are meaningful to users.

```go
import "github.com/sdslabs/kiwi"

// ...

// Type of str value.
// str is an example, you can put your own type here.
func (v *Value) Type() kiwi.ValueType {
  return "str"
}
```

## DoMap

Each type implements its own actions according to the need and scope. DoMap
returns a `map` that maps actions with the respective functions.

So if you want a **GET** action for str which on invoking will return the value
of str, and an **UPDATE** action for updating the str it can be done by mapping
the actions with corresponding functions as:

```go
import (
  "fmt"

  "github.com/sdslabs/kiwi"

// ...

// Mapping the actions with the functions defining their behaviour.
func (v *Value) DoMap() map[kiwi.Action]kiwi.DoFunc {
  return map[kiwi.Action]kiwi.DoFunc{
    "GET": func(params ...interface{}) (interface{}, error) {
      // don't need any params
      return string(*v), nil
    },

    "UPDATE": func(params ...interface{}) (interface{}, error) {
      // requires one parameter
      if len(params) < 1 {
        return nil, fmt.Errorf("str.Value.Update requires 1 argument")
      }

      // parameter should be a valid string
      newStr, ok := params[0].(string)
      if !ok {
        return nil, fmt.Errorf("str.Value.Update takes only string argument")
      }

      *v = Value(newStr)
      return newStr, nil
    },
  }
}
```

Now an update action can be called on the type using the `Do` method:

```go
import "github.com/sdslabs/kiwi"

// ...

store.Do("key_name", "UPDATE", "hello world")
```

## ToJSON

ToJSON returns the data in JSON format.

```go
import (
  "encoding/json"

  "github.com/sdslabs/kiwi"
)

// ...

func (v *Value) ToJSON() (json.RawMessage, error) {
	c, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(c), nil
}
```

So, in this case, our string `hello` would be `"hello"` in JSON.

## FromJSON

FromJSON populates the value with the data from `json.RawMessage`.

```go
import (
  "encoding/json"

  "github.com/sdslabs/kiwi"
)

// ...

func (v *Value) FromJSON(rawmessage json.RawMessage) error {
	return json.Unmarshal(rawmessage, v)
}
```

So if we pass `"hello"` as raw message, it populates the data of our value
as `hello`.

## Register Value

After implementing the complete interface, the new value should be registered
by using `kiwi.RegisterValue`. It registers a new value type with the package.
This should be called in the `init` method so that if any error occurs, or
there is any conflict in type names, it is caught in during initialisation
of the app.

```go
import "github.com/sdslabs/kiwi"

// ...

func init() {
  kiwi.RegisterValue(func() kiwi.Value { return new(Value) })
}
```

This is the actual implementation for
[github.com/sdslabs/kiwi/values/str](https://pkg.go.dev/github.com/sdslabs/kiwi/values/str)
package. To see more examples, take a look at implementations of
[standard packages](https://github.com/sdslabs/kiwi/tree/main/values).
