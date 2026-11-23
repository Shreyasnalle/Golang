# 03 - Long Time No See: Go Revision & Core Concepts

A comprehensive revision guide summarizing the key learnings, rules, code structures, and explanations practiced in `03_revision_flow_control` (originally the `3_longtimenosee_basics` folder, a "long time no see" revision session).

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

### Package-Level Scope Conflict Across Files
- All files in the same folder share the same package namespace.
- If you declare a package-level variable (e.g., `var age int = 20` in `basics.go`), declaring another package-level variable with the same name `age` in another file in the same directory (like `conditionswitch.go`) will cause a **re-declaration compile error**!

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

## 5. Flow Control: `if` with Short Statements

In Go, an `if` statement can start with a short initialization statement before the condition, separated by a semicolon `;`:

```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    } else {
        fmt.Printf("sorry")
    }
    return lim
}
```

### Key Rules:
1. **Short Statement syntax**: `if <initialization>; <condition> { ... }`.
2. **Variable Scoping**: Any variable declared in the short statement (like `v := math.Pow(x, n)`) is **only accessible inside the `if` and its corresponding `else` blocks**. It is not available outside!
3. **`else` Placement Rule**: In Go, `else` or `else if` **must be placed on the exact same line as the closing brace `}`** of the `if` block (`} else {`). If you put `else` on a new line, the Go compiler will automatically insert a semicolon and throw a syntax error.

---

## 6. Flow Control: `switch` Statements

Go's `switch` is cleaner and safer than in languages like C, C++, or Java.

```go
func switchh() {
    fmt.Print("Go runs on ") // fmt.Print does not append a newline, keeping the output on one line

    // Short statement before switch condition:
    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X")
    case "linux":
        fmt.Println("Linux")
    default:
        // freebsd, openbsd, windows...
        fmt.Printf("%s.\n", os)
    }
}
```

### Key Differences & Features:
1. **Automatic Break (No `break` needed!)**:
   - Go runs **only the selected case** and then immediately exits the switch.
   - You do **not** need to write `break;` at the end of each case!
   - (If you ever need fall-through behavior, Go provides an explicit `fallthrough` keyword).
2. **Short Initialization Statement**:
   - `switch os := runtime.GOOS; os {` initializes `os` and tests it immediately.
   - `os` is scoped strictly within the switch block.
3. **`runtime.GOOS`**:
   - Built-in constant from the `runtime` package representing the target operating system (e.g., `"darwin"` for macOS, `"linux"` for Linux, `"windows"` for Windows).
4. **`default` case**:
   - Executes if none of the cases match.

---

## 7. Summary Cheat Sheet

| Feature | Syntax / Rule | Example |
| :--- | :--- | :--- |
| **Executable Entry** | Package must be `main` and have `func main()` | `package main` |
| **Directory Rule** | All files in a folder share one package name | Subfolders needed for other packages |
| **Outside Functions** | Must use `var` | `var count int = 10` |
| **Inside Functions** | Can use `:=` or `var` | `count := 10` |
| **Multiple Returns** | Specified in parentheses `(type1, type2)` | `(string, error)` |
| **Named Returns** | Declared in return signature, uses `=` then `return` | `func f() (res int) { res = 5; return }` |
| **`if` with Init** | Statement runs before condition check; variable scoped to `if`/`else` | `if v := calc(); v > 0 { ... }` |
| **`else` Placement** | Must be on same line as closing brace | `} else {` |
| **`switch` Break** | Automatic breaking; no explicit `break` needed | `case "linux": fmt.Println(...)` |
| **`switch` with Init** | Short statement runs before case matching | `switch os := runtime.GOOS; os {` |
| **Tagless `switch`** | `switch { case cond: }` = cleaner if-else chain | `switch { case h < 12: ... }` |
| **`fallthrough`** | Runs the next case body without checking it | `case 1: ...; fallthrough` |

---

## 8. More `if` / `switch` Forms (`flow_control.go`)

### `if / else if / else`
```go
if marks >= 90 {
    fmt.Println("A")
} else if marks >= 60 {
    fmt.Println("B")
} else {
    fmt.Println("C")
}
```
No `()` around conditions, but `{}` are **always** required.

### `switch` with no condition = `if-else` chain
```go
switch {              // same as: switch true {
case hour < 12:
    fmt.Println("Good morning")
case hour < 17:
    fmt.Println("Good afternoon")
default:
    fmt.Println("Good evening")
}
```
Cases are checked **top to bottom**; the first `true` one runs.

### Multiple values in one case
```go
switch day {
case "sat", "sun":
    fmt.Println("weekend")
}
```

### `fallthrough`
```go
switch n {
case 1:
    fmt.Println("one")
    fallthrough        // runs case 2's body WITHOUT checking case 2
case 2:
    fmt.Println("two")
}
```
Rarely needed. It only falls into the **very next** case.

---

## 9. Files in This Folder

| File | What it covers |
| :--- | :--- |
| `main.go` | comments only: package `main`, static vs dynamic, compiled vs interpreted, `go run` vs `go build` (no `main()` here, which is allowed) |
| `basics.go` | the directory rule, `add`, `swap`, named vs plain returns, package-level `var` |
| `conditionswitch.go` | `if` with short statement (`pow`), `switch` with init (`switchh`), and **the only `main()`** |
| `flow_control.go` | `if/else if/else`, tagless switch, multi-value case, `fallthrough` |

> Fixes made: `swap` now actually swaps (`return y, x`), the `"Linu"` typo now prints `"Linux"`, and `main()` calls every function.

## 10. How to Run

```bash
go run ./03_revision_flow_control
```

## 11. Revision Checklist

- [ ] Can two files in the same folder have different package names?
- [ ] Why does declaring `age` in two files of the same folder fail?
- [ ] Where can the variable from `if v := ...; v < lim` be used?
- [ ] Why must `else` be on the same line as `}`?
- [ ] How do you write an if-else chain using `switch`?
- [ ] What does `fallthrough` do, and does it check the next case's condition?
