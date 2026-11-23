# Learning Go 🐹

My notes and code from learning Go, one topic per folder, in the order I learned them. Every folder is a runnable program with its own `doc.md` for revision.

---

## Roadmap

| # | Folder | Topics | Key files |
| :-: | :--- | :--- | :--- |
| 01 | [`01_introduction`](01_introduction/doc.md) | `package main`, imports, `func main()`, exported names, compiled vs interpreted, `go run` vs `go build`, `GOOS`/`GOARCH`, `go.mod` | `main.go`, `go.mod` |
| 02 | [`02_basics`](02_basics/doc.md) | functions, multiple & named returns, `:=` vs `var`, zero values, static typing, `for` loops, constants, type conversion, format specifiers | `main.go`, `loops.go`, `types.go`, `format_specifiers.go` |
| 03 | [`03_revision_flow_control`](03_revision_flow_control/doc.md) | the one-folder-one-package rule, `if` with short statement, `switch`, tagless `switch` as if-else, `fallthrough` | `basics.go`, `conditionswitch.go`, `flow_control.go` |
| 04 | [`04_defer`](04_defer/doc.md) | `defer`, LIFO stack order, argument evaluation, cleanup, `panic` / `recover`, same-package access | `defer.go`, `main.go` |
| 05 | [`05_pointers`](05_pointers/doc.md) | reference types, `&` and `*`, pass by value vs pointer, nil pointers, `new` | `references.go`, `pointers_more.go` |
| 06 | [`06_structs`](06_structs/doc.md) | structs, literals, pointer to struct, nested & anonymous structs, comparison, exported fields, JSON + struct tags | `struct.go`, `struct_more.go` |
| 07 | [`07_arrays`](07_arrays/doc.md) | fixed-size arrays, `[...]`, arrays are copied, 2D arrays | `arrays.go`, `arrays_more.go` |
| 08 | [`08_slices`](08_slices/doc.md) | slicing, len/cap, `make`, `append` growth, nil slice, `copy`, `range`, 2D/3D slices | `slices.go`, `append.go`, `slices_more.go`, `slices_2d.go` |
| 09 | [`09_maps`](09_maps/doc.md) | maps, `make`, insert/update/delete, comma ok, iteration order, nil map, word count, map vs struct | `maps.go`, `maps_more.go` |
| 10 | [`10_function_values_closures`](10_function_values_closures/doc.md) | functions as values, function types, higher-order functions, closures (adder, fibonacci, counter) | `function_values.go`, `closure.go`, `*_more.go` |
| 11 | [`11_methods`](11_methods/doc.md) | methods vs functions, value vs pointer receivers, auto `&`/`*`, methods on non-struct types | `methods.go`, `methods_more.go` |
| 12 | [`12_interfaces`](12_interfaces/doc.md) | interfaces, polymorphism, `any`, type assertion & switch, `Stringer`, `error`, `errors.Is`/`As`, `%w` | `interfaces.go`, `stringer.go`, `error.go`, `*_more.go` |
| 13 | [`13_generics`](13_generics/doc.md) | type parameters, `comparable`/`any`, custom constraints with `~`, generic types (`stack[T]`) | `generics.go`, `generics_more.go` |
| 14 | [`14_goroutines`](14_goroutines/doc.md) | `go` keyword, concurrency vs parallelism, scheduler, `sync.WaitGroup`, race conditions, `sync.Mutex` | `concurrency.go`, `waitgroup_mutex.go` |
| 15 | [`15_channels`](15_channels/doc.md) | channels, buffered vs unbuffered, `close` + `range`, directional channels, `select`, timeouts, worker pool, deadlock | `channels.go`, `bufferunbuffer.go`, `select.go`, `channels_more.go` |

---

## How to Run

Requires Go **1.22+** (`go version`).

```bash
# from the repo root, run any topic:
go run ./02_basics
go run ./15_channels

# check everything compiles (01_introduction is its own module, see below):
go vet ./...
(cd 01_introduction && go run .)
```

Always use `go run .` / `go run ./folder`, **not** `go run main.go`. Each folder's code is spread across several files of the same package, and `go run main.go` compiles only that one file.

---

## How This Repo Is Organised

* **One folder = one topic = one `package main`.** Every folder has its own `func main()` that calls every example in that folder, with a `--- heading` before each one so the output is easy to follow.
* **The original file** (e.g. `struct.go`) holds what I first wrote while learning. **`*_more.go` files** add extra examples for the same topic.
* **Comments** explain the *why* in plain words. Every function has a comment above it saying what it does. Notes starting with `fixed:` or `(note: ...)` / `(more exactly: ...)` mark corrections added during revision; my original comments are kept as they were.
* **`doc.md`** in every folder: the concepts, code snippets, a table of the files, how to run it, and a **revision checklist**.
* **Two modules:** the root `go.mod` is `module golang` (folders 02–15). `01_introduction` has its **own** `go.mod` (`module mod_file`) as a go.mod learning example, so run it from inside that folder.

---

## Fixes Made During Revision

| Where | Bug | Fix |
| :--- | :--- | :--- |
| `03` `swap` | returned `x, y` (nothing swapped) | `return y, x` |
| `03` `switchh` | printed `"Linu"` | `"Linux"` |
| `08` `printSlices` | `fmt.Println` with `%d` | `fmt.Printf` |
| `09` `mapping_in_declaration` | printed `m` instead of `mapping` | prints `mapping` |
| `13` `index` | `func index(t comparable)(s []t, x t)`, invalid syntax | `func index[T comparable](s []T, x T)` |
| `13` `generic` | still called the non-generic functions | calls `index` |
| `14` `say` | `i < 5: i ++`, didn't compile | `i < 5; i++` |
| `15` `fibonacci` | `x*y` gave `0 1 0 0…` | `x+y` |
| all folders | many functions were never called | every `main()` now runs everything |

---

## Quick Cheat Sheet

```go
x := 10                       // declare + assign (inside functions only)
var y int = 20                // works anywhere
const pi = 3.14

func f(a, b int) (int, error) // multiple returns, error last
if v, err := f(1, 2); err != nil { }

for i := 0; i < n; i++ { }    // the only loop keyword
for k, v := range m { }

s := make([]int, 0, 10)       // slice: len 0, cap 10
s = append(s, 1)
m := map[string]int{}         // map
type T struct{ A int }        // struct
func (t *T) Set(a int) { t.A = a }  // pointer-receiver method
type Shape interface{ Area() float64 }
func Index[T comparable](s []T, x T) int  // generics

go work()                     // goroutine
ch := make(chan int, 3)       // buffered channel
select { case v := <-ch: _ = v; default: }
defer cleanup()
```
