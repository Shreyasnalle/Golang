package main

// Go is directory-based, all .go files in the exact same folder MUST share the exact same package name, we cannot mix different package names in one folder. If we want a separate package to import, we must create a subdirectory (subfolder) for it.

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}
// here the data type of the parameters are declared in the function, along with that the return data type is also mentioned

func swap(x string, y string) (string, string) {
	return x, y
}
// here the data type present after the parameter declaration is the data type for the RETURN values

func split(sum int) (x int, y int) {
	x = sum + 2
	y = sum - 2
	return
}
// here we have declared the varibale which in itself is the return of this function, so we don't use : here

func splitting(sum int) (int, int) {
	x := sum + 2
	y := sum - 2
	return x, y
}
// here we are just specifing the data type of the return object, not declaring the variable even though we are return x, y, we are declaring the x and y inside the function

var name string
var age int = 20
// so if we want to declare any variable outside funtion, then we use this format, we don't use :, : is only used inside the function

func main(){
	fmt.Println("hello, golang")
}
