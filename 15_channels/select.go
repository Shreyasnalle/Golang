package main

import (
	"fmt"
	"time"
)

// so only when the channel receive somehing it prints the tick or boom and if not then shows ... 
// select_working -> select waits on many channel operations at once and runs whichever is ready first
// time.Tick gives a channel that receives every 100ms, time.After gives a channel that receives once after 500ms
// default runs when NO channel is ready, so the loop never blocks and prints the dots in between
func select_working() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
