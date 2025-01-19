package main

import (
	"fmt"
	"time"
)

func main() {
	// Создание двух каналов для передачи строковых сообщений
	chanel := make(chan string)
	chanel2 := make(chan string)

	// Первая горутина отправляет строку "Запись 200ms" в канал chanel каждые 200 миллисекунд
	go func() {
		for {
			chanel <- "Запись 200ms"
			time.Sleep(time.Millisecond * 200)
		}
	}()

	// Вторая горутина отправляет строку "Запись 1s" в канал chanel2 каждую секунду
	go func() {
		for {
			chanel2 <- "Запись 1s"
			time.Sleep(time.Second)
		}
	}()

	// Главный цикл программы с select для обработки сообщений из обоих каналов
	for {
		select {
		// Когда сообщение поступает из канала chanel
		case msg := <-chanel:
			fmt.Println(msg)
		// Когда сообщение поступает из канала chanel2
		case msg := <-chanel2:
			fmt.Println(msg)
		default:
		}
	}
}
