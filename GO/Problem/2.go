package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const timeoutLimit = 90

type Result struct {
	msg string
	err error
}

func fakeDownload(url string) Result {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	if r > timeoutLimit {
		return Result{
			err: errors.New(fmt.Sprintf("failed to download data %s\n", url)),
		}
	}

	return Result{
		msg: fmt.Sprintf("dowloaded data from %s\n", url),
	}
}

func download(urls []string) ([]string, error) { //arr so big
	var (
		arr []string
		err error
		mu  sync.Mutex
		wg  sync.WaitGroup
	)

	for _, str := range urls {
		wg.Add(1) // 1 go run

		go func(url string) { // go is start, go end to Done
			defer wg.Done()

			msg := fakeDownload(url) // функция запускается в горутине, сразу лочим её через мютекс
			mu.Lock()
			defer mu.Unlock()

			if msg.err != nil {
				errors.Join(err, msg.err)
			} else {
				arr = append(arr, msg.msg)
			}
		}(str)
	}

	wg.Wait()
	return arr, err
}

func main() { // block
	msgs, err := download([]string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml",
		"https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml",
		"https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml",
		"https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml",
		"https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(msgs)
}
