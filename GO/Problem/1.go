package main

import (
	"fmt"
	"sync"
)

func fillChan(n int) <-chan int {
	chanel := make(chan int) // 1 ; 1

	go func() {
		defer close(chanel)
		for i := 0; i < n; i++ {
			chanel <- i
		}
	}()

	return chanel
}

func merge(cs ...<-chan int) <-chan int {
	chanel := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(len(cs) + 1)
	go func() {
		defer wg.Done()
		for _, arr := range cs { // 1,2,3
			go func(ch <-chan int) {
				defer wg.Done()
				for v := range ch {
					chanel <- v
				}
			}(arr)
		}
	}()

	go func() {
		wg.Wait()
		defer close(chanel)
	}()
	return chanel
}

func main() {
	a := fillChan(2)
	b := fillChan(3)
	c := fillChan(4)
	//
	d := merge(a, b, c)

	for v := range d {
		fmt.Println(v)
	}

}
