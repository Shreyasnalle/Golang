package main

import "fmt"

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