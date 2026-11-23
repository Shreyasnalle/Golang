package main

import "fmt"

// as you can see, i have not imported the function from the defer.go file, to run it from here, and that's so beautiful thing of golang, damn man, so in the same package or you can also say in the same directory, there is no need to import other function, classses or anythig from other codes in the same directory or pacakge

// so even the variables in the same package can be used in other files without doing the importing stuff, the variable which are outside aany function, the ones which are inside alredy come with the function import

// main -> calls every defer example one by one, with a heading for each so the output is easy to read
func main() {
	fmt.Println("--- basic defer")
	defer_go()
	fmt.Println("--- defer is a stack (LIFO)")
	defer_stack()
	fmt.Println("--- defer arguments are evaluated immediately")
	defer_args_evaluated_now()
	fmt.Println("--- real world: closing a connection")
	open_db_connection()
	fmt.Println("--- defer still runs on panic")
	defer_runs_on_panic()
	fmt.Println("main continues normally after the recovered panic")
}

// go run main.go only runs the main file which is the starting point of the executable code in go, while go run . runs all the files in that directory
// note: in this folder go run main.go would FAIL with undefined: defer_go, because defer.go is not included, so always use go run . here
