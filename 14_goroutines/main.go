package main

import "fmt"

// main -> runs every goroutine example in this folder
// main_exits_early is kept LAST on purpose, main returns right after it, which kills its go routine before it can print
func main() {
	fmt.Println("--- go routine running next to the current one")
	runnning_concurrency()
	fmt.Println("--- sync.WaitGroup")
	using_waitgroup()
	fmt.Println("--- race condition")
	race_condition()
	fmt.Println("--- sync.Mutex")
	using_mutex()
	fmt.Println("--- nobody waits for a go routine by default (main ends now)")
	main_exits_early()
}
