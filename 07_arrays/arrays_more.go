package main

import "fmt"

// array_loop -> going over an array with a normal for and with range
// len(arr) gives the size of the array
func array_loop() {
	for i := 0; i < len(natural); i++ {
		fmt.Print(natural[i], " ")
	}
	fmt.Println("<- normal for using len()")

	sum := 0
	for _, value := range natural { // _ because we don't need the index here
		sum += value
	}
	fmt.Println("sum of natural using range:", sum)
}

// array_length_infer -> [...] lets the compiler count the elements for us, the size is still fixed after this
func array_length_infer() {
	colors := [...]string{"red", "green", "blue"}
	fmt.Printf("%v has len %d and type %T\n", colors, len(colors), colors) // type is [3]string
}

// array_is_a_value -> assigning an array to another variable COPIES the whole array
// so changing the copy does not change the original (slices behave differently, see 08_slices)
func array_is_a_value() {
	original := [3]int{1, 2, 3}
	copied := original
	copied[0] = 100
	fmt.Println("original:", original, "| copied:", copied)

	// and because the size is part of the type, [3]int and [4]int are two DIFFERENT types
	// var x [4]int = original -> compile error, cannot use [3]int as [4]int
}

// array_2d -> an array of arrays, like a small matrix / grid
func array_2d() {
	var grid [2][3]int // 2 rows, 3 columns, all 0 at the start
	grid[0][1] = 5
	grid[1][2] = 9
	for _, row := range grid {
		fmt.Println(row)
	}
}
