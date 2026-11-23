package main

import "fmt"

// main -> runs every interface, stringer and error example in this folder
func main() {
	fmt.Println("--- basic interface")
	interface_working()
	fmt.Println("--- polymorphism with shape")
	polymorphism()
	fmt.Println("--- empty interface / any")
	empty_interface()
	fmt.Println("--- type assertion")
	type_assertion()
	fmt.Println("--- type switch")
	describe(21)
	describe("gopher")
	describe(circle{2})
	describe(3.14)
	fmt.Println("--- nil interface")
	nil_interface()
	fmt.Println("--- stringer interface")
	stringe_working()
	fmt.Println("--- error interface (custom error)")
	error_working()
	fmt.Println("--- error patterns")
	errors_patterns()
}
