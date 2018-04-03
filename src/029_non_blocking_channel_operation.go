package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//go func() {
	//	messages <- "hi5"
	//}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("received messages",msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <- signals:
		fmt.Println("received signals",sig)
	default:
		fmt.Println("no activity")
	}
}
