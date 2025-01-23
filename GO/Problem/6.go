package main

import "fmt"

func mergeSorted(a, b <-chan int) <-chan int {
	channel := make(chan int)
	go func() {
		defer close(channel)

		val1, ok1 := <-a
		val2, ok2 := <-b

		for ok1 && ok2 { // true && true
			if val1 < val2 {
				channel <- val1
				val1, ok1 = <-a
			} else {
				channel <- val2
				val2, ok2 = <-b
			}
		}

		for ok1 {
			channel <- val1
			val1, ok1 = <-a
		}

		for ok2 {
			channel <- val2
			val2, ok2 = <-b
		}
	}()

	return channel
}

func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)

	c := mergeSorted(a, b)

	for val := range c {
		fmt.Printf("%d ", val)
	}
}
