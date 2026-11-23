package main

import "fmt"

// array_use -> declares an array of 2 strings, fills it index by index and prints it
func array_use() {
	var a [2]string
	a[0] = "Hey,"
	a[1] = "I am Shreyas"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

// so arrays are basically of fixed size in golang, so we can't append anything in this, such that the maximum size of the array get's breached
// so we use array only when we know that we will have this strict parameters in that array, or else we use slice, we will learn this next

// so the array size is the part of the array type, we cannot access anything beyound that array size, yeah but we can edit things

var natural = [6]int{1, 2, 3, 4, 5, 6}

// so outside function, the var should be declared and assigned on the same line, but in function it can be done on seperate line
