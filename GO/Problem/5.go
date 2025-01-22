package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(jobs <-chan int, results chan<- int) {
	for v := range jobs {
		time.Sleep(time.Second)
		results <- v * v
	}
}

const numJobs = 15
const numWorkers = 3

func main() {
	jobs := make(chan int, numJobs)    //15
	results := make(chan int, numJobs) // 15
	wg := sync.WaitGroup{}

	for i := 0; i < numJobs; i++ {
		jobs <- i + 1
	}
	close(jobs)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(jobs, results)
		}() //
	}

	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}

}
