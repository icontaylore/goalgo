package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 10 * time.Millisecond

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	channel := make(chan struct{})

	go func() {
		executeTask()
		channel <- struct{}{}
		close(channel)
	}()

	select {
	case <-ctx.Done():
		return errors.New("time out")
	case <-channel:
		return nil
	}
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
