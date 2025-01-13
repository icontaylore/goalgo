package двусвязный_список

import (
	"errors"
	"fmt"
)

type Node struct {
	Data int
	Prev *Node
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

	prev := l.Tail

	n = &Node{Data: i, Next: nil, Prev: prev}
	l.Tail.Next = n
	l.Tail = n
}

func (l *LinkedList) PushFront(i int) {
	n := &Node{Data: i, Next: l.Head}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	l.Head.Prev = n
	l.Head = n
}

func (l *LinkedList) Insert(i int, j int) error {
	if l.Head == nil {
		return errors.New("пустой список")
	}

	cur := l.Head
	for cur != nil {
		if cur.Data == j {
			left := cur
			right := cur.Next

			n := &Node{Data: i, Next: right, Prev: left}
			cur.Next = n
			cur = n
			return nil
		}

		cur = cur.Next
	}

	return nil
}

func (l *LinkedList) Erase(i int) error {
	current := l.Head

	if l.Head == nil {
		errors.New("список пуст")
	}

	if l.Head.Data == i {
		fmt.Println("yes")
		l.Head = l.Head.Next
		l.Head.Prev = nil

		return nil

	} else if l.Tail.Data == i {
		fmt.Println("yes2")
		left := l.Tail.Prev
		left.Next = nil

		return nil
	}

	for current != nil {
		if current.Next.Data == i {
			left := current
			right := current.Next.Next

			left.Next = right
			right.Prev = left

			return nil
		}

		current = current.Next
	}
	return nil
}

func (l *LinkedList) Range() {
	cur := l.Head

	for cur != nil {

		fmt.Println(cur.Data, "-data;", cur.Prev, "-prev;", cur.Next, "-next")
		cur = cur.Next
	}
}

func main() {
	l1 := &LinkedList{}

	l1.PushFront(1)
	l1.PushFront(2)
	l1.PushFront(3)
	l1.PushFront(4)

	l1.Range()

}
