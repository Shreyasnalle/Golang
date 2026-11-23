package main

import (
	"fmt"
	"time"
)

// unbuffered_handoff -> make(chan string) with no size is unbuffered
// the send in the go routine blocks until main is ready to receive, so "sent" can only print after the receive happened
func unbuffered_handoff() {
	ch := make(chan string)
	go func() {
		fmt.Println("sender: about to send (will block until someone receives)")
		ch <- "ping"
		fmt.Println("sender: sent")
	}()
	time.Sleep(100 * time.Millisecond) // the sender is stuck at ch <- during this sleep
	fmt.Println("receiver: got", <-ch)
	time.Sleep(10 * time.Millisecond) // give the sender a moment to print "sent"
}

// receive_with_ok -> v, ok := <-ch, ok is false only when the channel is closed AND empty
// this is how a receiver can tell "a real 0 was sent" apart from "channel closed, you got the zero value"
func receive_with_ok() {
	ch := make(chan int, 2)
	ch <- 0
	close(ch)
	v, ok := <-ch
	fmt.Println("v =", v, "ok =", ok) // 0 true -> a real value, still in the buffer after close
	v, ok = <-ch
	fmt.Println("v =", v, "ok =", ok) // 0 false -> closed and empty
}

// producer -> chan<- int means this function can ONLY SEND on out, receiving would be a compile error
func producer(out chan<- int) {
	for i := 1; i <= 3; i++ {
		out <- i
	}
	close(out)
}

// consumer -> <-chan int means this function can ONLY RECEIVE from in
func consumer(in <-chan int, done chan<- bool) {
	for v := range in {
		fmt.Println("consumed", v)
	}
	done <- true
}

// directional_channels -> a normal chan int is passed to both, go narrows it to send only / receive only automatically
// the arrow shows the direction of the data, chan<- data goes INTO the channel, <-chan data comes OUT of it
func directional_channels() {
	ch := make(chan int)
	done := make(chan bool)
	go producer(ch)
	go consumer(ch, done)
	<-done // wait until the consumer says it finished, a channel used like a WaitGroup
}

// fibonacci_select -> keeps sending fibonacci numbers on c until something arrives on quit
// select picks whichever case is ready, sending to c when the receiver wants a value, or returning when quit is signalled
func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit received")
			return
		}
	}
}

// select_quit -> the receiver takes 10 numbers and then tells the generator to stop
func select_quit() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(<-c, " ")
		}
		fmt.Println()
		quit <- 0
	}()
	fibonacci_select(c, quit)
}

// select_timeout -> wait for a result, but not forever, time.After acts as a timer
func select_timeout() {
	slow := make(chan string)
	go func() {
		time.Sleep(300 * time.Millisecond)
		slow <- "slow result"
	}()
	select {
	case res := <-slow:
		fmt.Println(res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout! gave up waiting after 100ms")
	}
}

// pool_worker -> takes jobs from the jobs channel until it is closed, sends job*job to results
func pool_worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		time.Sleep(20 * time.Millisecond) // pretend work
		results <- fmt.Sprintf("worker %d squared %d = %d", id, j, j*j)
	}
}

// worker_pool -> 3 go routines share 6 jobs, this combines everything, buffered channels, directions, close and range
func worker_pool() {
	jobs := make(chan int, 6)
	results := make(chan string, 6)
	for w := 1; w <= 3; w++ {
		go pool_worker(w, jobs, results)
	}
	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs) // no more jobs, the workers' range loops end after the last one
	for r := 1; r <= 6; r++ {
		fmt.Println(<-results) // exactly 6 receives, one per job
	}
}

// deadlock note -> if every go routine is blocked (e.g. main sends on an unbuffered channel that nobody receives from)
// go crashes with: fatal error: all goroutines are asleep - deadlock!
// example (not called, it would crash): ch := make(chan int); ch <- 1
