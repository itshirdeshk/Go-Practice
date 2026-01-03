package main

import "fmt"

type paymenter interface {
	makePayment(amount float32)
}

type Payment struct {
	gateway paymenter
}

// Here, we are breaking open close principle but by using interfaces, we have solved this problem
func (p Payment) makePayment(amount float32) {
	// razorpayGateway := razorpay{}
	// razorpayGateway.makePayment(amount)
	// stripeGateway := stripe{}
	p.gateway.makePayment(amount)
}

type razorpay struct{}

func (r razorpay) makePayment(amount float32) {
	fmt.Println("Payment made via Razorpay", amount)
}

type stripe struct{}

func (r stripe) makePayment(amount float32) {
	fmt.Println("Payment made via Stripe", amount)
}

func main() {
	// stripeGateway := stripe{}
	razorpayGateway := razorpay{}
	newPayment := Payment{gateway: razorpayGateway}
	newPayment.makePayment(100)
}
