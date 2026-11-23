# 05 - Pointers & References

A **pointer** holds the **memory address** of a value. Through that address we can read or change the original value.

---

## 1. Reference Types in Go

Pointers, slices, maps, channels and functions are often called **reference types**: they refer to underlying data instead of holding it directly. Pointers are the most explicit of these.

---

## 2. `&` and `*`

| Symbol | Name | Meaning | Example |
| :--- | :--- | :--- | :--- |
| `&x` | address-of | "give me the address of `x`" | `p := &i` |
| `*T` | pointer type | "pointer to a `T`" | `var p *int` |
| `*p` | dereference | "go to the address in `p` and read/write there" | `*p = 21` |

```go
i, j := 42, 2781
p := &i         // p points to i
fmt.Println(*p) // 42, read i through p
*p = 21         // write i through p
fmt.Println(i)  // 21

q := &j
*q = *q / 37    // j = 2781 / 37 = 75
```

---

## 3. Why Pointers? Go Passes Copies

```go
func add_ten_value(n int)    { n = n + 10 }    // changes a COPY
func add_ten_pointer(n *int) { *n = *n + 10 }  // changes the ORIGINAL

score := 50
add_ten_value(score)     // score still 50
add_ten_pointer(&score)  // score now 60
```

**Rule:** a function gets a copy of every argument. To let it change your variable, pass the address `&x`.

---

## 4. Nil Pointers

* Zero value of a pointer = `nil` (points to nothing).
* `*p` on a `nil` pointer → **panic: nil pointer dereference**.
* Always check `if p != nil { ... }` when a pointer might be nil.

---

## 5. `new(T)`

```go
p := new(int)  // allocates a zero int, returns *int
*p = 99
```
Same as `var x int; p := &x`.

---

## 6. No Pointer Arithmetic

Unlike C, Go does **not** allow `p++` to move to the next memory address. This keeps Go memory safe.

---

## 7. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `references.go` | reference types overview, `&` / `*`, reading and writing through pointers |
| `pointers_more.go` | pass by value vs pointer, printing addresses `%p`, nil pointer, `new()` |
| `main.go` | runs every example |

> Pointers to **structs** (`p := &v; p.y = 75`) are covered in `06_structs`.

## 8. How to Run

```bash
go run ./05_pointers
```

## 9. Revision Checklist

- [ ] What is the difference between `&x` and `*p`?
- [ ] What is the type of `&i` when `i` is an `int`?
- [ ] Why didn't `add_ten_value` change `score`?
- [ ] What happens if you dereference a `nil` pointer?
- [ ] What does `new(int)` return?
