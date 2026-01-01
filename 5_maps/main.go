package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["one"] = 1
	m["two"] = 1
	fmt.Println(m)

	newMap := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(newMap)

	v, ok := newMap["four"]
	fmt.Println(v, ok)

	if v, ok := newMap["three"]; ok {
		fmt.Println("Value:", v)
	} else {
		fmt.Println("Key not found")
	}
}
