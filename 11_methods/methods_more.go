package main

import (
	"fmt"
	"strings"
)

// scale_value -> the SAME code as scalping but with a value receiver, it only changes its own copy
// this is the mistake to avoid: if a method must change the struct, it needs a pointer receiver
func (s vertex) scale_value(f float64) {
	s.x = s.x * f
	s.y = s.y * f
}

// value_receiver_does_not_mutate -> proves the copy behaviour side by side
func value_receiver_does_not_mutate() {
	v := vertex{3, 4}
	v.scale_value(10)
	fmt.Println("after value receiver scale  :", v) // {3 4}, unchanged
	v.scalping(10)
	fmt.Println("after pointer receiver scale:", v) // {30 40}, changed
}

// auto_address_and_deref -> go writes the & and * for us when calling methods
// v.scalping(2) on a plain value is turned into (&v).scalping(2)
// p.shreyas() on a pointer is turned into (*p).shreyas()
func auto_address_and_deref() {
	v := vertex{1, 1}
	v.scalping(2) // go: (&v).scalping(2)
	p := &vertex{3, 4}
	fmt.Println("v:", v, "| p.shreyas():", p.shreyas()) // go: (*p).shreyas()
	// with a NORMAL function there is no such help, abs(p) would be a compile error, we must write abs(*p)
	fmt.Println("abs(*p):", abs(*p))
}

// my_float -> methods are not only for structs, ANY type declared in this package can have methods
type my_float float64

// absolute -> a method on a non struct type
func (f my_float) absolute() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// word_list -> a named slice type with a method on it
type word_list []string

// shout -> joins all words in capital letters
func (w word_list) shout() string {
	return strings.ToUpper(strings.Join(w, " "))
}

// methods_on_other_types -> calling the methods on my_float and word_list
func methods_on_other_types() {
	f := my_float(-7.5)
	fmt.Println("my_float(-7.5).absolute() =", f.absolute())
	fmt.Println(word_list{"go", "has", "methods"}.shout())
	// note: we cannot add methods to types from OTHER packages directly, e.g. func (i int) double() is not allowed
	// that is exactly why we create our own type (my_float) first
}

// bank_account -> a small real world example, a pointer receiver for changes and a value receiver for reading
type bank_account struct {
	owner   string
	balance float64
}

// deposit -> changes the balance, so pointer receiver
func (b *bank_account) deposit(amount float64) {
	b.balance += amount
}

// withdraw -> changes the balance, returns false if there is not enough money
func (b *bank_account) withdraw(amount float64) bool {
	if amount > b.balance {
		return false
	}
	b.balance -= amount
	return true
}

// summary -> only reads, so value receiver is fine
func (b bank_account) summary() string {
	return fmt.Sprintf("%s has %.2f", b.owner, b.balance)
}

// bank_example -> uses all three methods
func bank_example() {
	acc := bank_account{owner: "shreyas"}
	acc.deposit(500)
	ok := acc.withdraw(800)
	fmt.Println("withdraw 800 ok?", ok)
	acc.withdraw(200)
	fmt.Println(acc.summary())
}
