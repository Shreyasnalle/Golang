package main

// closure is an inner functio that references another funcion the function whcih is present on above, it can access those variables and modify those ourter variables even after the outer function has stopped executing

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}