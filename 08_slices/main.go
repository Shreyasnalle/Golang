package main 

import "fmt"

// main -> runs every slice example in this folder, with a heading for each
func main() {
	fmt.Println("--- slicing an array")
	slices_use()
	fmt.Println("--- slice literals")
	slicing_initiliazation()
	fmt.Println("--- nil slice")
	null_slice()
	fmt.Println("--- make")
	make_using()
	fmt.Println("--- append")
	append_slice()
	fmt.Println("--- append growth (wasted capacity)")
	append_waste_size()
	fmt.Println("--- append with make (no waste)")
	using_make()
	fmt.Println("--- range over a slice")
	range_over_slice()
	fmt.Println("--- slice defaults")
	slice_defaults()
	fmt.Println("--- copy")
	copy_slice()
	fmt.Println("--- append beyond cap breaks sharing")
	append_breaks_sharing()
	fmt.Println("--- 2d slice (tic tac toe)")
	slice_2d_literal()
	fmt.Println("--- 2d slice with make")
	for _, row := range slice_2d_make(3, 4) {
		fmt.Println(row)
	}
	fmt.Println("--- 3d slice")
	slice_3d()
	fmt.Println("--- 2d slice of struct")
	slice_of_struct_2d()
}
