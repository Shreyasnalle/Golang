package main

import "fmt"

// main -> runs the non generic version, the generic version and the extra generic examples
func main() {
	fmt.Println("--- without generics (two functions)")
	generics_running()
	fmt.Println("--- with generics (one function)")
	generic()
	fmt.Println("--- constraints, multiple type params, generic types")
	generics_more()
}
