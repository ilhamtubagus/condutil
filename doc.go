/*
Package condutil provides utility functions for conditional checks and value comparisons in Go.

# Features

Currently, this package provides the following functionality:

  - IsZeroValue: A function to check if a given value is the zero value for its type.
  - IsEmptySlice: A function to check if a given slice is empty.
  - IsEmptyMap: A function to check if a given map is empty.

# Installation

To use this package in your Go project, you can install it using:

	go get github.com/ilhamtubagus/go-conditionals

# Usage

Here's an example of how to use the functions:

	package main

	import (
	    "fmt"
	    "github.com/ilhamtubagus/go-conditionals"
	)

	func main() {
	    // IsZeroValue examples
	    fmt.Println(go_conditionals.IsZeroValue(0))          // true
	    fmt.Println(go_conditionals.IsZeroValue(""))         // true
	    fmt.Println(go_conditionals.IsZeroValue(1))          // false

	    // IsEmptySlice examples
	    fmt.Println(go_conditionals.IsEmptySlice([]int{}))   // true
	    fmt.Println(go_conditionals.IsEmptySlice([]int{1}))  // false

	    // IsEmptyMap examples
	    fmt.Println(go_conditionals.IsEmptyMap(map[string]int{}))  // true
	    fmt.Println(go_conditionals.IsEmptyMap(map[string]int{"a": 1}))  // false
	}

# Function Documentation

## IsZeroValue

	func IsZeroValue(value interface{}) bool

IsZeroValue checks if the given value is the zero value for its type.

It supports all Go types including:
  - Booleans
  - Integers (all sizes, signed and unsigned)
  - Floats
  - Complex numbers
  - Strings
  - Pointers
  - Arrays
  - Slices
  - Maps
  - Structs
  - Channels
  - Functions

Parameters:
  - value: interface{} - The value to be checked. Can be of any type.

Returns:
  - bool: true if value is the zero value for its type, false otherwise.

## IsEmptySlice

	func IsEmptySlice(slice interface{}) bool

IsEmptySlice checks if the given slice is empty.

Parameters:
  - slice: interface{} - The value to be checked. Should be a slice of any type.

Returns:
  - bool: true if slice is empty or nil, false otherwise.

## IsEmptyMap

	func IsEmptyMap(m interface{}) bool

IsEmptyMap checks if the given map is empty.

Parameters:
  - m: interface{} - The value to be checked. Should be a map of any type.

Returns:
  - bool: true if m is a map and contains no elements, false otherwise.

# Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

# License

This project is licensed under the MIT License - see the LICENSE file for details.
*/
package condutil

// This file is left intentionally empty.
// It's used only for package documentation.
