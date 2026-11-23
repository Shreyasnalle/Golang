package main

// we can perform append operations of the slices, and there is an internal working of this, like how does the go append over the slices, let's understand it, it's important

// the append function has two parameters, first which takes the slice over which the append is to be perfoemd and second tha parameter which is to be added into the slice, at one time we can add more than one elements as well in the slice

// the length depends on what number of elements are present in the slice, but the capacity does work like that, when the length get's equal to the capacity the capacity get's increased by 2x, so until the 2xth capacity get's equal to the length the appending can continue to be done, after 256 capacity the 2x never happens rather it get increased by 1.25x

import "fmt" 

// printSlices -> prints length, capacity and the slice itself
func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // fixed: was fmt.Println, but Println does not understand %d, only Printf does
}

// append_slice -> appends one value, then another, then three at once, printing len/cap every time
func append_slice() {
	var s []int 
	printSlices(s)
	s = append(s, 0)
	printSlices(s)
	s = append(s, 1)
	printSlices(s)
	s = append(s, 2, 3, 4)
	printSlices(s)
}

// here after the capacity get's 1, it will e increased to 2, then from 2 to 4 adn then from 4 to 8, but we are appending the numbers max until 5 only, so the rest 3 spacing for teh elements which will never get appended will never be utilized, thus this wastes the sapce and we would not want that to happen, for that we use make function

// append_waste_size -> appends 0 to 4 one by one, watch cap jump 1 -> 2 -> 4 -> 8, and 3 slots of 8 are never used
func append_waste_size() {
	var s[] int
	printSlices(s)
	for i := 0; i < 5; i ++ {
		s = append(s, i)
		printSlices(s)
	}
}

// using_make -> we already know we need 5 elements, so we reserve cap 5 at the start, no re-growing, no waste
func using_make() {
	s := make([]int, 0, 5) // so this will make sure that capacity is always equal to 5 even if we use append here, make function tool has this format of parameters (data structure, length, variable)
	// note: the 3rd parameter is the capacity, and cap stays 5 only while we append 5 or less, a 6th append would grow it again
	printSlices(s)
	for i := 0; i < 5; i ++ {
		s = append(s, i)
		printSlices(s)
	}
}
