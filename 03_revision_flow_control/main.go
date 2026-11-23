package main
// in go every file should mention at top which package it belong to, main is a researved package page which tells go complier that this is the executable program and not a shared library, if the package name is somehting other than main then go complier things this file is to be imported by other code and it's not a direct standalone executable code

// statically -> code checked at compile time (before running the code), while dynamically -> code is checked at runtime(while running the code), here variable types are checked, or the code simply putting
// compiled -> uses a complier to translate your source code into machine code, while interpreted -> the interpreter or runtime executes the program

// with go run, the exe is stored in cache and deleted automatically later once it run's while with go build the binary file stays on the disk