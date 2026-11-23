package main
// golang, expects the in directory files to have the same package name so they name it under the name of the folder

import (
	"fmt"
)
// consider defer to be like a stack, where all the items are stored in it, and which are being executed only when the other code in the function is being executed, due to this feature of defer, it is immedetiately called for after the main execution, so that, if by any chance the code breaks in b etween, then the db connection, the method call, payment gateway is never left open because of the defer it get's closed up

// defer_go -> the deferred print runs AFTER Hello, World, i.e. just before defer_go returns
func defer_go() {
	defer fmt.Println("From Shreyas to Golang")
	fmt.Println("Hello, World")
}

// so, as the requiremets.txt in python and the package.json in next.js or an js applications, here in go we have go mo, to install it just do go mod init <filename>

// defer_stack -> proves the stack (LIFO, last in first out) behaviour, the defer written LAST runs FIRST
// output -> counting, done, 2, 1, 0
func defer_stack() {
	fmt.Println("counting")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i) // pushed on the stack as 0, 1, 2, so popped as 2, 1, 0
	}
	fmt.Println("done")
}

// defer_args_evaluated_now -> the ARGUMENTS of a deferred call are calculated at the moment defer is written, not when it finally runs
// so even though x becomes 20 later, the deferred print already captured 10
func defer_args_evaluated_now() {
	x := 10
	defer fmt.Println("deferred x =", x) // x is read right here -> 10
	x = 20
	fmt.Println("current x =", x)
}

// open_db_connection -> a fake db connection, just to show the real world use case of defer (cleanup)
func open_db_connection() {
	fmt.Println("db: connection opened")
	defer fmt.Println("db: connection closed") // written right next to open, so we never forget to close it
	fmt.Println("db: running query 1")
	fmt.Println("db: running query 2")
	// even if we add 10 more return statements in this function, the close will still run before every one of them
}

// defer_runs_on_panic -> panic means the program crashed in between, still the deferred calls run
// recover() inside a deferred function catches the panic, so the program does not die and continues after this function
// this is exactly the "code breaks in between but the connection still gets closed" case from the comment above
func defer_runs_on_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()
	defer fmt.Println("payment gateway: closed (runs even though we panicked)")
	fmt.Println("payment gateway: opened")
	panic("something went wrong in between")
	// nothing after panic runs, only the defers do
}
