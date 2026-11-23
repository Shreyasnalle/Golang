package main 

// so interface is like a job or a contract, let's say like a job description, it actually cares about what actions and methods it can do, as the struct has the elements data type and field initialized right, in the same way interface is like a playbook in hich the methods mentioned has to be something which anyone calls ahs to play

import "fmt"

// i -> an interface with one method m(), ANY type that has a method m() automatically satisfies i
type i interface {
	m()
}

// t -> a struct with one string field
type t struct {
	s string
}

// m -> because t has this method, t implements the interface i, there is no "implements" keyword in go, it's implicit
func (tt t) m() {
	fmt.Println(tt.s)
}

// interface_working -> stores a t inside an interface variable of type i and calls m() through the interface
func interface_working() {
	var ii i = t{"hello"}
	ii.m()
	// other way around to write this 
	// var ii i
	// ii = t{"hello"}
	// ii.m()
}
// so as we can see in the code, in the function we use a var with the inteface as reference and we then pass the t struct to it with the text hello to it, later we have to make sure that the instance of this interface has to use the m function or else it will show an error

// interface cannot work with an empty intilization, there ahs to be something assigned with it
// (more exactly: var ii i is allowed, it's just a nil interface, but calling ii.m() on it panics because there is no concrete value inside to call m() on, see nil_interface() in interfaces_more.go)
