package main

import (
	"fmt"
	"strings"
)

// main -> runs every function value and closure example in this folder
func main() {
	fmt.Println("--- function as a value / parameter")
	asynchronous_functioon()
	fmt.Println("--- named function types + map of functions")
	function_types()
	fmt.Println("--- immediately invoked anonymous function")
	immediately_invoked()
	fmt.Println("--- higher order function")
	fmt.Println(map_strings([]string{"go", "closure"}, strings.ToUpper))
	fmt.Println("--- function returning a function")
	double, triple := multiplier(2), multiplier(3)
	fmt.Println(double(5), triple(5))
	fmt.Println("--- closure: adder")
	closure_adder()
	fmt.Println("--- closure: fibonacci")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), " ")
	}
	fmt.Println()
	fmt.Println("--- closure: counter with private state")
	next := counter()
	fmt.Println(next(), next(), next())
	fmt.Println("--- closure changing the outer variable")
	closure_modifies_outer()
}
