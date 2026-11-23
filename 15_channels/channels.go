package main 

// channel is a way of communication in between different go routines, so we can acually orchestrate, look over and make different go routines have a coordination and communication in between them with go channel

import "fmt"

// sum -> adds up the slice and SENDS the result into the channel c instead of returning it
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send the sum to c, which is the channel
}
// here the c chan int, int is the type of the data which is being channeled through or commmunicated via the channel 

// channel -> splits the slice in two halves, sums each half in its own go routine, then receives both results
// x, y := <-c, <-c blocks until both go routines have sent, so the channel also does the "waiting" for us
// the order of x and y is not fixed, whichever go routine finishes first is received first
func channel() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x*y)
}
