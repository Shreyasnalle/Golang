package main
// so in go, every .go file should mention at the top telling which pakage it belongs to, main is a reserved package name, which tells go compiler that this file is a executable program and not a shared library 
// if the name of the package was other than main, then go complier will thing that this file is to be imported by other code and not directly run as an standalone executable

import (
	"fmt"
	"math/rand"
	"math"
)
// fmt stands for formatting, which is an build in package from go's standard library

func main(){
	fmt.Println("Hello, Golang")
	fmt.Println("My favourite number is,", rand.Intn(19))
	fmt.Println(math.Pi)
	// so this calls the println function from the fmt package provided by go's standard library 
}

// go is complied programming language, which basically means that it can create a binary file for any type of os, and by directly running this file in that os, we can get the same results, without donwloading or installing go and this can be done by GOOS and GOARCH

// compile languages, they complie the binary file first before running it, while interpreted language run at runtime, that is the binary codes are generated in the RAM, along the way and send forwad to CPU, without ever storing those binaries on disk

// with go run, the binary file is actually created but it is stored in the cache and deleted automatically once it run's while with go build, the binary file stays on the disk

// GOOS=windows GOARCH=amd64 go build -o app.exe main.go, the GOOS tells the targeted OS while GOARCH tells about the CPU chip for that os