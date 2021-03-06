package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}
	close(requests)

	// Time Limiter
	limiter := time.Tick(500 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request ", req, time.Now())
	}

	// 滑动窗口的限流
	bustyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		bustyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			bustyLimiter <- t
		}
	}()

	bustyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		bustyRequests <- i
	}
	close(bustyRequests)

	for req := range bustyRequests {
		<-bustyLimiter
		fmt.Println("Busty Request ", req, time.Now())
	}
}
