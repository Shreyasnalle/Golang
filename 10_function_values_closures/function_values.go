package main

import (
	"fmt"
	"math"
)

// asynchronous functions is a function which uses one function or calls one function in another function and this is a very common ad importatn feature of any programming language, i mean yeah, it's really important
// NOTE (correct name): passing a function into another function is called a FUNCTION VALUE / HIGHER ORDER FUNCTION / callback, not asynchronous
// asynchronous means "runs in the background without blocking", and in go that is done with goroutines -> see 14_goroutines
// in go, functions are values just like int or string, we can store them in variables, pass them as arguments and return them from functions

// compute -> takes a FUNCTION as its parameter (fn), whose type is func(float64, float64) float64, and calls it with 3 and 4
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// asynchronous_functioon -> stores an anonymous function in hypot, calls it directly, then passes it to compute
func asynchronous_functioon() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt((x*x + y*y))
	}
	fmt.Println(hypot(5, 12))
	// so as we can see here we have used a function and stored that returned value in hypot and then passsed down the paramters to that object of the function to find the square root
	fmt.Println(compute(hypot))
}
