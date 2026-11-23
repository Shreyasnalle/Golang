package main

import "fmt"

// main -> runs every array example in this folder
func main() {
	fmt.Println("--- basic array")
	array_use()
	fmt.Println("package level array natural:", natural)
	fmt.Println("--- looping over an array")
	array_loop()
	fmt.Println("--- [...] length inference")
	array_length_infer()
	fmt.Println("--- arrays are copied on assignment")
	array_is_a_value()
	fmt.Println("--- 2d array")
	array_2d()
}
