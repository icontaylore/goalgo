package односвязный_список

import "fmt"

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Data int
	Next *Node
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

func (l *LinkedList) EraseLast() {
	if l.Head == nil {
		fmt.Println("нету элементов")
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		l.Tail = nil
	}

	current := l.Head
	var prev *Node

	prev = current

	for current.Next != nil {
		prev = current
		current = current.Next
	}

	prev.Next = nil
}

func (l *LinkedList) EraseFirst() {
	if l.Head == nil {
		fmt.Println("пусто")
		return
	}

	l.Head = l.Head.Next
}

func (l *LinkedList) Erase(i int) {
	current := l.Head

	if l.Head.Data == i {
		l.Head = current.Next
		fmt.Println("эл. был первым")

		return
	}

	var left *Node

	for current != nil {
		if current.Data == i && current.Next != nil {
			left.Next = current.Next
			break

		} else if current.Next == nil {
			l.Tail = left
		}

		left = current
		current = current.Next
	}

}

func (l *LinkedList) Range() {
	current := l.Head

	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func main() {
	n, i := 0, 0
	fmt.Scan(&n)

	l1 := &LinkedList{}
	for i < n {
		a := 0
		fmt.Scan(&a)
		l1.PushBack(a)
		i++
	}
	l1.EraseLast()
	l1.Erase(6)
	fmt.Println(l1.Tail)

	l1.Range()

}
