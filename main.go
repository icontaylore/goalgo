package main

import "fmt"

// LinkedList
type LinkedList struct {
	Head *Node // first
	Tail *Node // last
}

// Node struct
type Node struct {
	Data int   // data
	Next *Node //
}

func (l *LinkedList) PushBack(num int) {
	n := &Node{Data: num, Next: nil}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}

	l.Tail.Next = n
	l.Tail = n
}

func (l *LinkedList) PushFront(num int) {
	n := &Node{Data: num, Next: l.Head}

	if l.Head == nil {
		l.Tail = n
	}

	l.Head = n
}

func (l *LinkedList) Search(num int) {
	current, count := l.Head, 1 // текущий эл списка. счётчик

	for current != nil {
		if current.Data == num {
			fmt.Println("yes this num")
			return
		}
		count++
		current = current.Next
	}

	fmt.Println(current.Data, count)
}

func (l *LinkedList) Insert(num int, j int) {
	current := l.Head
	n := &Node{Data: num}

	for current != nil { // 6 9 ; 10 11 ; 6>9
		if current.Data == j {
			fmt.Println("зацепились")
			n.Next = current.Next
			current.Next = n

			return
		}
		current = current.Next
	}

}

func (l *LinkedList) Erase(num int) {
	current := l.Head // начальное состояние
	prev := &Node{}

	if num == current.Data {
		current = current.Next
	}

	for current != nil && current.Data != num { // 6 9 10
		prev = current
		current = current.Next
	}

	if current.Next == nil {
		l.Tail = prev
	}
	prev.Next = current.Next

}

func main() {
	l1 := &LinkedList{}
	l1.PushBack(1)
}
