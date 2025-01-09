package main

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	data int   // значение
	next *Node // ссылка на след эл.
}

func (l *LinkedList) PushBack(i int) {
	n := &Node{data: i, next: nil}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}

	l.Tail.next = n
	l.Tail = n
}

func (l *LinkedList) PushFront(i int) {
	n := &Node{data: i, next: l.Head}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}

	l.Head = n
}

func (l *LinkedList) searchEl(i int) {
	current, count := l.Head, 1

	for current.next != nil {
		if current.data == i {
			fmt.Println("yes", current)
			break
		}
		count++
		current = current.next
	}

	fmt.Println(count)
}

func main() {
	l1 := &LinkedList{}
	l1.PushBack(2)  // 1
	l1.PushBack(3)  // 2
	l1.PushBack(6)  // 3
	l1.PushBack(9)  // 4
	l1.PushBack(10) // 5
	l1.PushBack(11) // 6
	l1.PushBack(12) // 7
	l1.PushBack(16) // 8
	l1.searchEl(9)
}
