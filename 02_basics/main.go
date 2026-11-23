package main

import "fmt"

// add takes two ints and returns their sum, the type is written AFTER the variable name (x int), and the return type is written after the brackets
func add(x int, y int) int {
	return x + y
}

// go is a statistically typed language which means the data types are checked at the complie time i.e before the code runs, this is more safer, reliable and faster
// while python is dynamically typed language which means the data types are checked at the runtime i.e while the code is running, this is not safe at all, neither that much reliable and slower

// swap returns TWO values, go functions can return multiple values, the (string, string) tells the return types
func swap(x string, y string) (string, string) {
	return y, x
}

// split uses NAMED return values, x and y are already declared in the signature so we just assign them and do a naked return
func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
// here in the return data casting itself we are declaring x, y thus we have used = here

// splitting does the same work as split, but with normal (un-named) return values
func splitting(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}
// here we are just specifing the data type return and not decalring the variable, so we use := which means declare and assign

// : is the declaration of the variable while = is it's assignement

var name string
var agee int = 20
// if we want to declare a variable outside any function then we have to follow this format, we cant declare and assign here, we cant use := outside the function

// variable shows both ways of declaring a variable inside a function, := (short) and var (long)
func variable() {
	namee := "Shreyas"
	var age int = 20
	fmt.Println(namee, age)
}
// := this can only we used inside the function for the variable

// varr shows the zero values, variables declared without a value get 0, 0, false and "" (empty string)
func varr() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // %q shows the empty string as "" so we can actually see it
}

// loop is the classic 3 part for loop -> init; condition; post
func loop() {
	sum := 0
	for i := 0; i < 10; i ++ {
		sum += 1
	}
	fmt.Println("The final sum is:", sum)
	// note: sum += 1 adds 1 every time, so this counts the iterations (10), if we did sum += i it would add 0+1+...+9 = 45
}

// main calls every function of this folder one by one, so go run . shows everything
func main() {
	fmt.Println(add(10, 15))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(10))
	fmt.Println(splitting(10)) // same answer as split(10), just written the other way
	fmt.Println("package level vars:", name, agee) // name is "" because it's only declared, agee is 20
	variable()
	varr()
	v := "shreyas"
	fmt.Printf("type of the v is %T\n", v)
	loop()
	loop_types()
	constants_and_types()
	format_specifiers()
}

// % are called as formatting specifiers, some of the most used are, %v for value print, %T type of the value, %+v the whole struct, %s for string without quotes, %q string with quotes
