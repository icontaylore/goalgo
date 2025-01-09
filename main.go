package main

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int   // значение
	Next *Node // ссылка на след эл.
}

func (l *LinkedList) PushBack(i int) {
	n := &Node{Data: i, Next: nil}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *LinkedList) PushFront(i int) {
	n := &Node{Data: i, Next: l.Head}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}

	l.Head = n
}

func (l *LinkedList) SearchEl(i int) int {
	current, count := l.Head, 1

	for current.Next != nil {
		if current.Data == i {
			fmt.Println("yes", current)
			break
		}
		count++
		current = current.Next
	}

	return current.Data
}

func (l *LinkedList) Insert(i int, j int) {
	current := l.Head
	n := &Node{Data: i}

	for current != nil {
		if current.Data == j {
			n.Next = current.Next
			current.Next = n

			if current == l.Tail {
				l.Tail = n
			}
			return
		}
		current = current.Next
	}

	fmt.Println(n.Next.Data)
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
	l1.Insert(20, 9)
}
