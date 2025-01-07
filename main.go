package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Структура для пользователя, включающая ID, email и лог действий
type User struct {
	id    int
	email string
	log   []Cache
}

// Доступные действия для логов пользователей
var LoggyAction = []string{
	"replay", "delete", "start rec", "stop rec", "back rec",
	"start rec", "stor back", "exit", "start",
}

// Структура для хранения действия с временем
type Cache struct {
	A string    // Описание действия
	t time.Time // Время действия
}

func main() {
	// Инициализация генератора случайных чисел
	rand.NewSource(time.Now().Unix())

	// Генерация 100 пользователей
	u := generateUser(100)

	// Печать информации о каждом пользователе
	for _, v := range u {
		writeFile(v)
	}

}

func writeFile(usr User) error {
	filename := fmt.Sprintf("Structures/структура, вывод кэша/uid_%d.txt", usr.id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(usr.getInfo())
	defer file.Close()
	return nil
}

// Функция для создания списка пользователей
func generateUser(n int) []User {
	// Создание среза пользователей
	user := make([]User, n+1)
	for i := 1; i <= n; i++ {
		// Создание пользователя с случайными действиями
		user[i] = User{id: i, email: fmt.Sprintf("email%d@mail.com", i), log: generateRandomAction(5)}
	}
	return user
}

// Функция для генерации случайных действий
func generateRandomAction(n int) []Cache {
	log := make([]Cache, n)
	for i := 0; i < n; i++ {
		// Генерация случайного действия из LoggyAction
		log[i] = Cache{A: LoggyAction[rand.Intn(len(LoggyAction))], t: time.Now()}
	}
	return log
}

// Метод для получения информации о пользователе
func (u User) getInfo() string {
	// Стартовое сообщение с информацией о пользователе
	out := fmt.Sprintf("User: %s - id: %d\n", u.email, u.id)
	// Добавление информации о действиях пользователя
	for _, v := range u.log {
		// Форматирование времени для действий
		out += fmt.Sprintf("  Action: %s at %s\n", v.A, v.t.Format(time.RFC3339))
	}
	return out
}
