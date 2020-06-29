// Copyright (c) 2020 SDSLabs
// Use of this source code is governed by an MIT license
// details of which can be found in the LICENSE file.

// Package kiwi implements a minimalistic in-memory key value store.
//
// Each key is thread safe as it is protected by its own mutex, though different
// keys can be accessed by various threads.
//
// To get started, create a store with the NewStore function and add keys to it
// using AddKey. Each key is associated with a value which has a specific type.
// These types are extendible and can be created by implementing the Value interface.
//
// Store can also be initialized with a schema, which is basically a map of keys
// and value types.
package kiwi
