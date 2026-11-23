# 08 - Slices

A **slice** is a flexible, growable **view into an array**. In real Go code you use slices far more than arrays.

---

## 1. What a Slice Is Made Of

A slice is a small header with three parts:

| Part | Meaning | Get it with |
| :--- | :--- | :--- |
| **pointer** | where in the underlying array the slice starts | — |
| **length** | how many elements the slice currently has | `len(s)` |
| **capacity** | how many elements fit from the start of the slice to the end of the underlying array | `cap(s)` |

---

## 2. Slicing an Array: Slices Share Memory

```go
names := [9]string{"Shreyas", "Tech", "Finance", "Games"}
a := names[0:2]   // [Shreyas Tech]
b := names[1:3]   // [Tech Finance]
b[0] = "Advit"    // changes names[1], so a[1] changes too!
```

`s[low:high]` includes `low` and excludes `high`. Defaults: `low = 0`, `high = len(s)`, so `s[:2]`, `s[2:]` and `s[:]` all work.

---

## 3. Creating Slices

```go
q := []int{1, 2, 3}                     // literal (no size inside [])
var s []int                             // nil slice: len 0, cap 0, s == nil
a := make([]int, 5)                     // len 5, cap 5, all zeros
b := make([]int, 0, 5)                  // len 0, cap 5
s := []struct{ i int; b bool }{{1, true}} // slice of anonymous struct
```

`make(type, length, capacity)`: the third argument is the **capacity** and is optional.

---

## 4. `append` and Growth

```go
s = append(s, 0)          // ALWAYS reassign: s = append(s, ...)
s = append(s, 2, 3, 4)    // many values at once
```

When `len == cap` and you append, Go allocates a **new, bigger array** and copies everything:

```
cap: 0 → 1 → 2 → 4 → 8 → ...   (doubling while small)
after ~256: grows by ~1.25x
```

* Appending **several values at once** can give an odd capacity. In this repo, `append(s, 2, 3, 4)` on cap 2 gave **cap 6**.
* If you already know the size, **pre-allocate** with `make([]int, 0, n)` to avoid repeated re-growing and wasted capacity.
* After a re-grow, the new slice **no longer shares** memory with the old one:

```go
base := make([]int, 3, 3)
grown := append(base, 4)  // new array
grown[0] = 100            // base is NOT affected
```

---

## 5. `range` and `copy`

```go
for i, v := range pow { }    // index + value
for i := range pow { }       // index only

independent := make([]int, len(a))
copy(independent, a)          // real copy, no shared memory
shared := a                   // NOT a copy, same array
```

---

## 6. Multi-dimensional Slices

```go
board := [][]string{          // 2D literal
    {"_", "_", "_"},
    {"_", "_", "_"},
}
board[0][0] = "X"

grid := make([][]int, rows)   // 2D with make:
for i := range grid {
    grid[i] = make([]int, cols) // every row needs its own make
}

cube := [][][]int{{{1, 2}}}   // 3D: just one more []
```

Rows can have **different lengths** (jagged), which a 2D array can't.

---

## 7. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `slices.go` | slicing an array, shared memory, slice literals, slice of struct, nil slice, `make`, re-slicing |
| `append.go` | `append`, len/cap growth, wasted capacity, pre-allocating with `make` |
| `slices_more.go` | `range`, slice defaults, `copy`, append beyond cap breaks sharing |
| `slices_2d.go` | 2D slice literal (tic-tac-toe), 2D with `make`, 3D slice, 2D slice of struct |
| `main.go` | runs every example |

> Fix made: `printSlices` used `fmt.Println` with `%d`. `Println` doesn't understand format verbs, so it is now `fmt.Printf`.
> Maps used to live in this folder; they are now in `09_maps`.

## 8. How to Run

```bash
go run ./08_slices
```

## 9. Revision Checklist

- [ ] What three things make up a slice?
- [ ] Why did changing `b[0]` also change `a`?
- [ ] `make([]int, 0, 5)`: what are len and cap?
- [ ] Why must you write `s = append(s, x)` and not just `append(s, x)`?
- [ ] When does an appended slice stop sharing memory with the original?
- [ ] `b := a` vs `copy(b, a)`?
- [ ] How do you build a 3×4 grid with `make`?
