package main
// so in go, every .go file should mention at the top telling which pakage it belongs to, main is a reserved package name, which tells go compiler that this file is a executable program and not a shared library
// if the name of the package was other than main, then go complier will thing that this file is to be imported by other code and not directly run as an standalone executable

import (
	"fmt"
	"math/rand"
	"math"
)
// fmt stands for formatting, which is an build in package from go's standard library
// math/rand -> the package name is the LAST part of the import path, so we call it as rand.Intn() and not math/rand.Intn()
// math -> gives us constants like math.Pi and functions like math.Sqrt()

// main() is the entry point, go always starts executing the program from here and the program ends when main() returns
func main(){
	fmt.Println("Hello, Golang")
	fmt.Println("My favourite number is,", rand.Intn(19))
	fmt.Println(math.Pi)
	// so this calls the println function from the fmt package provided by go's standard library

	exported_names()
}

// so in go, a name is EXPORTED (can be used from outside its package) only if it starts with a capital letter
// that's why we write fmt.Println, math.Pi, rand.Intn, all start with capital, if we wrote math.pi it would give an error, because pi with small p is not exported, it's private to the math package
// this is go's way of public / private, there are no public or private keywords like java, the first letter itself decides it
func exported_names() {
	fmt.Println("math.Pi is exported, so we can use it:", math.Pi)
	// fmt.Println(math.pi) -> compile error: name not exported by package math
}

// go is complied programming language, which basically means that it can create a binary file for any type of os, and by directly running this file in that os, we can get the same results, without donwloading or installing go and this can be done by GOOS and GOARCH

// compile languages, they complie the binary file first before running it, while interpreted language run at runtime, that is the binary codes are generated in the RAM, along the way and send forwad to CPU, without ever storing those binaries on disk

// with go run, the binary file is actually created but it is stored in the cache and deleted automatically once it run's while with go build, the binary file stays on the disk

// GOOS=windows GOARCH=amd64 go build -o app.exe main.go, the GOOS tells the targeted OS while GOARCH tells about the CPU chip for that os
