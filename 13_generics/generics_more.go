package main

import "fmt"

// number -> a custom constraint, it's an interface but used as a LIST OF ALLOWED TYPES (a type set)
// the ~ means "this type or any type whose underlying type is this", so a type like my_int (below) is allowed too
type number interface {
	~int | ~int64 | ~float64
}

// my_int -> a custom type with int underneath, allowed in number because of ~int
type my_int int

// sum -> adds all numbers of any number type, comparable would not work here because comparable does not allow +
func sum[T number](values []T) T {
	var total T // zero value of T, 0 for int and 0.0 for float64
	for _, v := range values {
		total += v
	}
	return total
}

// map_slice -> two type parameters, T for the input and U for the output
// converts a []T into a []U by applying fn to every element
func map_slice[T any, U any](items []T, fn func(T) U) []U {
	out := make([]U, 0, len(items))
	for _, item := range items {
		out = append(out, fn(item))
	}
	return out
}

// filter -> keeps only the elements for which keep returns true
func filter[T any](items []T, keep func(T) bool) []T {
	var out []T
	for _, item := range items {
		if keep(item) {
			out = append(out, item)
		}
	}
	return out
}

// stack -> a GENERIC TYPE, a stack that can hold elements of any single type T
type stack[T any] struct {
	items []T
}

// push -> adds on the top, pointer receiver because it changes the stack
func (s *stack[T]) push(item T) {
	s.items = append(s.items, item)
}

// pop -> removes from the top, returns false if the stack is empty
func (s *stack[T]) pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

// pair -> a generic type with two type parameters
type pair[K comparable, V any] struct {
	key   K
	value V
}

// generics_more -> runs all the extra generic examples
func generics_more() {
	fmt.Println("sum ints  :", sum([]int{1, 2, 3}))
	fmt.Println("sum floats:", sum([]float64{1.5, 2.25}))
	fmt.Println("sum my_int:", sum([]my_int{10, 20}))
	// sum([]string{"a"}) -> compile error, string does not satisfy number

	lengths := map_slice([]string{"go", "generics"}, func(s string) int { return len(s) })
	fmt.Println("map_slice string -> int:", lengths)

	evens := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool { return n%2 == 0 })
	fmt.Println("filter evens:", evens)

	var s stack[string] // T = string for this stack
	s.push("first")
	s.push("second")
	top, _ := s.pop()
	fmt.Println("popped:", top, "| left:", s.items)
	var nums stack[int]
	_, ok := nums.pop()
	fmt.Println("pop on empty int stack ok?", ok)

	p := pair[string, int]{key: "age", value: 20}
	fmt.Printf("pair: %+v\n", p)
}
