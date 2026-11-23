package main

// so slices in golang is like array slicing or list comprehension

import "fmt"

// so slices has three things, one is the length which is the totaal length of that slice, as the name suggets, the second is capacity which is the total capacity from the starting to the end of that slice and pointer is bascially the pointer in that slice

// how slicing can be done on the array

func slices_use() {
	names := [9]string {
		"Shreyas",
		"Tech",
		"Finance",
		"Games",
	}
	fmt.Println(names)
	a := names[0 : 2]
	b := names [1 : 3]
	fmt.Println(a, b)
	b[0] = "Advit"
	fmt.Println(a, b)
	fmt.Println(names)
}

// initialising slices in the code

// so here as we can see, the slicing is the same as that opf the array but the slicing nver has strict size unlike the array, also we have a slice of struct, that is a slice having elements of any data types

func slicing_initiliazation() {
	q := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(q)

	r := []bool{true, false, true, false, true, true}
	fmt.Println(r)

	s := []struct {
		i  int
		b bool
	} {
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(s)
}

func null_slice(){
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nillllll, let's go")
	}
} 

// make is a inbuild golang tool which is used to intilize and assign objects to some specific refernces type, one of which is the slices 

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func make_using() {
	a := make([]int, 5) // make tool has this format: type, length, optional capacity
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2] // performing slicing over the objets created by the making tool
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

// like we can create 2d, 3d and blah blah dimensional arrays in python, in the same way, we can also create 2d, 3d slices in golang as well, i will not code this part, it can be done by ai, i know th concept and i think that's enough, obv i can use make function there too, use strcut over that slice to store elements with varying data types