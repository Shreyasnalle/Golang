package main

import "fmt"

// main -> runs every pointer example in this folder
func main() {
	fmt.Println("--- reading and writing through a pointer")
	pointer_working()
	fmt.Println("--- value vs pointer when passing to a function")
	pass_by_value_vs_pointer()
	fmt.Println("--- printing addresses")
	print_addresses()
	fmt.Println("--- nil pointer")
	nil_pointer()
	fmt.Println("--- new()")
	using_new()
}
