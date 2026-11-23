package main 

import "fmt"

// indexint -> returns the index of x in an int slice, -1 if it is not there
func indexint(s []int, x int) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// indexstring -> the exact same logic as indexint, only for strings, this duplication is the problem generics solve
func indexstring(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// generics_running -> calls the two non generic versions
func generics_running() {
	si := []int{1, 2, 3, 4}
	fmt.Println(indexint(si, 15))
	ss := []string{"shreyas", "advit"}
	fmt.Println(indexstring(ss, "hello"))
}

// so as we can see, we have two different functions, which infact has the same purpose as well, but their is only a difference in between the type of the slice being passed to the function, one is the int and the other one is the string, so tis increases the codeline thus increasing the size of the boilerplate

// so we have a concept of generics in go, which eases this work out
// index -> ONE function for every type that can be compared with ==
// [T comparable] is the type parameter list, T is a placeholder type, comparable is the constraint (T must support == and !=)
// fixed: this was written as index(t comparable)(s []t, x t), type parameters go inside [ ] right after the function name
func index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

// generic -> the same calls as generics_running but now through the single generic index function
// fixed: this was still calling indexint and indexstring, now it calls index for both
func generic() {
	si := []int{1, 2, 3, 4}
	fmt.Println(index(si, 3))   // T is inferred as int -> 2
	ss := []string{"shreyas", "advit"}
	fmt.Println(index(ss, "advit")) // T is inferred as string -> 1
	fmt.Println(index[string](ss, "hello")) // we can also write the type explicitly -> -1
}

// so as we can see the two fucntions are simply being worked upon into a single one, this has reduced the codelines and thus imporving the code
