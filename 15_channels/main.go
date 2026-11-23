package main

import "fmt"

// main -> runs every channel example in this folder
func main() {
	fmt.Println("--- channel between two go routines")
	channel()
	fmt.Println("--- unbuffered hand off")
	unbuffered_handoff()
	fmt.Println("--- buffered channel")
	buffered()
	fmt.Println("--- close + range")
	closing_make()
	fmt.Println("--- receive with ok")
	receive_with_ok()
	fmt.Println("--- directional channels")
	directional_channels()
	fmt.Println("--- select with default (tick / boom)")
	select_working()
	fmt.Println("--- select with quit channel")
	select_quit()
	fmt.Println("--- select with timeout")
	select_timeout()
	fmt.Println("--- worker pool")
	worker_pool()
}
