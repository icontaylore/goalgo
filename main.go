package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) PushBack(i int) {
	n := &Node{Data: i, Next: nil}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *LinkedList) Check() bool {
	cur := l.Head
	count := 0

	for cur != nil {
		if count == 3 {
			break
		}

		count++
		cur = cur.Next
	}

	if count == 3 {
		return false
	}

	return true
}

func main() {
	l1 := &LinkedList{}
	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	r := l1.Check()
	fmt.Println(r)
}
