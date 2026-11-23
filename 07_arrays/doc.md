# 07 - Arrays

An **array** is a **fixed-size** sequence of elements of the **same type**.

---

## 1. Declaring

```go
var a [2]string          // 2 strings, both "" at the start
a[0] = "Hey,"
a[1] = "I am Shreyas"

var natural = [6]int{1, 2, 3, 4, 5, 6}   // declare + initialise
colors := [...]string{"red", "green"}    // [...] → compiler counts the length
```

* **Outside a function:** `var` and the value on the same line (`:=` not allowed).
* **Inside a function:** declare first and assign later, or use `:=`.

---

## 2. The Size Is Part of the Type

* `[3]int` and `[4]int` are **different types**.
* You **cannot append** to an array or grow it. Accessing past the size is an error.
* You **can** change elements: `a[0] = "Hi"`.
* Use an array only when the size is truly fixed; otherwise use a **slice** (`08_slices`).

---

## 3. Looping

```go
for i := 0; i < len(natural); i++ { natural[i] }  // index-based
for _, value := range natural { sum += value }    // range
```

---

## 4. Arrays Are Values (Copied!)

```go
original := [3]int{1, 2, 3}
copied := original      // the WHOLE array is copied
copied[0] = 100
// original: [1 2 3]   copied: [100 2 3]
```

Slices behave differently: slices of the same array share memory.

---

## 5. 2D Arrays

```go
var grid [2][3]int   // 2 rows × 3 columns
grid[0][1] = 5
```

---

## 6. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `arrays.go` | declaring, assigning by index, fixed size rules, package-level array |
| `arrays_more.go` | looping, `[...]`, arrays are copied on assignment, 2D arrays |
| `main.go` | runs every example |

## 7. How to Run

```bash
go run ./07_arrays
```

## 8. Revision Checklist

- [ ] Can you append to an array?
- [ ] Are `[3]int` and `[4]int` the same type?
- [ ] After `b := a; b[0] = 9`, does `a` change?
- [ ] What does `[...]` do?
