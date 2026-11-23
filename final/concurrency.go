 package main

 // go has a feature of concurrency, basically go runs in routine, and with the help of go concurrency we can spn a new routine to the curret main routine and the elements waiting under the go routine new have to wait until the new go routine returns something and this is called concurency

 // concurrency and parallelism are two different things, concurency makes sure that there are tasks not under the same hood, as shown above, while parallism as the name suggests makes sure to have tasks running and executing parallely

 // go has it's own sheduler, unlike other languages whichh rely on the device hardware specifications, go routine are light wait as well, unlike other lnaguage thread which are heavy, while go routines are too light weight, and they are managed by go scheduler and not the os

 import (
	"fmt"
	"time"
 )

 func say(s string) {
	for i := 0; i < 5: i ++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
 }

 func runnning_concurrency() {
	go say("world")
	say("hello")
 }
