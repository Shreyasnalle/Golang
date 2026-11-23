package main

// we can perform append operations of the slices, and there is an internal working of this, like how does the go append over the slices, let's understand it, it's important

// the append function has two parameters, first which takes the slice over which the append is to be perfoemd and second tha parameter which is to be added into the slice, at one time we can add more than one elements as well in the slice

// the length depends on what number of elements are present in the slice, but the capacity does work like that, when the length get's equal to the capacity the capacity get's increased by 2x, so until the 2xth capacity get's equal to the length the appending can continue to be done, after 256 capacity the 2x never happens rather it get increased by 1.25x

import "fmt" 

func printSlices(s []int) {
	fmt.Println("len=%d cap=%d %v\n", len(s), cap(s), s)
}

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

func append_waste_size() {
	var s[] int
	printSlices(s)
	for i := 0; i < 5; i ++ {
		s = append(s, i)
		printSlices(s)
	}
}

func using_make() {
	s := make([]int, 0, 5) // so this will make sure that capacity is always equal to 5 even if we use append here, make function tool has this format of parameters (data structure, length, variable)
	printSlices(s)
	for i := 0; i < 5; i ++ {
		s = append(s, i)
		printSlices(s)
	}
}