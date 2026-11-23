package main

import (
	"fmt"
	"time"
)

// this myerror struct is not a custom error of golang, it's something which we have coded out
type MyError struct {
	when time.Time
	what string
}

// this error method is a by default golang method signature, it's by default from go, it's not a custome name
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.when, e.what)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}	
}

func error_working() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// so the error_working() function runs the run() function, which returns an change to the myerror struct, as it has an pointer to mystruct, and as the e is the instance of he mystruct, the error message get's formated into the format given in the inbuild error pointer method function, which is returning an sprintf, which get's printed out from the error_working() function