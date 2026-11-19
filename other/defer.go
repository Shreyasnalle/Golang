package main
// golang, expects the in directory files to have the same package name so they name it under the name of the folder

import (
	"fmt"
)
// consider defer to be like a stack, where all the items are stored in it, and which are being executed only when the other code in the function is being executed, due to this feature of defer, it is immedetiately called for after the main execution, so that, if by any chance the code breaks in b etween, then the db connection, the method call, payment gateway is never left open because of the defer it get's closed up

func defer_go() {
	defer fmt.Println("From Shreyas to Golang")
	fmt.Println("Hello, World")
}

// so, as the requiremets.txt in python and the package.json in next.js or an js applications, here in go we have go mo, to install it just do go mod init <filename>

