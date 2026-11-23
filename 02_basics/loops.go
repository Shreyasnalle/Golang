package main

import "fmt"

// so go has only ONE loop keyword and that is for, there is no while and no do-while in go
// but the same for can be written in different ways to act like a while loop or an infinite loop

// loop_types shows all the shapes of the for loop
func loop_types() {
	// 1. classic for -> init; condition; post (same as the loop() function in main.go)
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println("<- classic for")

	// 2. while style for -> only the condition, init and post are dropped
	n := 1
	for n < 100 {
		n *= 2
	}
	fmt.Println(n, "<- while style for, keeps doubling until n >= 100")

	// 3. infinite for -> no condition at all, we have to come out with break (or return)
	count := 0
	for {
		count++
		if count == 3 {
			break // break jumps out of the loop
		}
	}
	fmt.Println(count, "<- infinite for, stopped using break")

	// 4. continue -> skips the rest of the current iteration and moves to the next one
	for i := 0; i < 6; i++ {
		if i%2 == 0 {
			continue // skip even numbers
		}
		fmt.Print(i, " ")
	}
	fmt.Println("<- only odd numbers, evens skipped using continue")

	// 5. range -> loops over a collection, gives index and value (we will use it a lot with slices and maps later)
	for index, value := range []string{"go", "is", "fun"} {
		fmt.Print(index, ":", value, " ")
	}
	fmt.Println("<- range gives index and value")
	// if we don't need the index we write _ (blank identifier) -> for _, value := range ...
}
