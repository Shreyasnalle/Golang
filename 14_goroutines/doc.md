# 14 - Goroutines & Concurrency

A **goroutine** is a lightweight thread managed by the **Go runtime scheduler**, not the OS. You start one by putting `go` in front of a function call.

---

## 1. Starting a Goroutine

```go
func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

go say("world")   // runs in a NEW goroutine, this line returns immediately
say("hello")      // runs in the current goroutine
// output: hello / world mixed together
```

> Fix made: the loop was `i < 5: i ++` (`:` instead of `;`), which didn't compile.

---

## 2. Concurrency vs Parallelism

| Concurrency | Parallelism |
| :--- | :--- |
| *Dealing with* many tasks at once by interleaving them | *Doing* many tasks at the exact same instant |
| possible on one CPU core | needs multiple cores |
| about **structure** | about **execution** |

Go gives you concurrency. The scheduler runs goroutines in parallel when cores are available.

---

## 3. Why Goroutines Are Cheap

* They start with a tiny stack (a few KB) that grows when needed. OS threads take ~1MB.
* The **Go scheduler** multiplexes thousands or millions of goroutines onto a few OS threads.
* Creating one is almost free compared to a thread in Java or C++.

---

## 4. Nobody Waits by Default

* `go f()` does **not** block. The caller carries on.
* When `main()` returns, **every goroutine is killed**, finished or not.
* `time.Sleep` is **not** a real way to wait. Use a `WaitGroup` (or channels, `15_channels`).

---

## 5. `sync.WaitGroup`: Wait for Goroutines to Finish

```go
var wg sync.WaitGroup
for i := 1; i <= 3; i++ {
    wg.Add(1)            // +1 BEFORE starting the goroutine
    go worker(i, &wg)    // pass a POINTER
}
wg.Wait()                // blocks until the counter reaches 0

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()      // -1 when the worker finishes
}
```

---

## 6. Race Conditions & `sync.Mutex`

`counter++` is really **read → add → write**. Two goroutines can read the same old value and one update gets lost. In this repo, 1000 unprotected increments gave **964**.

```go
type safe_counter struct {
    mu sync.Mutex
    v  map[string]int
}
func (c *safe_counter) inc(key string) {
    c.mu.Lock()           // only ONE goroutine past this point at a time
    defer c.mu.Unlock()
    c.v[key]++
}
```

Detect races with `go run -race ./14_goroutines` (this compiles in extra checks, so it is slower and uses more memory).

---

## 7. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `concurrency.go` | `go` keyword, concurrency vs parallelism, the Go scheduler, `say` example |
| `waitgroup_mutex.go` | main doesn't wait, `sync.WaitGroup`, a race condition, `sync.Mutex` with `safe_counter` |
| `main.go` | runs every example (`main_exits_early` is last on purpose) |

## 8. How to Run

```bash
go run ./14_goroutines
```

## 9. Revision Checklist

- [ ] What does the `go` keyword do?
- [ ] What happens to running goroutines when `main` returns?
- [ ] Why call `wg.Add(1)` *before* `go worker(...)`?
- [ ] Why must the WaitGroup be passed as a pointer?
- [ ] Why can 1000 `counter++` calls end with less than 1000?
- [ ] Concurrency vs parallelism in one line each?
