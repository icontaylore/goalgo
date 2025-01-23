package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type result struct {
	msg string
	err error
}
type searh func() *result
type replicas []searh

func fakeSearch(kind string) searh {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func getFirstResult(ctx context.Context, replic replicas) *result {
	ch := make(chan *result, len(replic))

	for _, v := range replic {
		go func(f searh) {
			select {
			case <-ctx.Done():
				fmt.Println("time out")
			case ch <- f():
			}
		}(v)
	}

	select {
	case <-ctx.Done():
		return &result{err: fmt.Errorf("time Out")}
	case val := <-ch:
		val.err = fmt.Errorf("none")
		return val
	}
}

func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	returnArray := []*result{}
	wg := sync.WaitGroup{}

	wg.Add(len(replicaKinds))
	for _, arr := range replicaKinds {
		go func() {
			defer wg.Done()
			returnArray = append(returnArray, getFirstResult(ctx, arr))
		}()
	}

	wg.Wait()
	return returnArray
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	replicaKinds := []replicas{
		replicas{fakeSearch("web1"), fakeSearch("web2")},
		replicas{fakeSearch("image1"), fakeSearch("image2")},
		replicas{fakeSearch("video1"), fakeSearch("video2")},
	}

	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}

}
