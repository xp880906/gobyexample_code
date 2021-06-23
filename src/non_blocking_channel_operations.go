package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)
	default:
		fmt.Println("No message")
	}

	// 为什么这里发送不了
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("Message sent")
	default:
		fmt.Println("No message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Message Received: ", msg)
	case sig := <-signals:
		fmt.Println("Signal Received: ", sig)
	default:
		fmt.Println("No activity")
	}
}
