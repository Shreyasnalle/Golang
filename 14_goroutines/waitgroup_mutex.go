package main

import (
	"fmt"
	"sync"
	"time"
)

// main_exits_early -> starts a go routine and returns immediately without waiting for it
// it is called as the LAST line of main, so main ends right away and the go routine is killed before its print
// (if we added time.Sleep in main after this, the line WOULD print, sleep is not a real way to wait, WaitGroup is)
func main_exits_early() {
	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("you will NOT see this line, nobody waited for me")
	}()
	fmt.Println("returned without waiting")
}

// worker -> pretends to do some work, then tells the WaitGroup it is done
// wg is a pointer, a copied WaitGroup would count separately and never reach 0
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Done() = counter - 1, defer makes sure it runs even if the worker returns early
	time.Sleep(time.Duration(id) * 20 * time.Millisecond)
	fmt.Println("worker", id, "done")
}

// using_waitgroup -> Add(1) before starting each go routine, Done() inside it, Wait() blocks until the counter is 0
func using_waitgroup() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("all workers finished")
}

// race_condition -> 1000 go routines do counter++ at the same time WITHOUT protection
// counter++ is really read -> add -> write, two go routines can read the same old value, so updates get lost
// the result is often less than 1000 (run with go run -race . to let go detect it)
func race_condition() {
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // unsafe on purpose
		}()
	}
	wg.Wait()
	fmt.Println("without mutex counter =", counter, "(may be less than 1000)")
}

// safe_counter -> a struct that protects its map with a mutex, only one go routine can hold the lock at a time
type safe_counter struct {
	mu sync.Mutex
	v  map[string]int
}

// inc -> lock, change, unlock, defer unlock means we can never forget to unlock
func (c *safe_counter) inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

// value -> reading also needs the lock, because another go routine might be writing at that moment
func (c *safe_counter) value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

// using_mutex -> the same 1000 increments, but now through the mutex, always exactly 1000
func using_mutex() {
	c := safe_counter{v: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.inc("visits")
		}()
	}
	wg.Wait()
	fmt.Println("with mutex counter =", c.value("visits"))
}
