package main

// maps is a concept in golang which is the same as that of the hashmaps in python, they are used in fast searching and editing, deleteing, as we acces the element not by linear sarch, butrather by direct key search, thus maps in golang has key value pair like thing

import "fmt" 

type vertex struct {
	lat float64
	long float64
}

var m map[string]vertex // this is the declaration of the map

func mapping_use() {
	m = make(map[string]vertex) // this is the allocation or initiliazation of the map, basiially assigning values in the map
	m["Bell Labs"] = vertex{
		40.12345, -75.6767,
	}
	fmt.Println(m["Bell Labs"])
} 

// so here we have decalred and initiliazided the values to the map in the same line itself, where as aboe we had declared earlier and initiliazed later in the function

var mapping = map[string]vertex {
	"bitcoin" : {
		58000.12,
		120000.15,
	},
}
func mapping_in_declaration() {
	fmt.Println(m)
}

func other_operation_mapping() {
	m := make(map[string]int)
	m["answer"] = 42 
	fmt.Println("the value of the answer is: ", m["answer"])
	m["answer"] = 48
	fmt.Println("the new value fo the answer key is: ", m["answer"])
	delete(m, "answer") // this is the syntax of the delete function, here we first have to keep the object on which the delete operation is to be performed and then we add the key which is to be deleted
	fmt.Println("the delete value is: ", m["answer"])
	v, ok := m["answer"]
	fmt.Println("the value is: ", v, "present?", ok) 
	// so when we get any value from the amp, it sends us two things, first the value itself, and second the status of that value, that whetehr that value was actually present or not
}