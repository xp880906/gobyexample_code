package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			job, ok := <-jobs
			if ok {
				fmt.Println("received job: ", job)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
		fmt.Println("Sent job: ", i)
	}

	fmt.Println("Sent all jobs")
	close(jobs)
	<-done
}
