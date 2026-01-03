package main

import (
	"fmt"
	"time"
)

type Order struct {
	id        int
	amount    string
	status    string
	createdAt time.Time
}

func newOrder(id int, amount string, status string) *Order {
	return &Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
}

func (o *Order) changeStatus(newStatus string) {
	o.status = newStatus
}

func main() {
	order := Order{
		id:        1,
		amount:    "100",
		status:    "pending",
		createdAt: time.Now(),
	}

	order.changeStatus("paid")

	newOrder := *newOrder(2, "200", "pending")
	newOrder.changeStatus("paid")

	fmt.Println(newOrder)

	language := struct {
		name   string
		isGood bool
	}{
		name:   "Go",
		isGood: true,
	}

	fmt.Println(language)
}
