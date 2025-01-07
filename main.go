package main

import (
	"fmt"
	"math/rand"
	"os"
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

var ActArr = []string{
	"delete",
	"reg",
	"exit",
	"photo",
	"foxus",
	"starty run",
}

// Метод createTableCahce выводит список действий апокалипсиса и их временные метки
func (a Apocalypse) createTableCahce() string {
	out := ""
	for i, v := range a.LogChaos {
		// Используем сохраненное время v.t, а не текущее время time.Now()
		out = fmt.Sprintf("%d: Action: %s, Time: %s\n", i, v.Action, v.t.Format(time.RFC3339))
	}
	return out
}

func main() {
	rand.New(rand.NewSource(time.Now().Unix()))

	u := genetateUser(100)

	for i, user := range u {
		write(user, i+1)
	}
}

func write(u Apocalypse, id int) error {
	fmt.Printf("writen user: id - %d", u.Name)

	filename := fmt.Sprintf("logs/uid_%d.txt", id)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	_, err = file.WriteString(u.createTableCahce())
	return nil
}

func genetateUser(count int) []Apocalypse {
	user := make([]Apocalypse, count)

	for i := 1; i < count; i++ {
		user[i] = Apocalypse{
			Name:      fmt.Sprintf("apocalypse%d reg", i),
			TypeChaos: "storm",
			LogChaos:  generateLog(rand.Intn(1000)),
		}
	}
	return user
}

func generateLog(count int) []ActionApocalypse {
	log := make([]ActionApocalypse, count)

	for i := 1; i < count; i++ {
		log[i] = ActionApocalypse{t: time.Now(), Action: ActArr[rand.Intn(len(ActArr)-1)]}
	}
	return log
}
