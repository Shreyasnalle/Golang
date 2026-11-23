package main

import "fmt"

// main -> runs every struct example in this folder
func main() {
	fmt.Println("--- basic struct")
	struct_function()
	fmt.Println("--- pointer to struct")
	struct_references()
	fmt.Println("--- struct literals (v1, v2, v3, p from struct.go)")
	fmt.Println(v1, v2, v3, p) // p prints as &{1 2}, the & tells it's a pointer to a struct
	fmt.Println("--- nested struct")
	nested_struct()
	fmt.Println("--- anonymous struct")
	anonymous_struct()
	fmt.Println("--- comparing structs")
	compare_structs()
	fmt.Println("--- struct <-> json")
	struct_and_json()
}
