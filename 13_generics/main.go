package main

import "fmt"

//	func printSlice[T any](items []T) {
//		for _, item := range items {
//			fmt.Println(item)
//		}
//	}
//
//	func printSlice[T int | string](items []T) {
//		for _, item := range items {
//			fmt.Println(item)
//		}
//	}
// 
func printSlice[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack[T int | string] struct {
	elements []T
}

func main() {

	myStack := stack[string]{
		elements: []string{"true", "false"},
	}

	fmt.Println(myStack)

	printSlice([]int{1, 2, 3}, "numbers")
}
