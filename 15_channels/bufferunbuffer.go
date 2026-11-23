package main

// so there are two types of channels, first is the uffered and the second is the unffered, the uffered channel has the concept of quene, whicha basically means that their is a fixed size to this channel, after which no more data can be send over it until that is beig received by the go routine reciever, while the unbuffered channel dosen't have any quene and it keeps on sending to the go routine

import "fmt"

// this is a buffered channel
func buffered() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// so how does the receiver know that the sender go routine is done with sendting the values or is their anything mpre to send about, this is done by close(), which is an inbuild go function, it is the same as thaat of the commit() function in the db, which sends the whole execution at once, in the same way, we can close the sending from the sender routine to make sure that the reciever understands that the sending is being done

// we can also, iterate over the items being quened over in the channel, and channel is a refernce, it's on of the references of go, so we can use make function over it and thus set a fixed size of the quene of the channel thus this is really helpful

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i ++ {
		c <- x
		x, y = y, x*y
	}
	close(c)
}

func closing_make() {
	c := make(chan int, 30)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}