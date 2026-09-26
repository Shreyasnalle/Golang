package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(10, 15))
}

// go is a statistically typed language which means the data types are checked at the complie time i.e before the code runs, this is more safer, reliable and faster
// while python is dynamically typed language which means the data types are checked at the runtime i.e while the code is running, this is not safe at all, neither that much reliable and slower