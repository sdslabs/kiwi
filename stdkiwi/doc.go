// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

// Package stdkiwi implements an API with standard values.
//
// These values include all the values in the package github.com/sdslabs/kiwi/values.
// It implements all the actions as type-safe functions.
//
//
// Get Started
//
// Create a store, add key and play with it. It's that easy!
//
//	store := stdkiwi.NewStore()
//
//	if err := store.AddKey("my_string", "str"); err != nil {
//	  // handle error
//	}
//
//	myString := store.Str("my_string")
//
//	if err := myString.Update("Hello, World!"); err != nil {
//	  // handle error
//	}
//
//	str, err := myString.Get()
//	if err != nil {
//	  // handle error
//	}
//
//	fmt.Println(str) // Hello, World!
//
// For documentation visit https://kiwi.sdslabs.co/docs/concepts-stdkiwi.html
package stdkiwi
