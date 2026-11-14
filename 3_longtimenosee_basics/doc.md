# 03 - Long Time No See: Go Revision & Core Concepts

A comprehensive revision guide summarizing the key learnings, rules, code structures, and explanations practiced in `3_longtimenosee_basics`.

---

## 1. Packages & Directory Architecture

### Package `main` vs. Other Packages
- **`package main`**: A reserved package name indicating to the Go compiler that this file is part of an **executable program**. It must contain an entry point: `func main()`.
- **Any other package name** (e.g., `package basics`, `package utils`): Signals a **shared library / reusable package** intended to be imported by other code, not directly executed as a standalone binary.

### The Directory Rule: One Directory = One Package
- Go organizes packages **strictly by directory**.
- **Every `.go` file in the exact same folder MUST declare the exact same package name.**
- You **cannot** mix package names in the same folder (e.g., having `package main` in `main.go` and `package basics` in `basics.go` side-by-side produces a compiler error: `found packages main (...) and basics (...) in ...`).
- **To create a separate package to import:** You must create a **subdirectory (subfolder)** and put files for that package inside it.

---

## 2. Compilation & Execution Models

| Concept | Explanation | Go Context |
| :--- | :--- | :--- |
| **Statically Typed** | Types are verified at **compile time** before running. Variable types are fixed once declared. | Go catches type mismatches early during compilation. |
| **Dynamically Typed** | Types are resolved at **runtime** while code is running (e.g., Python, JavaScript). | Faster to prototype, but errors only show up when the line runs. |
| **Compiled** | Source code is translated directly into machine code binaries by a compiler. | Go compiles directly into machine code for fast native execution. |
| **Interpreted** | Code is read and executed line-by-line by an interpreter or runtime virtual machine. | Python, Ruby, PHP. |

### `go run` vs. `go build`
- **`go run <files>`**: Compiles the source files, stores the temporary executable binary in cache, executes it immediately, and cleans it up afterward. Great for quick development and testing.
- **`go build`**: Compiles the package and saves a permanent standalone binary file directly to your disk for production or distribution.

---

## 3. Variable Declarations: `:=` vs. `var`

### Scope Rules for `:=` and `var`

```go
// Package-level (OUTSIDE any function):
// Must use the 'var' keyword; the short declaration operator ':=' is NOT allowed here!
var name string        // Declared without initial value (defaults to zero value: "")
var age int = 20       // Declared with explicit type and initial value

func main() {
    // Function-level (INSIDE a function):
    // ':=' is short declaration & initialization (infers type automatically)
    city := "Bangalore" 
    
    // '=' is used strictly for assignment / reassignment
    city = "Hyderabad"
}
```

- **`:=`**: Declares and initializes a variable in a single step with automatic type inference. **Only valid inside functions.**
- **`=`**: Assigns/updates a value into an already declared variable.
- **`var`**: Can be used both outside (package-level) and inside functions.

---

## 4. Functions & Return Signatures

### A. Basic Function with Parameter & Return Types
In Go, parameter types and the return type are explicitly specified:

```go
func add(x int, y int) int {
    return x + y
}
// Parameters: x (int), y (int)
// Return type: int
```

---

### B. Multiple Return Values
Go natively supports returning multiple values from a function:

```go
func swap(x string, y string) (string, string) {
    return x, y
}
// Returns two strings: (string, string)
```

---

### C. Named Return Values vs. Explicit Return Values

Go offers two ways to handle return values:

#### 1. Named Returns (with "Naked" Return)
```go
func split(sum int) (x int, y int) {
    x = sum + 2    // 'x' is already declared in the return signature, so use '='
    y = sum - 2    // 'y' is already declared in the return signature, so use '='
    return         // Naked return: automatically returns the current values of x and y
}
```
- The variables `x` and `y` are declared directly in the return signature `(x int, y int)`.
- Because they are already declared, you assign to them using `=` (do **not** use `:=`).
- A bare `return` automatically returns `x` and `y`.

#### 2. Standard Explicit Returns
```go
func splitting(sum int) (int, int) {
    x := sum + 2   // Declare and initialize local variables inside the function
    y := sum - 2
    return x, y    // Explicitly return x and y
}
```
- Only the data types `(int, int)` are declared in the return signature.
- Variables `x` and `y` are declared locally inside the function body using `:=`.
- You must explicitly write `return x, y`.

---

## 5. Summary Cheat Sheet

| Feature | Syntax / Rule | Example |
| :--- | :--- | :--- |
| **Executable Entry** | Package must be `main` and have `func main()` | `package main` |
| **Directory Rule** | All files in a folder share one package name | Subfolders needed for other packages |
| **Outside Functions** | Must use `var` | `var count int = 10` |
| **Inside Functions** | Can use `:=` or `var` | `count := 10` |
| **Multiple Returns** | Specified in parentheses `(type1, type2)` | `(string, error)` |
| **Named Returns** | Declared in return signature, uses `=` then `return` | `func f() (res int) { res = 5; return }` |
