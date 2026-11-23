package main

import "fmt"

// range_over_slice -> range gives index and value for every element
func range_over_slice() {
	pow := []int{1, 2, 4, 8, 16}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i := range pow { // only the index, value skipped
		pow[i] = pow[i] * 10
	}
	fmt.Println("after multiplying by 10:", pow)
}

// slice_defaults -> in s[low:high] low defaults to 0 and high defaults to len(s)
func slice_defaults() {
	s := []int{10, 20, 30, 40, 50}
	fmt.Println(s[:2], s[2:], s[:]) // [10 20] [30 40 50] [10 20 30 40 50]
}

// copy_slice -> copy(dst, src) copies elements into a NEW backing array, so the two slices are independent
// compare this with b := a, there both point to the same array and changing one changes the other
func copy_slice() {
	a := []int{1, 2, 3}
	shared := a // same underlying array
	independent := make([]int, len(a))
	n := copy(independent, a) // returns how many elements were copied
	shared[0] = 99
	fmt.Println("a:", a, "| shared:", shared, "| independent:", independent, "| copied:", n)
}

// append_breaks_sharing -> when append needs more than cap, go makes a NEW bigger array
// after that the new slice and the old one no longer share memory
func append_breaks_sharing() {
	base := make([]int, 3, 3) // len 3, cap 3, full
	grown := append(base, 4)  // cap exceeded -> new array allocated
	grown[0] = 100
	fmt.Println("base:", base, "| grown:", grown) // base still starts with 0
}
