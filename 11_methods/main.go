package main

import "fmt"

// main -> runs every method example in this folder
func main() {
	fmt.Println("--- function vs method")
	running()
	running_method()
	fmt.Println("--- pointer receiver then value receiver")
	normal_variable_pointer_variable_method()
	fmt.Println("--- value receiver cannot change the struct")
	value_receiver_does_not_mutate()
	fmt.Println("--- automatic & and *")
	auto_address_and_deref()
	fmt.Println("--- methods on non struct types")
	methods_on_other_types()
	fmt.Println("--- bank account example")
	bank_example()
}
