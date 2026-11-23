package main

import "fmt"

// as you can see, i have not imported the function from the defer.go file, to run it from here, and that's so beautiful thing of golang, damn man, so in the same package or you can also say in the same directory, there is no need to import other function, classses or anythig from other codes in the same directory or pacakge

// so even the variables in the same package can be used in other files without doing the importing stuff, the variable which are outside aany function, the ones which are inside alredy come with the function import

func main() {
	defer_go()
	struct_function()
	struct_references()
	fmt.Println(v1, v2, v3, p)
}

// go run main.go only runs the main file which is the starting point of the executable code in go, while go run . runs all the files in that directory