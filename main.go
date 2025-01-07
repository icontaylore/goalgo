package main

import (
	"fmt"
	"time"
)

type Apocalypse struct {
	Name      string
	TypeChaos string
	LogChaos  []ActionApocalypse
}

type ActionApocalypse struct {
	Action string
	t      time.Time
}

func (a Apocalypse) createTableCahce() {
	for _, v := range a.LogChaos {
		fmt.Println(v.Action, time.Now())
	}
}

func main() {
	Tsunami := Apocalypse{Name: "Tsunami", TypeChaos: "Mega Wave", LogChaos: []ActionApocalypse{
		{Action: "MegaSuname", t: time.Now()},
		{Action: "MegaSuname2", t: time.Now()},
		{Action: "MegaSuname3", t: time.Now()},
		{Action: "MegaSuname4", t: time.Now()},
		{Action: "MegaSuname5", t: time.Now()},
	},
	}
	Tsunami.createTableCahce()
}
