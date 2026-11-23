package main 

// so interface is like a job or a contract, let's say like a job description, it actually cares about what actions and methods it can do, as the struct has the elements data type and field initialized right, in the same way interface is like a playbook in hich the methods mentioned has to be something which anyone calls ahs to play

import "fmt"

type i interface {
	m()
}

type t struct {
	s string
}

func (tt t) m() {
	fmt.Println(tt.s)
}

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