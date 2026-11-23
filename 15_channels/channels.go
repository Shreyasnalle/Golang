package main 

// channel is a way of communication in between different go routines, so we can acually orchestrate, look over and make different go routines have a coordination and communication in between them with go channel

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send the sum to c, which is the channel
}
// here the c chan int, int is the type of the data which is being channeled through or commmunicated via the channel 

func channel() {
	s := []int{1, 2, 3, 4, 5, 6}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x*y)
}