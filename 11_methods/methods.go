package main

import (
	"fmt"
	"math"
)

// vertex -> the struct on which we attach all the methods in this folder
type vertex struct {
	x, y float64
}

// abs -> a NORMAL function, vertex is passed as an argument, called as abs(v)
func abs(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// running -> calls the normal function version
func running() {
	v := vertex{3, 4}
	fmt.Println(abs(v))
}

// so basically method is another way of calling function in languages which has the support of classes, there we create classs, create an object of that classss and then call a method over it, but here in golang we don't have classse so we directly method that is th dot notation like

// abss -> the SAME logic as abs but as a METHOD, (v vertex) is the receiver, called as v.abss()
func (v vertex) abss() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// just a different way of notation

// running_method -> calls the method version with dot notation
func running_method() {
	v := vertex{3, 5}
	fmt.Println(v.abss())
}

// so using methods inplace of the functions makes sure that the object or the scope of that method being called is limited to that instance only, like for example in if we had classes, then that class wopuld have 2 objects, then these objects would have diffferent methods being claled for that object, the same thing is here too
// func (v vertex) abs() float64 {}, then the v1.abs() and v2.abs() are two different call's, oop concept

// now there are two two ways i which we can use methods, firt is the normal function way and the second is the pointer varaible way, so as we can understand, in the function way, we can't directly assigs the variables, but in the pointer varaible, we have the pointer to the data strcuture, so this helps us to change the value presetn at that data strcutre

// shreyas -> VALUE receiver, gets a copy of the vertex, only reads it
func (s vertex) shreyas() float64 {
	return math.Sqrt(s.x*s.x + s.y*s.y)
}
// scalping -> POINTER receiver, gets the address of the vertex, so it changes the original x and y
func (s *vertex) scalping(f float64) {
	s.x = s.x * f
	s.y = s.y * f
}
// normal_variable_pointer_variable_method -> first mutate with the pointer method, then read with the value method
func normal_variable_pointer_variable_method() {
	v := vertex{3, 4}
	v.scalping(10)
	fmt.Println(v.shreyas())
}

// so here, we first call the scalping, which changes the value of the vertex struct, thenw we call the shreyas which does the return over the updated vertex elements, it just reads, never updates the struct, so the (s vertex) is the s copy of the vertex instance, while (s *vertex) s is the pointer to the vertex instance

// we use the pointer reciever method when we are sure about the mutation of the struct, that is when we have to chnage the valuesof that struct, then we use the pointer receiver method. For better performance, we use this as well, becausse we can accees only tht specific part of that struct, there is no need to copy the whole struct all together. It's better of consistency
// while the varaible reciever method is used when we don't want any mutatin, we are sure that we will go with the immutation and for better safety, as it makes sure that we don't make any changes in the struct
