# Add new value type

A value is something that can be associated with a key in a Kiwi store.
A value implements its own methods and can be accessed by type assertion.
To define a new third part value you need to implement the `kiwi.Value` interface
which consists of type, DoMap, ToJSON and FromJSON.

## Type

Each value has a static type:
1. Has a unique name.
2. Usually named after data types that are meaningful to users.

```go
// Type of str value.
// str is an example, you can put your own type here.
func (v *Value) Type() kiwi.ValueType {
	return "str"
}
```

## DoMap

Each type implements its own actions according to the need and scope. DoMap
returns a map that maps Actions with the respective functions.

So if you want a **Get** action for str which on invoking will return the value 
of str, and an **Update** action for updating the str it can be done by mapping
the actions with corresponding functions as:

```go
// Defining actions

Get kiwi.Action = "GET"
Update kiwi.Action = "UPDATE"

// Mapping the actions with the functions defining their behaviour.
func (v *Value) DoMap() map[kiwi.Action]kiwi.DoFunc {
	return map[kiwi.Action]kiwi.DoFunc{
		Get: func(params ...interface{}) (interface{}, error) {
			// don't need any params
			return string(*v), nil
        },
        Update: func(params ...interface{}) (interface{}, error) {
			// requires one string parameter
			if len(params) < 1 {
				return nil, fmt.Errorf("str.Value.Update requires 1 argument")
			}

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

## ToJSON

ToJSON returns the raw byte array of the value's data.

```go
func (v *Value) ToJSON() (json.RawMessage, error) {
	c, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(c), nil
}
```

## FromJSON

FromJSON populates the value with the data from `RawMessage`.

```go
func (v *Value) FromJSON(rawmessage json.RawMessage) error {
	return json.Unmarshal(rawmessage, v)
}
```

## Register Value

After implementing the complete interface, the new value can be registered by
`kiwi.RegisterValue`. It registers a new value type with the package.

```go
kiwi.RegisterValue(func() kiwi.Value { return new(Value) })
```
