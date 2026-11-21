package main

// there is a difference in between array and struct, array store elements with same data types while struct store elements of difference data types, it's like list in python

import (
	"fmt"
)

type vertex struct {
	x int
	y int
}

func struct_function() {
	v := vertex{1, 2}
	fmt.Println(v)
	v.x = 10
	fmt.Println(v)
}

// we can also use struct with references, and more specifically with the pointer references

func struct_references() {
	v:= vertex{1, 5}
	p := &v
	p.y = 75
	fmt.Println(v)
}

// o struct is something which is really important in golang, it's used everywhere, as it has the ability to store without the condition of having elements with same data type, it can store eleements with any data type, so that's really good, in real example,it's being used for validation, like unlike js, go dosent have the compatibility for json with it, so it uses struct to do teh validation and store the json payloads here in the struct, and then the further operations are done over this stored struct

var (
	v1 = vertex {1, 2} // has type vertex
	v2 = vertex {x: 1} // here the x has being assigned the value, but y is not assigned so golang by default passes it 0
	v3 = vertex {} // by deafult assigngs 0
	p = &vertex {1, 2} // has the location to this data structure
)

// we can also use the struct in another package, but for this we have to use uppercase while declaring the parameters of that struct, the first letter should be in uppercase. If it's in small then that struct can be used only in the curent package