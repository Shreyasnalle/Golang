# 01 - Introduction to Go: Concepts & Learnings

This document summarizes the core concepts, architecture, commands, and code learned in this introductory section.

---

## 1. What is Go and How Does It Run?

### Compiled vs. Interpreted Languages
* **Compiled Languages (Go, C, Rust):** 
  * The source code is translated ahead-of-time (AOT) by the compiler directly into native machine code (binary `0`s and `1`s).
  * The output is a permanent, standalone binary executable file on disk.
  * The target machine **does not need Go installed** to run the compiled binary.
* **Interpreted Languages (Python, JavaScript):**
  * The source code is read and converted line-by-line into CPU instructions on the fly in memory (RAM).
  * No standalone binary executable is saved to disk.
  * The target computer **must have the interpreter/runtime installed** (e.g., Python, Node.js).

---

## 2. Go Execution: `go run` vs. `go build`

| Command | What It Does Behind the Scenes | Output File |
| :--- | :--- | :--- |
| `go run main.go` | Compiles code into a temporary system cache folder (e.g., `/tmp/go-build...`), runs it immediately, and deletes the binary when done. | None left on disk (temporary) |
| `go build main.go` | Compiles source code into a permanent standalone executable binary for the current system. | `./01_introduction` or `./main` |

---

## 3. Cross-Compilation (`GOOS` & `GOARCH`)

Go has built-in cross-compilation out of the box, meaning you can generate executables for other operating systems and CPU chips without needing external tools.

* **`GOOS` (Target Operating System):**
  * `windows` — Microsoft Windows
  * `linux` — Linux
  * `darwin` — macOS
* **`GOARCH` (Target CPU Architecture):**
  * `amd64` — 64-bit Intel and AMD processors (most laptops/desktops)
  * `arm64` — 64-bit ARM processors (Apple Silicon M-series, Raspberry Pi, mobile)
  * `386` — 32-bit x86 processors

### Example: Building a Windows executable from Linux
```bash
GOOS=windows GOARCH=amd64 go build -o app.exe main.go
```
The resulting `app.exe` can run on any 64-bit Windows PC without installing Go.

---

## 4. Go Modules & `go.mod`

A Go module is a collection of Go packages stored in a file tree with a `go.mod` file at its root.

### Why `go.mod` is Essential:
1. **Project Identification:** Defines the module path/name (e.g., `module mod_file` or `github.com/username/project`).
2. **Version Control:** Specifies the Go toolchain version (e.g., `go 1.22.2`), ensuring consistency across machines.
3. **Dependency Registry:** Tracks third-party packages, external dependencies, and their exact versions so other developers can reproduce the build.

### Key Module Commands:
```bash
# Initialize a new module
go mod init <module-name>

# Download needed dependencies and clean up unused ones
go mod tidy

# Add a third-party package
go get <package-path>
```

> **Comments in `go.mod`:** Both single-line (`//`) and multi-line (`/* ... */`) comments are valid in `go.mod`.

---

### Breakdown:

1. **`package main`:**
   * Every `.go` file must belong to a package.
   * `main` is a **reserved package name** indicating an executable application (not a library).
   * It tells the compiler to look for the `func main()` entry point.
   * If named anything else (e.g., `package utils`), Go treats it as a reusable library intended to be imported.

2. **`import (...)`:**
   * Brings in code from the standard library or third-party packages.
   * Multiple packages can be grouped inside parentheses.
   * Packages used here:
     * **`"fmt"`** (*Formatting*): Formatted I/O functions like `Println`, `Printf`, `Sprintf`.
     * **`"math/rand"`**: Pseudo-random number generation (e.g., `rand.Intn(n)` returns a random integer in `[0, n)`).
     * **`"math"`**: Mathematical constants (like `math.Pi`) and functions.

3. **`func main()`:**
   * The entry point of any executable Go program.
   * Execution starts and ends here.
