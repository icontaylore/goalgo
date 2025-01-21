package main

import (
	"fmt"
	"sync"
)

func fillchan(n int) <-chan int {
	chanel := make(chan int)

	go func() {
		defer close(chanel)
		for i := 0; i < n; i++ {
			chanel <- i
		}
	}()
	return chanel
}

func merge(cs ...<-chan int) <-chan int {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for _, v := range cs {
		wg.Add(1)
		go func(chanel <-chan int) {
			defer wg.Done()
			for val := range chanel {
				ch <- val
			}
		}(v)
	}

	go func() {
		wg.Wait()
		defer close(ch)
	}()

	return ch
}

func main() {
	a := fillchan(2)
	b := fillchan(3)
	c := fillchan(4)

	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}
}
