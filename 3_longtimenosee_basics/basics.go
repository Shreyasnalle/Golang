package main

// Go is directory-based, all .go files in the exact same folder MUST share the exact same package name, we cannot mix different package names in one folder. If we want a separate package to import, we must create a subdirectory (subfolder) for it.

import (
	"fmt"
)

func main(){
	fmt.Println("hello, golang")
}
