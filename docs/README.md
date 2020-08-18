---
home: true
heroImage: /sdslabs-logo.png
tagline: A minimalistic in-memory key value store.
actionText: Get Started →
actionLink: /docs/
features:
- title: Extendable
  details: Custom data structures can be added with ease.
- title: One Binary
  details: Kiwi can be directly plugged in as a go library which produces one binary for your application.
- title: In-Memory
  details: All the data is stored in memory for fast access.
footer: MIT Licensed | Copyright © 2020 SDSLabs
---

Create a store, add key and play with it. It's that easy!

```go
store := stdkiwi.NewStore()

if err := store.AddKey("my_string", "str"); err != nil {
  // handle error
}

myString := store.Str("my_string")

if err := myString.Update("Hello, World!"); err != nil {
  // handle error
}

str, err := myString.Get()
if err != nil {
  // handle error
}

fmt.Println(str) // Hello, World!
```
