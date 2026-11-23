package main

import "fmt"

// main -> runs every map example in this folder
func main() {
	fmt.Println("--- declare with var, initialise with make")
	mapping_use()
	fmt.Println("--- declared + initialised together")
	mapping_in_declaration()
	fmt.Println("--- insert, update, delete, comma ok")
	other_operation_mapping()
	fmt.Println("--- short map literal + len")
	map_literal_short()
	fmt.Println("--- iterating (random vs sorted)")
	iterate_map()
	fmt.Println("--- comma ok inside if")
	comma_ok_in_if()
	fmt.Println("--- nil map")
	nil_map()
	fmt.Println("--- word count")
	fmt.Println(word_count("go is fun and go is fast"))
	fmt.Println("--- map of slices")
	map_of_slices()
	fmt.Println("--- map vs struct")
	map_vs_struct()
}
