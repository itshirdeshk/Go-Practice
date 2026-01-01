package main

import "fmt"

func main() {
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// default:
	// 	fmt.Println("Other than one or two")
	// }

	addLine := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("String")
		default:
			fmt.Println("Other")
		}
	}

	addLine(1)

	// var name string = "hirdesh";

}
