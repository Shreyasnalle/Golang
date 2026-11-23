package main

import (
	"fmt"
	"math"
)

// shape -> any type that has BOTH area() and perimeter() is a shape
type shape interface {
	area() float64
	perimeter() float64
}

// rect and circle -> two totally different structs, both will satisfy shape
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// area -> rect's area
func (r rect) area() float64 { return r.width * r.height }

// perimeter -> rect's perimeter
func (r rect) perimeter() float64 { return 2 * (r.width + r.height) }

// area -> circle's area
func (c circle) area() float64 { return math.Pi * c.radius * c.radius }

// perimeter -> circle's perimeter
func (c circle) perimeter() float64 { return 2 * math.Pi * c.radius }

// measure -> works for ANY shape, it does not know or care if it's a rect or a circle (polymorphism)
func measure(s shape) {
	fmt.Printf("%T -> area %.2f, perimeter %.2f\n", s, s.area(), s.perimeter())
}

// polymorphism -> one slice of shape holding different types, and one function handling all of them
func polymorphism() {
	shapes := []shape{rect{3, 4}, circle{1}}
	for _, s := range shapes {
		measure(s)
	}
}

// empty_interface -> interface{} (or its new name any) has ZERO methods, so EVERY type satisfies it
// that's why fmt.Println(a ...any) can take values of any type
func empty_interface() {
	var anything any
	anything = 42
	fmt.Printf("(%v, %T)\n", anything, anything)
	anything = "hello"
	fmt.Printf("(%v, %T)\n", anything, anything)
	anything = []int{1, 2}
	fmt.Printf("(%v, %T)\n", anything, anything)
}

// type_assertion -> getting the concrete value back OUT of an interface
// value.(string) panics if it's not a string, so we use the comma ok form to check safely
func type_assertion() {
	var value any = "golang"
	s, ok := value.(string)
	fmt.Println("as string:", s, ok)
	n, ok := value.(int) // not an int -> n is 0 and ok is false, no panic
	fmt.Println("as int:", n, ok)
}

// describe -> a type switch, like a normal switch but on the TYPE stored inside the interface
func describe(value any) {
	switch v := value.(type) {
	case int:
		fmt.Println(v, "is an int, double is", v*2)
	case string:
		fmt.Println(v, "is a string of length", len(v))
	case shape:
		fmt.Println("it's a shape with area", v.area())
	default:
		fmt.Printf("don't know the type %T\n", v)
	}
}

// nil_interface -> var ii i is a nil interface, it is legal to declare, but calling a method on it panics
func nil_interface() {
	var ii i
	fmt.Printf("ii = %v, type = %T, ii == nil ? %v\n", ii, ii, ii == nil)
	if ii != nil {
		ii.m() // only call when something is inside
	} else {
		fmt.Println("ii is nil, skipping ii.m() (it would panic)")
	}
}
