package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) string {
	b := strings.Builder{}
	flag := true
	for i := 0; i < len(str); i++ {
		if flag == true {
			b.WriteString(strings.ToUpper(string(str[i])))
			flag = false
			i++
		}
		if string(str[i]) == " " {
			flag = true
		}
		b.WriteString(string(str[i]))
	}
	fmt.Println(b.String())
	return b.String()
}
