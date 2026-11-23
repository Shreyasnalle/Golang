package main

import "fmt"

// operation -> we can give a function signature its own type name, so long func(...) types become readable
type operation func(int, int) int

// apply -> uses the named function type operation as a parameter
func apply(a, b int, op operation) int {
	return op(a, b)
}

// function_types -> passing different functions to the same apply, and keeping functions inside a map
func function_types() {
	add := func(a, b int) int { return a + b }
	mul := func(a, b int) int { return a * b }
	fmt.Println("apply add:", apply(3, 4, add), "| apply mul:", apply(3, 4, mul))

	calculator := map[string]operation{
		"+": add,
		"*": mul,
		"-": func(a, b int) int { return a - b }, // anonymous function written directly as the value
	}
	fmt.Println("calculator[\"-\"](10, 3) =", calculator["-"](10, 3))
}

// immediately_invoked -> an anonymous function that is defined and called in the same line, the () at the end calls it
func immediately_invoked() {
	result := func(name string) string {
		return "hello " + name
	}("shreyas")
	fmt.Println(result)
}

// map_strings -> a small higher order function, it applies fn to every element of the slice and returns the new slice
func map_strings(items []string, fn func(string) string) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, fn(item))
	}
	return out
}

// multiplier -> a function that RETURNS a function, the returned function remembers factor (this is a closure)
func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
