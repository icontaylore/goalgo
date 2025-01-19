package main

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	mu    sync.RWMutex // можно без *, и так передается ссылка на поле структуры
	value int
}

func (u *User) Increment() {
	u.mu.Lock() // блокируем доступ к конкретному значению для других горутин
	u.value++
	u.mu.Unlock() // разблокируем
}

func (u *User) Val() int {
	u.mu.RLock()         // блокируем только чтение
	defer u.mu.RUnlock() // разблокируем
	return u.value
}

func main() {
	u1 := &User{} // создается ссылка на структуру User

	for i := 0; i < 1500; i++ {
		go func() {
			u1.Increment()
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(u1.value)
}
