package main

import (
	"context"
	"fmt"
	"reflect"
)

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	channelReturn := make(chan interface{})

	go func(channelsRet chan interface{}) {
		defer close(channelReturn)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx is done")
			case channel, ok := <-in:
				if !ok {
					return
				}
				channelReturn <- channel
			}
		}
	}(channelReturn)
	return channelReturn
}

func main() {
	ch := make(chan interface{})
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var res []interface{}

	for v := range orDone(context.Background(), ch) {
		res = append(res, v)
	}

	fmt.Println(res)

	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	}
}
