---
home: true
heroImage: /kiwi-logo.png
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

***

### Contact

If you have a query regarding the product or just want to say hello then feel
free to visit [chat.sdslabs.co](https://chat.sdslabs.co) or drop a mail at
[contact@sdslabs.co.in](mailto:contact@sdslabs.co.in)

<p align="center">Made by <a href="https://sdslabs.co" target="_blank">SDSLabs<svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" x="0px" y="0px" viewBox="0 0 100 100" width="15" height="15" class="icon outbound"><path fill="currentColor" d="M18.8,85.1h56l0,0c2.2,0,4-1.8,4-4v-32h-8v28h-48v-48h28v-8h-32l0,0c-2.2,0-4,1.8-4,4v56C14.8,83.3,16.6,85.1,18.8,85.1z"></path> <polygon fill="currentColor" points="45.7,48.7 51.3,54.3 77.2,28.5 77.2,37.2 85.2,37.2 85.2,14.9 62.8,14.9 62.8,22.9 71.5,22.9"></polygon></svg></a></p>
