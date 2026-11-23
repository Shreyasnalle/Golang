package main

// Go is directory-based, all .go files in the exact same folder MUST share the exact same package name, we cannot mix different package names in one folder. If we want a separate package to import, we must create a subdirectory (subfolder) for it.

// add -> both parameters are int and the return is also int
func add(x int, y int) int {
	return x + y
}
// here the data type of the parameters are declared in the function, along with that the return data type is also mentioned

// swap -> takes two strings and gives them back in the reverse order
func swap(x string, y string) (string, string) {
	return y, x // fixed: earlier this was return x, y which gave back the same order, so nothing was actually swapped
}
// here the data type present after the parameter declaration is the data type for the RETURN values

// split -> named return values, x and y are born in the signature itself
func split(sum int) (x int, y int) {
	x = sum + 2
	y = sum - 2
	return
}
// here we have declared the varibale which in itself is the return of this function, so we don't use : here

// splitting -> same as split but with plain return types, x and y are declared inside using :=
func splitting(sum int) (int, int) {
	x := sum + 2
	y := sum - 2
	return x, y
}
// here we are just specifing the data type of the return object, not declaring the variable even though we are return x, y, we are declaring the x and y inside the function

var name string
var age int = 20
// so if we want to declare any variable outside funtion, then we use this format, we don't use :, : is only used inside the function
