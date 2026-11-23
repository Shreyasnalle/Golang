package main

import "fmt"

// closure_adder -> pos and neg are two separate closures, each has its own private sum
func closure_adder() {
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

// fibonacci -> returns a closure, every call gives the next fibonacci number
// a and b live on between calls because the inner function still references them
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

// counter -> classic closure, count is hidden, the ONLY way to change it is by calling the returned function
// this is how go does "private state" without classes
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// closure_modifies_outer -> a closure can change a variable of the function it is written in
func closure_modifies_outer() {
	total := 0
	addToTotal := func(n int) {
		total += n // changes the outer total, not a copy
	}
	addToTotal(5)
	addToTotal(10)
	fmt.Println("total changed by the closure:", total)
}
