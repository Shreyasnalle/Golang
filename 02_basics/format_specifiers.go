package main

import "fmt"

// a small struct only to show how %v, %+v and %#v print a struct differently
type user struct {
	Name string
	Age  int
}

// format_specifiers prints the same values with different % verbs so we can see the difference side by side
// Printf -> prints with format, Println -> prints values with spaces and a new line, Sprintf -> does NOT print, it returns the string
func format_specifiers() {
	u := user{"Shreyas", 20}

	fmt.Printf("%%v  -> %v\n", u)  // {Shreyas 20}
	fmt.Printf("%%+v -> %+v\n", u) // {Name:Shreyas Age:20}, field names also
	fmt.Printf("%%#v -> %#v\n", u) // go syntax representation, main.user{Name:"Shreyas", Age:20}
	fmt.Printf("%%T  -> %T\n", u)  // main.user, the type
	fmt.Printf("%%s  -> %s\n", "hello")
	fmt.Printf("%%q  -> %q\n", "hello")
	fmt.Printf("%%d  -> %d\n", 100)
	fmt.Printf("%%f  -> %f | %%.2f -> %.2f\n", 3.14159, 3.14159)
	fmt.Printf("%%t  -> %t\n", true)
	fmt.Printf("%%p  -> %p\n", &u) // memory address, changes every run

	s := fmt.Sprintf("%s is %d years old", u.Name, u.Age) // nothing printed here, just stored in s
	fmt.Println(s)
	// %% prints a literal % sign, that's why the labels above use %%
}
