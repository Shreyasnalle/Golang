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

func main() {
	fmt.Println(add(10, 15))
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
	fmt.Println(split(10))
	fmt.Println(split(10))
}
