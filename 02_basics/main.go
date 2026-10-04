package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// go is a statistically typed language which means the data types are checked at the complie time i.e before the code runs, this is more safer, reliable and faster
// while python is dynamically typed language which means the data types are checked at the runtime i.e while the code is running, this is not safe at all, neither that much reliable and slower

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
// here in the return data casting itself we are declaring x, y thus we have used = here

func splitting(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}
// here we are just specifing the data type return and not decalring the variable, so we use := which means declare and assign

// : is the declaration of the variable while = is it's assignement

var name string 
var agee int = 20
// if we want to declare a variable outside any function then we have to follow this format, we ant declare and assign here, we cant use := outside the function

func variable() {
	namee := "Shreyas"
	var age int = 20
	fmt.Println(namee, age)
}
// := this can only we used inside the function for the variable

func varr() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)
}

func loop() {
	sum := 0
	for i := 0; i < 10; i ++ {
		sum += 1
	}
	fmt.Println("The final sum is:", sum)
}

func main() {
	fmt.Println(add(10, 15))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(10))
	fmt.Println(split(10))
	variable()
	varr()
	v := "shreyas"
	fmt.Printf("type of the v is %T\n", v)
	loop()
}

// % are called as formatting specifiers, some of the most used are, %v for value print, %T type of the value, %+v the whole struct, %s for string without quotes, %q string with quotes