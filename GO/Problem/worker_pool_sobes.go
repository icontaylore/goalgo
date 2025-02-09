package main

import (
	"fmt"
	"sync"
	"time"
)

// new 12 problem

func main() {
	pool := make(chan int, 10)

	for i := 1; i < 11; i++ {
		pool <- i
	}
	close(pool)

	wg := &sync.WaitGroup{}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go worker(pool, wg)
	}
	wg.Wait()
}

func worker(pool <-chan int, wg *sync.WaitGroup) {
	for v := range pool {
		printNumber(v)
	}
	wg.Done()
}

func printNumber(n int) {
	time.Sleep(time.Second)
	fmt.Println(n)
}
