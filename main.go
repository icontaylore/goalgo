package main

import (
	"fmt"
	"time"
)

// Структура Apocalypse описывает апокалипсис
// Name — название апокалипсиса
// TypeChaos — тип катастрофы
// LogChaos — список действий апокалипсиса
type Apocalypse struct {
	Name      string
	TypeChaos string
	LogChaos  []ActionApocalypse
}

// Структура ActionApocalypse описывает отдельное действие апокалипсиса
// Action — название действия
// t — время, когда это действие произошло
type ActionApocalypse struct {
	Action string
	t      time.Time
}

// Метод createTableCahce выводит список действий апокалипсиса и их временные метки
func (a Apocalypse) createTableCahce() {
	for _, v := range a.LogChaos {
		// Используем сохраненное время v.t, а не текущее время time.Now()
		fmt.Printf("Action: %s, Time: %s\n", v.Action, v.t.Format(time.RFC3339))
	}
}

func main() {
	// Создаем объект Tsunami типа Apocalypse
	Tsunami := Apocalypse{
		Name:      "Tsunami",   // Название катастрофы
		TypeChaos: "Mega Wave", // Тип катастрофы
		LogChaos: []ActionApocalypse{ // Лог действий с сохранением времени
			{Action: "MegaSuname", t: time.Now()},
			{Action: "MegaSuname2", t: time.Now().Add(1 * time.Minute)}, // Добавляем 1 минуту
			{Action: "MegaSuname3", t: time.Now().Add(2 * time.Minute)}, // Добавляем 2 минуты
			{Action: "MegaSuname4", t: time.Now().Add(3 * time.Minute)}, // Добавляем 3 минуты
			{Action: "MegaSuname5", t: time.Now().Add(4 * time.Minute)}, // Добавляем 4 минуты
		},
	}

	// Вызываем метод для вывода данных
	Tsunami.createTableCahce()
}
