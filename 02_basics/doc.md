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
