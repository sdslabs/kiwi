# Import and export

Now that we have explored how to use Kiwi, there's just one final feature
that we should get familiar with. All Kiwi values are compatible with JSON,
i.e., they can be converted to and loaded from JSON.

## Import

Data in a Kiwi store can be imported from a JSON. For example, the following
is valid data for our store:

```json
{
  "school_name": {
    "type": "str",
    "data": "Old School Name"
  },
  "students": {
    "type": "hash",
    "data": {
      "999": "Std Kiwi"
    }
  }
}
```

Let's import this data into our store:

```go
jsonData := json.RawMessage(`
{
  "school_name": {
    "type": "str",
    "data": "Old School Name"
  },
  "students": {
    "type": "hash",
    "data": {
      "999": "Std Kiwi"
    }
  }
}
`)

if err := store.Import(jsonData, kiwi.ImportOpts{}); err != nil {
  panic(err)
}
```

## Values in JSON

Each value can also be individually imported from/exported into JSON:

```go
studentsJSON, err := store.ToJSON("students")
if err != nil {
  panic(err)
}

fmt.Println(string(studentsJSON))
// Outputs: {"007":"Kiwi","123":"SDSLabs","999":"Std Kiwi"}
```

***

## Final program

```go
package main

import (
  "encoding/json"
	"fmt"

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

  jsonData := json.RawMessage(`
{
  "school_name": {
    "type": "str",
    "data": "Old School Name"
  },
  "students": {
    "type": "hash",
    "data": {
      "999": "Std Kiwi"
    }
  }
}
`)

  if err := store.Import(jsonData, kiwi.ImportOpts{}); err != nil {
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

  studentsJSON, err := store.ToJSON("students")
  if err != nil {
    panic(err)
  }

  fmt.Println(string(studentsJSON))
  // Outputs: {"007":"Kiwi","123":"SDSLabs","999":"Std Kiwi"}
}
```
