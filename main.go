package main

import (
	"container/list"
	"fmt"
	"strings"
)

type Deque struct {
	Stack *list.List
}

func (s *Deque) createStackForFormula(str string) {
	newStr := strings.ReplaceAll(str, " ", "")
	arrStr := []rune{}

	for _, v := range newStr {
		arrStr = append(arrStr, v)
	}

	s.goToQueStack(arrStr)
}

func (s *Deque) goToQueStack(str []rune) {

	for i, v := range str {
		if string(v) == "(" {
			s.Stack.PushFront(int(i))

		} else if string(v) == ")" {

			slice := str[s.Stack.Front().Value.(int) : i+1]
			s.outSlice(slice)
			s.Stack.Remove(s.Stack.Front())

		}
	}

}

func (s *Deque) outSlice(arr []rune) string {
	str := ""

	for _, v := range arr {
		str += string(v)
	}

	fmt.Println(str)
	return str
}

func main() {
	s1 := &Deque{Stack: list.New()}
	data := "2 + 3 * (1 - 5 - (33 * x - 5)) + (a - b)"

	s1.createStackForFormula(data)
}
