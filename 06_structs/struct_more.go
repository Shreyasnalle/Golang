package main

import (
	"encoding/json"
	"fmt"
)

// address and student -> a struct can have another struct as a field (nested struct)
type address struct {
	city string
	pin  int
}

type student struct {
	name    string
	age     int
	address address // nested struct
}

// nested_struct -> we reach inner fields by chaining dots, s.address.city
func nested_struct() {
	s := student{name: "Shreyas", age: 20, address: address{city: "Pune", pin: 411001}}
	fmt.Printf("%+v\n", s) // %+v prints the field names too
	fmt.Println("city:", s.address.city)
}

// anonymous_struct -> a struct without a type name, used once on the spot (we used this inside a slice in 08_slices too)
func anonymous_struct() {
	point := struct {
		x, y int
	}{3, 4}
	fmt.Println("anonymous struct:", point)
}

// compare_structs -> two structs are equal with == when ALL their fields are equal
func compare_structs() {
	a := vertex{1, 2}
	b := vertex{1, 2}
	c := vertex{2, 1}
	fmt.Println("a == b ?", a == b, "| a == c ?", c == a)
}

// Payload -> the JSON use case from the comment in struct.go
// fields start with a CAPITAL letter so the encoding/json package (which is another package) can see them
// the `json:"..."` part is called a struct tag, it tells json what the key name should be
type Payload struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
	secret string // small letter -> not exported -> json ignores it completely
}

// struct_and_json -> struct -> json text (Marshal) and json text -> struct (Unmarshal)
func struct_and_json() {
	p := Payload{UserID: 1, Email: "test@example.com", Active: true, secret: "hidden"}
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("struct -> json:", string(bytes)) // secret is missing in the output

	var incoming Payload
	body := []byte(`{"user_id": 42, "email": "go@example.com", "active": false}`)
	if err := json.Unmarshal(body, &incoming); err != nil { // &incoming -> Unmarshal needs a pointer so it can fill our struct
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("json -> struct: %+v\n", incoming)
}
