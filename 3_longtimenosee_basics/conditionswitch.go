package main

import (
	"fmt"
	"math"
	"runtime"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("sorry")
	}
	return lim
}
// this is how conditions are declared in go, first we declare and assign a varaiable then we put ; later we write a condition after that, also the else should be immediately after the } of if condition

var agee int = 10
// here it shows error for age because it is already declared somewhere under the package name MAIN in this current same directory

func switchh() {
	fmt.Print("Go runs on ") // no ln after print so that the next print can be appended immediately after this print line
	// ; os tells the switch to evaluate and match against the value stored in os
	// runtime.GOOS get's the os in which we are working on
	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X")
	case "linux" :
		fmt.Println("Linu")
	default :
		fmt.Printf("%s.\n", os)
	}
}
// so we can observe here that go does not require break statemenet inside every case, it has automatic breaking and that's cool

func main() {
	switchh()
}
// in the main package, in same directory, there could only be one single main function, that's it