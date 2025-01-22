package main

import (
	"context"
	"fmt"
	"math/rand"
)

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	chanel := make(chan interface{})

	go func() {
		defer close(chanel)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx is cancel")
			case chanel <- fn():
			}
		}
	}()

	return chanel
}

func take(ctx context.Context, n <-chan interface{}, num int) <-chan interface{} {
	chanel := make(chan interface{})

	go func() {
		defer close(chanel)
		for i := 0; i < num; i++ {
			val, ok := <-n
			if !ok {
				fmt.Println("chanel is off")
				<-ctx.Done()
			} else {
				chanel <- val
			}
		}
	}()

	return chanel
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	randGen := func() interface{} { return rand.Intn(10) }

	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, randGen), 3) {
		res = append(res, num)
	}

	if len(res) != 3 {
		panic("wrong code")
	}

	// Печать значений, чтобы увидеть результат
	fmt.Println(res)
}
