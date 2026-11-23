# 04 - Defer

`defer` schedules a function call to run **just before the surrounding function returns**, whether it returns normally or because of a `panic`.

---

## 1. Basic Defer

```go
func defer_go() {
    defer fmt.Println("From Shreyas to Golang") // runs LAST
    fmt.Println("Hello, World")                 // runs first
}
// Output:
// Hello, World
// From Shreyas to Golang
```

---

## 2. Defer Is a Stack (LIFO: Last In, First Out)

```go
for i := 0; i < 3; i++ {
    defer fmt.Println(i)   // pushed: 0, 1, 2
}
// popped (printed): 2, 1, 0
```

The **last** defer written is the **first** one to run.

---

## 3. Arguments Are Evaluated Immediately

```go
x := 10
defer fmt.Println("deferred x =", x)  // x is read NOW → 10
x = 20
// prints: deferred x = 10
```

The **call** is delayed, but its **arguments** are computed when the `defer` line runs.

---

## 4. Why Defer? Cleanup

```go
fmt.Println("db: connection opened")
defer fmt.Println("db: connection closed")  // written right next to "open"
```

* Write the cleanup (close the file, DB connection or payment gateway) **right next to** the open.
* It runs no matter how many `return` statements the function has.
* Real code: `f, _ := os.Open(name); defer f.Close()`

---

## 5. Defer Runs Even on `panic`

```go
func defer_runs_on_panic() {
    defer func() {
        if r := recover(); r != nil {      // recover() catches the panic
            fmt.Println("recovered:", r)
        }
    }()
    defer fmt.Println("gateway closed")   // still runs!
    panic("something went wrong")
}
```

* `panic` = the program crashes at that point.
* Every pending defer still runs while the panic unwinds.
* `recover()` only works **inside a deferred function**. It stops the panic, and the caller continues normally.

---

## 6. Same Package = No Imports Needed

`main.go` calls `defer_go()` from `defer.go` **without importing anything**. All files in one folder are one package, and they share functions, types and package-level variables.

| Command | What gets compiled |
| :--- | :--- |
| `go run main.go` | **only** `main.go`, so it fails here with `undefined: defer_go` |
| `go run .` | **every** `.go` file in the folder |

---

## 7. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `defer.go` | basic defer, LIFO stack, argument evaluation, cleanup pattern, defer + `panic`/`recover` |
| `main.go` | calls every example; notes on same-package access and `go run .` vs `go run main.go` |

## 8. How to Run

```bash
go run ./04_defer
```

## 9. Revision Checklist

- [ ] When exactly does a deferred call run?
- [ ] Three defers in a row: which order do they run in?
- [ ] `x := 1; defer fmt.Println(x); x = 2`: what is printed?
- [ ] Does defer run if the function panics?
- [ ] Where must `recover()` be called for it to work?
