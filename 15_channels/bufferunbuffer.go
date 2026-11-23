package main

// so there are two types of channels, first is the uffered and the second is the unffered, the uffered channel has the concept of quene, whicha basically means that their is a fixed size to this channel, after which no more data can be send over it until that is beig received by the go routine reciever, while the unbuffered channel dosen't have any quene and it keeps on sending to the go routine
// (more exactly: an unbuffered channel has NO storage, so every send BLOCKS until some go routine receives it, the sender and receiver meet hand to hand, see unbuffered_handoff() in channels_more.go)

import "fmt"

// this is a buffered channel
// buffered() -> capacity 2, so two sends work without any receiver, a 3rd send here would block forever (deadlock)
func buffered() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// so how does the receiver know that the sender go routine is done with sendting the values or is their anything mpre to send about, this is done by close(), which is an inbuild go function, it is the same as thaat of the commit() function in the db, which sends the whole execution at once, in the same way, we can close the sending from the sender routine to make sure that the reciever understands that the sending is being done
// (small difference from db commit: close() does not send anything, it just marks "no more values will come", values already in the buffer can still be received, and only the SENDER should close, sending on a closed channel panics)

// we can also, iterate over the items being quened over in the channel, and channel is a refernce, it's on of the references of go, so we can use make function over it and thus set a fixed size of the quene of the channel thus this is really helpful

// fibonacci -> sends the first n fibonacci numbers into c, then closes c so the range loop in the receiver knows to stop
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i ++ {
		c <- x
		x, y = y, x+y // fixed: was x*y, which gave 0 1 0 0 0..., fibonacci adds the previous two numbers
	}
	close(c)
}

// closing_make -> range over a channel keeps receiving until the channel is closed AND empty
func closing_make() {
	c := make(chan int, 30)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
