# 10 - Function Values & Closures

In Go, **functions are values**, just like `int` or `string`. You can store them in variables, pass them as arguments, return them, and keep them in maps.

> **Naming note:** this topic was first written as "asynchronous functions" (`asynchronous_functioon`). The correct terms are **function values**, **higher-order functions** and **callbacks**. *Asynchronous* means running in the background without blocking, which Go does with **goroutines** (`14_goroutines`).

---

## 1. Function as a Value / Parameter

```go
func compute(fn func(float64, float64) float64) float64 {
    return fn(3, 4)           // calls whatever function was passed in
}

hypot := func(x, y float64) float64 {   // anonymous function stored in a variable
    return math.Sqrt(x*x + y*y)
}
hypot(5, 12)       // 13
compute(hypot)     // 5, hypot passed as an argument (no parentheses!)
```

---

## 2. Named Function Types & Maps of Functions

```go
type operation func(int, int) int

func apply(a, b int, op operation) int { return op(a, b) }

calculator := map[string]operation{
    "+": func(a, b int) int { return a + b },
}
calculator["+"](2, 3)
```

---

## 3. Immediately Invoked Function

```go
result := func(name string) string { return "hello " + name }("shreyas")
```

---

## 4. Higher-Order Functions

A function that **takes** or **returns** a function:
```go
map_strings([]string{"go"}, strings.ToUpper)   // takes a function
double := multiplier(2)                         // returns a function
double(5)                                       // 10
```

---

## 5. Closures

A **closure** is an inner function that **references variables from the outer function**. It can read and change them **even after the outer function has returned**.

```go
func adder() func(int) int {
    sum := 0                    // stays alive inside the returned function
    return func(x int) int {
        sum += x
        return sum
    }
}

pos, neg := adder(), adder()    // two closures, two SEPARATE sums
```

### Uses
| Example | Idea |
| :--- | :--- |
| `adder()` | running total per closure |
| `fibonacci()` | generator: every call returns the next number |
| `counter()` | private state, only changeable through the returned function (Go's "private field" without classes) |
| `closure_modifies_outer()` | closure changes `total` in the enclosing function |

---

## 6. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `function_values.go` | (was `asynchronous.go`) function passed as a parameter, anonymous function in a variable |
| `function_values_more.go` | named function types, map of functions, immediately invoked function, higher-order `map_strings`, `multiplier` |
| `closure.go` | definition of a closure, `adder()` |
| `closure_more.go` | two independent adders, fibonacci generator, counter with private state, closure changing an outer variable |
| `main.go` | runs every example |

## 7. How to Run

```bash
go run ./10_function_values_closures
```

## 8. Revision Checklist

- [ ] What is the type of `func(x, y float64) float64 { ... }`?
- [ ] Why do we pass `hypot` and not `hypot()` to `compute`?
- [ ] Why don't `pos` and `neg` share the same `sum`?
- [ ] How does `fibonacci()` remember `a` and `b` between calls?
- [ ] Is passing a function "asynchronous"? What is async in Go?
