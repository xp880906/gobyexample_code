package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker ", id, " started job ", j)
		time.Sleep(time.Second)
		fmt.Println("worker ", id, " started job ", j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		<-results
	}
}
