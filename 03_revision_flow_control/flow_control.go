package main

import "fmt"

// if_else_chain -> the normal if / else if / else ladder
// no () around the condition in go, but {} are always compulsory, even for a single line
func if_else_chain(marks int) {
	if marks >= 90 {
		fmt.Println(marks, "-> grade A")
	} else if marks >= 60 {
		fmt.Println(marks, "-> grade B")
	} else {
		fmt.Println(marks, "-> grade C")
	}
	// the else if / else should be on the same line as the closing } of the previous block, else go puts a ; there and it's a syntax error
}

// switch_as_if_else -> switch with NO condition after it, this is the same as switch true
// so every case is a full condition and the first case which is true runs, it is a cleaner way of writing a long if-else chain
func switch_as_if_else(hour int) {
	switch {
	case hour < 12:
		fmt.Println(hour, "-> Good morning")
	case hour < 17:
		fmt.Println(hour, "-> Good afternoon")
	default:
		fmt.Println(hour, "-> Good evening")
	}
}

// switch_multiple_values -> one case can match many values, seperated by comma
func switch_multiple_values(day string) {
	switch day {
	case "sat", "sun":
		fmt.Println(day, "-> weekend")
	default:
		fmt.Println(day, "-> weekday")
	}
}

// switch_fallthrough -> go breaks automatically after a case, but if we WANT the next case to run too we write fallthrough
// fallthrough runs the next case body WITHOUT checking its condition, so use it carefully, it's rarely needed
func switch_fallthrough(n int) {
	switch n {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two (reached because of fallthrough from case 1)")
	case 3:
		fmt.Println("three, never printed for n = 1 because case 2 has no fallthrough")
	}
}
