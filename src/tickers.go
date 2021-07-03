package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker Ticked")
			case <-done:
				return
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker Stopped")
}
