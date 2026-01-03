package main

import "fmt"

type OrderStatus string

// const (
//
//	Pending OrderStatus = iota
//	Processing
//	Shipped
//	Delivered
//
// )

const (
	Pending    OrderStatus = "pending"
	Processing OrderStatus = "processing"
	Shipped    OrderStatus = "shipped"
	Delivered  OrderStatus = "delivered"
)

func changeStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}

func main() {
	changeStatus(Shipped)
}
