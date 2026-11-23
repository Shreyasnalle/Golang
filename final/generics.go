package main 

import "fmt"

func indexint(s []int, x int) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func indexstring(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func generics_running() {
	si := []int{1, 2, 3, 4}
	fmt.Println(indexint(si, 15))
	ss := []string{"shreyas", "advit"}
	fmt.Println(indexstring(ss, "hello"))
}

// so as we can see, we have two different functions, which infact has the same purpose as well, but their is only a difference in between the type of the slice being passed to the function, one is the int and the other one is the string, so tis increases the codeline thus increasing the size of the boilerplate

// so we have a concept of generics in go, which eases this work out
func index(t comparable)(s []t, x t) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func generic() {
	si := []int{1, 2, 3, 4}
	fmt.Println(indexint(si, 15))
	ss := []string{"shreyas", "advit"}
	fmt.Println(indexstring(ss, "hello"))
}

// so as we can see the two fucntions are simply being worked upon into a single one, this has reduced the codelines and thus imporving the code