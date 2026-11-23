package main

import (
	"fmt"
)

// refernces are not unlike he data types, the are pointers while point to the underlying data structure that work in the background, there are five mani references, pointers, slices, functions, channels and maps

// pointers, help us to first get the address of the stored variable or data type and then change it or assign a new value if requred

// pointer_working -> reads and changes i and j THROUGH pointers, without touching i or j directly
// in short: &x -> "give me the address of x", *p -> "go to the address in p and read / write the value there"
func pointer_working() {
	i, j := 42, 2781
	p := &i // this basically points to i in the data structure, it points p to i in the computer
	fmt.Println(*p) // here we are reading i through the pointer
	*p = 21
	fmt.Println((i))
	
	q := &j // so & actually points to the data structure of i while * is the pointer acting to the change or make actions on the j
	*q = *q / 37
	fmt.Println(j)
}
