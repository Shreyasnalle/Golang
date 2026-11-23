package main 

import "fmt"

// person -> a struct that implements fmt.Stringer through its String() method
type person struct {
	name string
	age int
}

// String -> this exact name and signature, String() string, is the fmt.Stringer interface, so Println calls it automatically
func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.name, p.age) // so sprintf returns a string variable, not to be printed to the console, while print anf printf give the statments to be printed to the console or in the terminal

	// so stringer is an interface of go which makes sure that when this method is being called, using the String() it returns the string variable not to console to print, that's something which the print and printf does
}

// stringe_working -> Println prints a and b using our String() method, not the default {name age} format
func stringe_working() {
	a := person{"shreyas nalle", 20}
	b := person{"advit nalle", 10}
	fmt.Println(a, b)
	// and because person struct implemented the stringer interface, an object of that struct will do adhere to that stringer interface
}
