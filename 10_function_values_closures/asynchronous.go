package main

import (
	"fmt"
	"math"
)

// asynchronous functions is a function which uses one function or calls one function in another function and this is a very common ad importatn feature of any programming language, i mean yeah, it's really important

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func asynchronous_functioon() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt((x*x + y*y))
	}
	fmt.Println(hypot(5, 12))
	// so as we can see here we have used a function and stored that returned value in hypot and then passsed down the paramters to that object of the function to find the square root
	fmt.Println(compute(hypot))
}