 package main

 // go has a feature of concurrency, basically go runs in routine, and with the help of go concurrency we can spn a new routine to the curret main routine and the elements waiting under the go routine new have to wait until the new go routine returns something and this is called concurency
 // (note: the caller does NOT wait for the new go routine, go say("world") returns immediately and both run side by side, and when main() returns ALL go routines are killed, finished or not, to really wait we use sync.WaitGroup -> see waitgroup_mutex.go)

 // concurrency and parallelism are two different things, concurency makes sure that there are tasks not under the same hood, as shown above, while parallism as the name suggests makes sure to have tasks running and executing parallely

 // go has it's own sheduler, unlike other languages whichh rely on the device hardware specifications, go routine are light wait as well, unlike other lnaguage thread which are heavy, while go routines are too light weight, and they are managed by go scheduler and not the os

 import (
	"fmt"
	"time"
 )

 // say -> prints s five times with a 100ms sleep in between, the sleep gives the other go routine a chance to run
 func say(s string) {
	for i := 0; i < 5; i ++ { // fixed: was i < 5: i ++ (a : instead of ;), which did not compile
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
 }

 // runnning_concurrency -> "world" runs in a NEW go routine, "hello" runs in the current one, so the output is mixed
 func runnning_concurrency() {
	go say("world")
	say("hello")
 }
