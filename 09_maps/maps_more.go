package main

import (
	"fmt"
	"sort"
	"strings"
)

// map_literal_short -> inside a map literal, if the value type is a struct we can skip writing vertex{...} again and again
func map_literal_short() {
	places := map[string]vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(places, "| total keys:", len(places)) // len works on maps too
}

// iterate_map -> range over a map gives key and value
// the ORDER IS RANDOM every run, go does this on purpose so we never depend on map order
// if we need an order, we take the keys out, sort them and loop over the sorted keys
func iterate_map() {
	ages := map[string]int{"shreyas": 20, "advit": 10, "rahul": 25}
	for name, age := range ages {
		fmt.Print(name, "=", age, " ")
	}
	fmt.Println("<- random order")

	keys := make([]string, 0, len(ages))
	for k := range ages {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Print(k, "=", ages[k], " ")
	}
	fmt.Println("<- sorted by key")
}

// comma_ok_in_if -> the most common way to check a key, ok is true only if the key really exists
// needed because a missing key gives back the zero value (0), and 0 could also be a real stored value
func comma_ok_in_if() {
	stock := map[string]int{"apple": 0, "mango": 12}
	for _, fruit := range []string{"apple", "mango", "kiwi"} {
		if qty, ok := stock[fruit]; ok {
			fmt.Println(fruit, "found, qty =", qty)
		} else {
			fmt.Println(fruit, "not in the map")
		}
	}
}

// nil_map -> a map declared with var and not made is nil
// READING from a nil map is fine (gives zero value) but WRITING to it panics, so always make() it first
func nil_map() {
	var m map[string]int
	fmt.Println("nil map read:", m["x"], "| m == nil ?", m == nil)
	// m["x"] = 1 -> panic: assignment to entry in nil map
	m = make(map[string]int)
	m["x"] = 1
	fmt.Println("after make:", m)
}

// word_count -> a classic real use of a map, counting how many times each word comes
// m[word]++ works even for a new word because the missing key starts from the zero value 0
func word_count(text string) map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(text) { // Fields splits the text on spaces
		counts[word]++
	}
	return counts
}

// map_of_slices -> the value of a map can be anything, even a slice
func map_of_slices() {
	skills := map[string][]string{}
	skills["shreyas"] = append(skills["shreyas"], "go", "python") // nil slice + append works fine
	skills["advit"] = append(skills["advit"], "drawing")
	fmt.Println(skills)
}

// map_vs_struct -> same data in both forms, to see the difference from the comments in maps.go
// struct -> fixed fields known at compile time, accessed by dot, can't loop over fields
// map    -> keys can be added / removed at runtime, accessed by key, can loop over it
func map_vs_struct() {
	type config struct {
		host string
		port int
	}
	c := config{host: "localhost", port: 8080}
	fmt.Println("struct:", c.host, c.port)
	// c.timeout = 5 -> compile error, the struct has no field timeout

	settings := map[string]string{"host": "localhost", "port": "8080"}
	settings["timeout"] = "5s" // fine, new keys anytime
	delete(settings, "host")
	fmt.Println("map:", settings)
}
