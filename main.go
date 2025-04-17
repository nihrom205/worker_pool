package main

import (
	"fmt"
	"time"
)

func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Printf("worker %d start\n", id)
		time.Sleep(1 * time.Second)
		fmt.Printf("worker %d end\n", id)
		result <- j * 2
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for i := 0; i <= 3; i++ {
		go worker(i, jobs, result)
	}

	for i := 0; i < 10; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 10; i++ {
		fmt.Printf("worker %d end\n", <-result)
	}
}
