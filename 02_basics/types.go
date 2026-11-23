package main

import (
	"fmt"
	"math"
)

// const -> a value which is fixed at compile time and can never be changed later, we cannot use := with const
const pi = 3.14
const greeting = "Hello"

// we can group many declarations in a () block, works for var and const both
const (
	max_int8  int8  = 127           // biggest value int8 can hold (8 bits -> -128 to 127)
	max_int64 int64 = math.MaxInt64 // biggest value int64 can hold (64 bits)
)

// constants_and_types shows constants, the sized ints and type conversion
func constants_and_types() {
	fmt.Println(greeting, pi)
	// pi = 3.1415 -> compile error, cannot assign to a constant

	fmt.Println("int8 max:", max_int8, "| int64 max:", max_int64)

	// type conversion -> go NEVER converts types automatically (unlike js), we have to write T(value) ourselves
	var i int = 42
	var f float64 = float64(i) // int -> float64
	var u uint = uint(f)       // float64 -> uint
	fmt.Println(i, f, u)

	// even int and int64 are different types, so this needs conversion too
	var big int64 = int64(i)
	fmt.Printf("%v is %T, %v is %T\n", i, i, big, big)

	// type inference -> with := go picks the type from the value itself
	a := 10   // int
	b := 10.5 // float64
	c := "go" // string
	d := true // bool
	fmt.Printf("%T %T %T %T\n", a, b, c, d)
}
