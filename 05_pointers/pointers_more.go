package main

import "fmt"

// add_ten_value -> gets a COPY of the number, so changing n here does not change the caller's variable
func add_ten_value(n int) {
	n = n + 10
}

// add_ten_pointer -> gets the ADDRESS of the number, so *n = ... changes the caller's original variable
func add_ten_pointer(n *int) {
	*n = *n + 10
}

// pass_by_value_vs_pointer -> this is THE main reason pointers exist, go always passes a copy to a function
// if we want the function to change our variable, we pass its address (&score) instead of the value
func pass_by_value_vs_pointer() {
	score := 50
	add_ten_value(score)
	fmt.Println("after add_ten_value  :", score) // still 50, only the copy was changed
	add_ten_pointer(&score)
	fmt.Println("after add_ten_pointer:", score) // 60, the original was changed
}

// print_addresses -> p holds the address of i, so printing p shows the same address as &i
func print_addresses() {
	i := 7
	p := &i
	fmt.Printf("i = %v, &i = %p\n", i, &i)
	fmt.Printf("p = %p, *p = %v, type of p = %T\n", p, *p, p) // type is *int, "pointer to int"
}

// nil_pointer -> the zero value of a pointer is nil, it points to nothing
// reading *p when p is nil crashes the program (panic: nil pointer dereference), so always check for nil first
func nil_pointer() {
	var p *int
	fmt.Println("p is", p, "| p == nil ?", p == nil)
	if p != nil {
		fmt.Println(*p) // safe, only runs if p actually points somewhere
	}
}

// using_new -> new(T) creates a zero value of type T in memory and gives back a pointer to it
// so new(int) is the same as -> var x int; p := &x
func using_new() {
	p := new(int)
	fmt.Println("*p starts as", *p) // 0, the zero value of int
	*p = 99
	fmt.Println("*p is now", *p)
}

// note: go has pointers but NO pointer arithmetic (p++ to move to the next memory block like in C is not allowed), this keeps go memory safe
