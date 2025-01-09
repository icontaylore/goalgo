package main

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

func main() {
	l1 := &LinkedList{}
	l1.PushBack(1)  // Добавляем 1 в конец
	l1.PushFront(2) // Добавляем 2 в начало
}
