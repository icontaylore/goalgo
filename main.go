package main

// односвязанный список
type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int
	Next *Node
}

func (l *LinkedList) PushFront(i int) {
	n := &Node{Data: i, Next: l.Head}

	if l.Head == nil {
		l.Tail = n
		l.Head = n
	} else if l.Head != nil {
		l.Head = n
	}
}
func (l *LinkedList) PushBack(i int) {
	n := &Node{Data: i, Next: nil}

	if l.Head == nil {
		l.Tail, l.Head = n, n
	} else if l.Head != nil {
		l.Tail.Next = n
		l.Tail = n
	}
}

func main() {
	l1 := &LinkedList{}
	l1.PushBack(1)
	l1.PushFront(2)
}
