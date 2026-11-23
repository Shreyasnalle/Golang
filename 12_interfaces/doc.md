# 12 - Interfaces (incl. Stringer & Error)

An **interface** is a **contract**: a list of method signatures. Any type that has those methods **automatically** satisfies the interface. There is no `implements` keyword.

> A struct describes *what data a thing has*. An interface describes *what a thing can do*.

---

## 1. Basic Interface

```go
type i interface { m() }
type t struct { s string }
func (tt t) m() { fmt.Println(tt.s) }   // t now satisfies i (implicitly)

var ii i = t{"hello"}
ii.m()   // hello
```

---

## 2. Polymorphism

```go
type shape interface {
    area() float64
    perimeter() float64
}
// rect and circle both have area() + perimeter() → both are shapes

func measure(s shape) { ... }            // works for ANY shape
shapes := []shape{rect{3, 4}, circle{1}} // different types, one slice
```

---

## 3. Empty Interface: `interface{}` / `any`

* Has **zero** methods, so **every** type satisfies it.
* `any` is just another name for `interface{}` (Go 1.18+).
* This is why `fmt.Println(a ...any)` accepts anything.

---

## 4. Type Assertion & Type Switch

```go
var value any = "golang"
s, ok := value.(string)   // "golang", true
n, ok := value.(int)      // 0, false (no panic with the comma ok form)
// value.(int) WITHOUT ok → panic

switch v := value.(type) {
case int:    // v is an int here
case string: // v is a string here
case shape:  // works with interfaces too
default:
}
```

---

## 5. Nil Interface

```go
var ii i         // legal: a nil interface
ii == nil        // true
ii.m()           // ❌ panic: nothing concrete inside to call m() on
```

---

## 6. Stringer: `fmt.Stringer`

```go
type Stringer interface { String() string }

func (p person) String() string {
    return fmt.Sprintf("%v (%v years)", p.name, p.age)
}
fmt.Println(person{"shreyas nalle", 20})   // shreyas nalle (20 years)
```

`fmt` checks whether a value has `String()` and uses it automatically.
`Sprintf` **returns** a string; `Print`/`Printf`/`Println` **write to the terminal**.

---

## 7. Error: `error`

```go
type error interface { Error() string }   // built in
```

### Custom error (`error.go`)
```go
type MyError struct { when time.Time; what string }
func (e *MyError) Error() string { return fmt.Sprintf("at %v, %s", e.when, e.what) }
func run() error { return &MyError{time.Now(), "it didn't work"} }
```

### Everyday patterns (`errors_more.go`)
| Pattern | Code |
| :--- | :--- |
| simple error | `errors.New("cannot divide by zero")` |
| `(value, error)` return | `func divide(a, b float64) (float64, error)` |
| check it | `if err != nil { ... }` |
| sentinel error | `var err_not_found = errors.New("item not found")` |
| wrap with context | `fmt.Errorf("buy %q: %w", name, err)` |
| match through wraps | `errors.Is(err, err_not_found)` |
| extract custom type | `var e *MyError; errors.As(err, &e)` |

Errors are **values**: Go has no `try`/`catch`.

---

## 8. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `interfaces.go` | interface as a contract, implicit implementation, `var ii i = t{"hello"}` |
| `interfaces_more.go` | `shape` polymorphism, `any`, type assertion, type switch, nil interface |
| `stringer.go` | `fmt.Stringer` via `String()` on `person` |
| `error.go` | custom error type `MyError` with `Error()` |
| `errors_more.go` | `errors.New`, `(value, error)`, `strconv.Atoi` errors, `%w` wrapping, `errors.Is`, `errors.As` |
| `main.go` | runs every example |

## 9. How to Run

```bash
go run ./12_interfaces
```

## 10. Revision Checklist

- [ ] How does a type "implement" an interface in Go?
- [ ] Why can `[]shape` hold both `rect` and `circle`?
- [ ] What is `any`?
- [ ] `value.(int)` vs `v, ok := value.(int)`?
- [ ] What happens when you call a method on a nil interface?
- [ ] Which method makes `fmt.Println` print your struct nicely?
- [ ] What is `%w` for in `fmt.Errorf`, and how does `errors.Is` use it?
