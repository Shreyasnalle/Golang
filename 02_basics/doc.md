# 02 - Go Basics: Concepts & Learnings

This document summarizes the core language fundamentals, functions, variable declarations, type systems, and formatting learned in this section.

---

## 1. Type Systems: Statically Typed vs. Dynamically Typed

* **Statically Typed (Go, Rust, C++):**
  * Data types are verified at **compile time** (before the program runs).
  * Once a variable is declared with a type, it cannot hold a value of a different type.
  * **Benefits:** Catches bugs early, prevents unexpected type mismatches, and allows the compiler to produce faster, highly optimized machine code.
* **Dynamically Typed (Python, JavaScript):**
  * Types are checked at **runtime** (while the code is executing).
  * Variables are dynamic references that can point to an integer at one moment and a string the next.
  * **Trade-off:** Faster for quick scripting, but errors only reveal themselves when that exact line of code is executed.

---

## 2. Integer Sizes: `int8` vs. `int64`

Go provides explicitly-sized integer types:

| Type | Memory Size | Range (Min to Max) | Typical Use |
| :--- | :--- | :--- | :--- |
| **`int8`** | 8 bits (1 byte) | `-128` to `127` | Memory-sensitive systems, byte buffers |
| **`int64`** | 64 bits (8 bytes) | `≈ -9.22 × 10¹⁸` to `≈ 9.22 × 10¹⁸` | Timestamps, database IDs, large counts |
| **`int`** | Platform-dependent (64-bit on 64-bit OS) | Same as `int64` on 64-bit machines | General-purpose integer |

> **Note:** Even on a 64-bit machine, `int` and `int64` are distinct types in Go and require explicit type conversion (e.g., `int64(myInt)`).

---

## 3. Variables: Declaration vs. Assignment

### `:=` vs. `=`
* **`:=` (Short Declaration & Initialization):**
  * Declares a **new** variable and assigns an initial value in one step.
  * Uses type inference (Go figures out the type automatically).
  * **Rule:** Can **only** be used inside functions.
* **`=` (Assignment Only):**
  * Assigns or updates the value of a variable that has **already been declared**.

```go
// Inside a function:
name := "Shreyas"  // Declares 'name' (string) and initializes it
name = "Nalle"     // Reassigns 'name' using '='

// Package-level (outside any function):
var age int = 20   // Must use 'var' keyword; ':=' is not allowed here
var city string    // Declared without initial value (gets zero value)
```

---

## 4. Default Zero Values

In Go, variables declared without an explicit initial value are automatically given their **zero value** (never `undefined` or garbage memory):

| Type | Zero Value |
| :--- | :--- |
| `int`, `int8`, `int64`, etc. | `0` |
| `float32`, `float64` | `0` |
| `bool` | `false` |
| `string` | `""` (empty string) |
| Pointers, slices, maps, channels | `nil` |

---

## 5. Functions in Go

### A. Basic Functions & Multiple Return Values
Go natively supports returning multiple values from a function:

```go
// Single return value
func add(x int, y int) int {
    return x + y
}

// Multiple return values
func swap(x string, y string) (string, string) {
    return y, x
}
```

### B. Named Return Values vs. Standard Returns

```go
// 1. Named Return Values (with Naked Return)
func split(sum int) (x int, y int) {
    x = sum * 4 / 9   // 'x' is already declared in signature, use '='
    y = sum - x       // 'y' is already declared in signature, use '='
    return            // Naked return: returns current values of x and y
}

// 2. Standard Returns (Explicit Return)
func splitting(sum int) (int, int) {
    x := sum * 4 / 9  // Declare local variables using ':='
    y := sum - x
    return x, y       // Explicitly return x and y
}
```

* **Named Returns:** Document what each returned value means right in the function signature. Best kept to short functions.
* **Naked Returns:** The bare `return` keyword automatically returns named variables.

---

## 6. Formatting Verbs (Specifiers) in `fmt`

Used with `fmt.Printf()` and `fmt.Sprintf()` to format and print data:

| Verb | Purpose | Example | Output |
| :--- | :--- | :--- | :--- |
| **`%v`** | Value in default format | `fmt.Printf("%v", 42)` | `42` |
| **`%+v`** | Struct with field names | `fmt.Printf("%+v", u)` | `{Name:Shreyas Age:20}` |
| **`%T`** | Data **Type** of the value | `fmt.Printf("%T", "hi")` | `string` |
| **`%s`** | Plain unquoted string | `fmt.Printf("%s", "hello")`| `hello` |
| **`%q`** | Double-quoted string | `fmt.Printf("%q", "hello")`| `"hello"` |
| **`%d`** | Base-10 integer | `fmt.Printf("%d", 100)` | `100` |
| **`%f`** / **`%.2f`** | Float / Float with 2 decimal places | `fmt.Printf("%.2f", 3.1415)` | `3.14` |
| **`%t`** | Boolean (`true` or `false`) | `fmt.Printf("%t", true)` | `true` |
| **`%p`** | Memory address / pointer | `fmt.Printf("%p", &v)` | `0xc000014070` |
| **`%#v`** | Go-syntax representation | `fmt.Printf("%#v", u)` | `main.user{Name:"Shreyas", Age:20}` |
| **`%%`** | A literal `%` sign | `fmt.Printf("100%%")` | `100%` |

> `Println` prints values separated by spaces + newline. `Printf` prints using a format. `Sprintf` **returns** the formatted string and prints nothing.

---

## 7. Loops: `for` Is the Only Loop

Go has no `while` and no `do-while`; `for` covers all of them:

```go
for i := 0; i < 3; i++ { }        // 1. classic: init; condition; post
for n < 100 { n *= 2 }            // 2. while-style: condition only
for { if done { break } }         // 3. infinite: exit with break / return
for i := 0; i < 6; i++ {          // 4. continue skips to the next iteration
    if i%2 == 0 { continue }
}
for index, value := range items { } // 5. range over a collection
for _, value := range items { }     //    _ = blank identifier, ignore the index
```

---

## 8. Constants & Type Conversion

```go
const pi = 3.14                // fixed at compile time, can never be reassigned
const ( maxInt8 int8 = 127 )   // grouped declaration works for const and var

var i int = 42
var f float64 = float64(i)     // explicit conversion: T(value)
var u uint = uint(f)
```

* Go **never** converts types automatically. `int` → `float64` needs `float64(i)`.
* `:=` cannot be used with `const`.
* Type inference with `:=`: `10` → `int`, `10.5` → `float64`, `"go"` → `string`, `true` → `bool`.

---

## 9. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `main.go` | functions, multiple/named returns, `:=` vs `var`, zero values, classic loop, `main()` that runs everything |
| `loops.go` | every form of `for`: classic, while-style, infinite + `break`, `continue`, `range` |
| `types.go` | `const`, grouped declarations, `int8`/`int64` limits, type conversion, type inference |
| `format_specifiers.go` | `%v %+v %#v %T %s %q %d %f %t %p %%` side by side, `Sprintf` |

## 10. How to Run

```bash
go run ./02_basics      # from the repo root
```

## 11. Revision Checklist

- [ ] Difference between `:=` and `=`? Where is `:=` not allowed?
- [ ] Zero value of `int`, `bool`, `string`, slice?
- [ ] Named return vs normal return: when do we use `=` inside the function?
- [ ] How do you write a `while` loop in Go?
- [ ] Why does `var f float64 = i` fail when `i` is an `int`?
- [ ] `Printf` vs `Println` vs `Sprintf`?
