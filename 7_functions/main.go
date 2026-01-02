package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func languages() (string, string, string) {
	return "JS", "Go", "C++"
}

func processIt(fn func(a int) int) int {
	return fn(1)
}

func main() {
	result := add(4, 5)
	fmt.Println(result)
	fmt.Println(languages())

	fn := func(a int) int {
		return 1
	}

	fmt.Println(processIt(fn))
}
