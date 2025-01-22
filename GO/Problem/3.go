package main

import (
	"context"
	"fmt"
)

func generator(ctx context.Context, arr ...int) <-chan int {
	channel := make(chan int)

	go func() {
		defer close(channel)
		for _, v := range arr {
			select {
			case <-ctx.Done():
				fmt.Println("cancel")
			case channel <- v:

			}
		}
	}()

	return channel
}

func squarer(ctx context.Context, channel <-chan int) <-chan int {
	channelTwo := make(chan int)

	go func() {
		defer close(channelTwo)
		for v := range channel {
			select {
			case <-ctx.Done():
				fmt.Println("cancel")
			case channelTwo <- v * v:

			}
		}
	}()

	return channelTwo
}

func main() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))

	for x := range pipeline {
		fmt.Println(x)
	}
}
