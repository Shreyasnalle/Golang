package main

import (
	"fmt"
	"math"
	"runtime"
)

// pow -> returns x^n if it is smaller than lim, otherwise prints sorry and returns lim
// x, n, lim float64 -> when params share a type we write the type only once at the end
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("sorry, %g >= %g, ", v, lim) // v is still usable here in the else block, it dies after the if-else ends
	}
	return lim
}
// this is how conditions are declared in go, first we declare and assign a varaiable then we put ; later we write a condition after that, also the else should be immediately after the } of if condition

var agee int = 10
// here it shows error for age because it is already declared somewhere under the package name MAIN in this current same directory

// switchh -> prints which os this program is running on
func switchh() {
	fmt.Print("Go runs on ") // no ln after print so that the next print can be appended immediately after this print line
	// ; os tells the switch to evaluate and match against the value stored in os
	// runtime.GOOS get's the os in which we are working on
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X")
	case "linux" :
		fmt.Println("Linux")
	default :
		fmt.Printf("%s.\n", os)
	}
}
// so we can observe here that go does not require break statemenet inside every case, it has automatic breaking and that's cool

// switch can be used in place of if else condition

// main -> runs everything from all the files in this folder, go run . compiles all files of the folder together
func main() {
	fmt.Println(add(2, 3))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(10))
	fmt.Println(splitting(10))
	fmt.Println(name == "", age, agee) // name is "" (zero value), age is from basics.go, agee from this file

	fmt.Println(pow(3, 2, 10)) // 9 < 10 so returns 9
	fmt.Println(pow(3, 3, 20)) // 27 >= 20 so prints sorry and returns 20

	switchh()
	if_else_chain(75)
	switch_as_if_else(15)
	switch_multiple_values("sat")
	switch_fallthrough(1)
}
// in the main package, in same directory, there could only be one single main function, that's it