package main

import "fmt"

func main() {
	// var nums []int
	// fmt.Println(len(nums))

	nums := make([]int, 2, 5)
	fmt.Println(cap(nums))
}
