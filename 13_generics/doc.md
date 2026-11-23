# 13 - Generics

**Generics** (Go 1.18+) let you write **one** function or type that works for **many types**, instead of copy-pasting the same logic per type.

---

## 1. The Problem

```go
func indexint(s []int, x int) int { ... }          // same logic
func indexstring(s []string, x string) int { ... } // same logic, different type
```
More types means more copies and more boilerplate.

---

## 2. Generic Function

```go
func index[T comparable](s []T, x T) int {
    for i, v := range s {
        if v == x {
            return i
        }
    }
    return -1
}

index([]int{1, 2, 3}, 3)          // T inferred as int → 2
index([]string{"a", "b"}, "b")    // T inferred as string → 1
index[string](ss, "hello")        // T written explicitly
```

* `[T comparable]` is the **type parameter list**. It goes **right after the function name**, in square brackets.
* `T` is a placeholder type. `comparable` is the **constraint**: `T` must support `==` / `!=`.

> Fix made: the original was `func index(t comparable)(s []t, x t)`, which is not valid Go. Also, `generic()` now actually calls `index`.

---

## 3. Constraints

| Constraint | Allows |
| :--- | :--- |
| `any` | every type (same as `interface{}`) |
| `comparable` | types that support `==` / `!=` |
| custom type set | exactly the listed types |

```go
type number interface {
    ~int | ~int64 | ~float64     // a type set
}
func sum[T number](values []T) T { ... }   // + works because every type in the set supports +
```

`~int` means "`int` **or any type whose underlying type is `int`**", e.g. `type my_int int`.

---

## 4. Multiple Type Parameters

```go
func map_slice[T any, U any](items []T, fn func(T) U) []U
map_slice([]string{"go"}, func(s string) int { return len(s) })   // []int{2}
```

---

## 5. Generic Types

```go
type stack[T any] struct { items []T }
func (s *stack[T]) push(item T) { ... }
func (s *stack[T]) pop() (T, bool) { var zero T; ... }

var s stack[string]     // a stack of strings
var n stack[int]        // a stack of ints

type pair[K comparable, V any] struct { key K; value V }
```

`var zero T` is a handy way to get the zero value of whatever `T` is.

---

## 6. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `generics.go` | the duplication problem (`indexint` / `indexstring`) and the generic `index[T comparable]` |
| `generics_more.go` | custom `number` constraint with `~`, `sum`, `map_slice[T, U]`, `filter`, generic `stack[T]`, `pair[K, V]` |
| `main.go` | runs every example |

## 7. How to Run

```bash
go run ./13_generics
```

## 8. Revision Checklist

- [ ] Where do type parameters go in a function signature?
- [ ] Why does `index` need `comparable` and not `any`?
- [ ] Why can't `sum` use `comparable` as its constraint?
- [ ] What does `~` mean in `~int`?
- [ ] How do you declare a stack that only holds strings?
