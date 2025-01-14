package из_List_в_Deque

import (
	"container/list"
	"fmt"
)

type Deque struct {
	Stack *list.List
}

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func createLinked(arr []int) *LinkedList {
	l := &LinkedList{}

	for _, v := range arr {
		l.PushBack(v)
	}

	l.Item()
	return l
}

func createStack(l *LinkedList) *Deque {
	lnew := &Deque{Stack: list.New()}

	for cur := l.Head; cur != nil; cur = cur.Next {
		lnew.Stack.PushFront(cur.Data)
	}

	for cur := lnew.Stack.Front(); cur != nil; cur = cur.Next() {
		fmt.Println(cur.Value)
	}

	return lnew
}

func (l *LinkedList) PushBack(i int) {
	n := &Node{Data: i}

	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}

	left := l.Tail
	l.Tail.Next = n
	l.Tail = n
	l.Tail.Prev = left

}

func (l *LinkedList) Item() {
	for cur := l.Head; cur != nil; cur = cur.Next {
	}
}

func main() {
	l1 := createLinked([]int{1, 2, 3, 4, 5})

	d1 := createStack(l1)

	fmt.Println("ok", l1, d1)

}
