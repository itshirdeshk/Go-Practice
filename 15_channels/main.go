package main

import (
	"fmt"
	"time"
)

// send
// func processNum(num chan int) {
// 	for num := range num {
// 		fmt.Println("Processing number", num)
// 	}
// }

func sendResult(resultChan chan int, num1 int, num2 int) {
	resultChan <- num1 + num2
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("Process...")
}

func emailQueue(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 1
	}()

	go func() {
		chan2 <- "hello"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1 := <-chan1:
			fmt.Println(chan1)
		case chan2 := <-chan2:
			fmt.Println(chan2)
		}
	}

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailQueue(emailChan, done)

	// for i := range 100 {
	// 	emailChan <- fmt.Sprintf("example@example%d.com", i)
	// }

	// fmt.Println("Email in the queue")
	// <-done

	// emailChan <- "example@example1.com"
	// emailChan <- "example@example2.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)
	// go task(done)

	// fmt.Println(<-done)

	// resultChan := make(chan int)

	// go sendResult(resultChan, 1, 2)

	// fmt.Println(<-resultChan)

	// intChan := make(chan int)

	// go processNum(intChan)

	// for {
	// 	intChan <- rand.Intn(100)
	// }

	// time.Sleep(time.Second * 2)

	// messsageChannel := make(chan string)

	// messsageChannel <- "ping"

	// msg := <- messsageChannel
	// fmt.Println(msg)
}
