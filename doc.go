/*
Package go_conditionals provides utility functions for conditional checks and value comparisons in Go.

# Features

Currently, this package provides the following functionality:

  - IsZeroValue: A function to check if a given value is the zero value for its type.

# Installation

To use this package in your Go project, you can install it using:

	go get github.com/ilhamtubagus/go-conditionals

# Usage

Here's an example of how to use the IsZeroValue function:

	package main

	import (
	    "fmt"
	    "github.com/yourusername/go-conditionals"
	)

	func main() {
	    fmt.Println(go_conditionals.IsZeroValue(0))          // true
	    fmt.Println(go_conditionals.IsZeroValue(""))         // true
	    fmt.Println(go_conditionals.IsZeroValue(false))      // true
	    fmt.Println(go_conditionals.IsZeroValue([]int{}))    // true
	    fmt.Println(go_conditionals.IsZeroValue(struct{}{})) // true

	    fmt.Println(go_conditionals.IsZeroValue(1))          // false
	    fmt.Println(go_conditionals.IsZeroValue("hello"))    // false
	    fmt.Println(go_conditionals.IsZeroValue(true))       // false
	    fmt.Println(go_conditionals.IsZeroValue([]int{1,2})) // false
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

# Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

# License

This project is licensed under the MIT License - see the LICENSE file for details.
*/
package go_conditionals

// This file is left intentionally empty.
// It's used only for package documentation.
