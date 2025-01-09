package main

import "fmt"

// Односвязанный список
type LinkedList struct {
	Head *Node
	Tail *Node
}

// Структура узла
type Node struct {
	Data int
	Next *Node
}

// Добавление элемента в начало списка
func (l *LinkedList) PushFront(i int) {
	n := &Node{Data: i, Next: l.Head}

	// Если список пуст, новый элемент будет и Head, и Tail
	if l.Head == nil {
		l.Tail = n
	}
	l.Head = n
}

// Добавление элемента в конец списка
func (l *LinkedList) PushBack(i int) {
	n := &Node{Data: i, Next: nil}

	// Если список пуст, новый элемент будет и Head, и Tail
	if l.Head == nil {
		l.Head, l.Tail = n, n
	} else {
		// Если список не пуст, добавляем элемент в конец
		l.Tail.Next = n
		l.Tail = n
	}
}

func (l *LinkedList) searchEl(i int) {
	current, count := l.Head, 0

	for current != nil {
		if current.Data == i {
			fmt.Println(current.Data, current.Next, count)
		}
		count++
		current = current.Next
	}
}

func main() {
	l1 := &LinkedList{}
	l1.PushBack(1)   // Добавляем 1 в конец
	l1.PushFront(2)  // Добавляем 2 в начало
	l1.PushFront(3)  // Добавляем 3 в начало
	l1.PushFront(11) // Добавляем 3 в начало
	l1.PushFront(10) // Добавляем 3 в начало
	l1.PushBack(9)   // Добавляем 3 в начало
	l1.PushFront(6)  // Добавляем 3 в начало
	l1.PushFront(4)  // Добавляем 3 в начало
	l1.searchEl(9)
}
