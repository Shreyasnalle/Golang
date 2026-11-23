# 11 - Methods

Go has **no classes**. A **method** is a function with a **receiver**: it is attached to a type and called with dot notation.

---

## 1. Function vs Method

```go
func abs(v vertex) float64 { ... }      // function → abs(v)
func (v vertex) abss() float64 { ... }  // method   → v.abss()
//    └ receiver
```

Same logic, different notation. A method ties behaviour to the type, much like calling a method on an object in OOP (`v1.abss()` and `v2.abss()` each work on their own instance).

---

## 2. Value Receiver vs Pointer Receiver

```go
func (s vertex) shreyas() float64 { ... }   // VALUE receiver → gets a COPY
func (s *vertex) scalping(f float64) {      // POINTER receiver → gets the ADDRESS
    s.x = s.x * f                            // changes the original
}
```

```go
v := vertex{3, 4}
v.scale_value(10)  // value receiver   → v still {3 4}
v.scalping(10)     // pointer receiver → v now {30 40}
```

| Use a **pointer receiver** `(s *T)` when... | Use a **value receiver** `(s T)` when... |
| :--- | :--- |
| the method must **change** the struct | the method only **reads** |
| the struct is **large** (avoid copying it) | you want the safety of no mutation |
| other methods of the type already use pointers (consistency) | the type is small (e.g. a couple of numbers) |

---

## 3. Go Adds `&` and `*` for You (Methods Only)

```go
v.scalping(2)   // Go turns it into (&v).scalping(2)
p.shreyas()     // Go turns it into (*p).shreyas()
abs(p)          // ❌ functions get no such help → write abs(*p)
```

---

## 4. Methods on Non-Struct Types

```go
type my_float float64
func (f my_float) absolute() float64 { ... }

type word_list []string
func (w word_list) shout() string { ... }
```

* Any type **declared in your package** can have methods.
* You **can't** add methods to types from other packages (`func (i int) double()` is not allowed). Declare your own type first.

---

## 5. Real-World Example (`bank_account`)

* `deposit`, `withdraw` → **pointer receivers** (they change the balance)
* `summary` → **value receiver** (it only reads)

---

## 6. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `methods.go` | function vs method, value receiver vs pointer receiver, when to use which |
| `methods_more.go` | value receiver can't mutate, auto `&`/`*`, methods on `my_float` and `word_list`, bank account |
| `main.go` | runs every example |

## 7. How to Run

```bash
go run ./11_methods
```

## 8. Revision Checklist

- [ ] What is a receiver?
- [ ] Why did `scale_value` not change `v`?
- [ ] Can you call a pointer-receiver method on a plain variable `v`? Why?
- [ ] Can you add a method to `int` directly?
- [ ] Which receiver for a method that only formats a string?
