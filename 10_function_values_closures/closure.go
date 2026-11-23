package main

// closure is an inner functio that references another funcion the function whcih is present on above, it can access those variables and modify those ourter variables even after the outer function has stopped executing

// adder -> returns a function, and that inner function keeps its OWN sum alive between calls
// every call to adder() creates a fresh sum, so two adders never mix their totals
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
