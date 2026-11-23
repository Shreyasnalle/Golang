# 15 - Channels & Select

A **channel** is a typed pipe that goroutines use to **communicate** and **synchronise**.

> "Don't communicate by sharing memory; share memory by communicating."

---

## 1. Basics

```go
c := make(chan int)   // a channel that carries ints (channels are reference types, made with make)
c <- 5                // SEND 5 into c
v := <-c              // RECEIVE from c
```

```go
func sum(s []int, c chan int) { ...; c <- sum }

go sum(s[:len(s)/2], c)
go sum(s[len(s)/2:], c)
x, y := <-c, <-c        // waits for both results (order not guaranteed)
```

---

## 2. Unbuffered vs Buffered

| | Unbuffered `make(chan T)` | Buffered `make(chan T, n)` |
| :--- | :--- | :--- |
| Storage | none | a queue of `n` slots |
| Send blocks when... | **always**, until a receiver takes it (hand-off) | the buffer is **full** |
| Receive blocks when... | no sender is ready | the buffer is **empty** |
| Use for | synchronising two goroutines | decoupling a fast sender from a slower receiver |

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2      // fine, buffer now full
// ch <- 3   // would block forever → deadlock
```

---

## 3. `close` and `range`

```go
func fibonacci(n int, c chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x+y
    }
    close(c)                 // "no more values are coming"
}

for i := range c { }         // stops when c is closed AND empty
v, ok := <-c                 // ok == false → closed and empty
```

* Only the **sender** closes. Sending on a closed channel **panics**.
* Values already buffered can still be received after `close`.
* You don't *have* to close a channel. Close it only when the receiver needs to know the data is finished (e.g. `range`).

> Fix made: `fibonacci` used `x*y`, which gave `0 1 0 0 0…`. It's now `x+y`.

---

## 4. Directional Channels

```go
func producer(out chan<- int)   // send-only
func consumer(in <-chan int)    // receive-only
```

The compiler stops a function from using the channel in the wrong direction. A normal `chan int` converts to either one automatically.

---

## 5. `select`

`select` waits on **several channel operations** and runs whichever is ready first (picking randomly if several are ready).

```go
select {
case <-tick:                    // receive
case c <- x:                    // send
case <-quit:                    // stop signal
    return
case <-time.After(100 * time.Millisecond): // timeout
default:                        // runs if NOTHING is ready, so select never blocks
}
```

| Pattern | Function |
| :--- | :--- |
| ticker + bomb + `default` | `select_working` |
| generator with a **quit** channel | `fibonacci_select` / `select_quit` |
| **timeout** | `select_timeout` |

---

## 6. Deadlock

If **every** goroutine is blocked, Go crashes:
```
fatal error: all goroutines are asleep - deadlock!
```
Classic cause: `ch := make(chan int); ch <- 1` in `main` with no receiver.

---

## 7. Worker Pool (everything together)

```go
jobs := make(chan int, 6)
results := make(chan string, 6)
for w := 1; w <= 3; w++ { go pool_worker(w, jobs, results) }
for j := 1; j <= 6; j++ { jobs <- j }
close(jobs)                    // workers' range loops finish
for r := 1; r <= 6; r++ { <-results }
```

---

## 8. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `channels.go` | what a channel is, `sum` in two goroutines, send/receive |
| `bufferunbuffer.go` | buffered vs unbuffered, `close`, `range` over a channel, `fibonacci` |
| `select.go` | `select` with `time.Tick`, `time.After` and `default` |
| `channels_more.go` | unbuffered hand-off, `v, ok := <-ch`, directional channels, select + quit, select + timeout, worker pool, deadlock note |
| `main.go` | runs every example |

## 9. How to Run

```bash
go run ./15_channels
```

## 10. Revision Checklist

- [ ] When does a send on an unbuffered channel block? On a buffered one?
- [ ] Who should close a channel? What happens if you send after `close`?
- [ ] How does `for v := range ch` know when to stop?
- [ ] What does `ok` mean in `v, ok := <-ch`?
- [ ] What's the difference between `chan<- int` and `<-chan int`?
- [ ] What does `default` do inside `select`?
- [ ] How do you add a timeout to a channel receive?
- [ ] What causes "all goroutines are asleep - deadlock!"?
