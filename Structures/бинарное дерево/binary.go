package main

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Push(i int) {
	t.Root = Insert(t.Root, i)
}

func Insert(node *Node, i int) *Node {
	n := &Node{Data: i}

	if node == nil {
		return n
	}

	if i < node.Data {
		node.Left = Insert(node.Left, i)
	} else if i > node.Data {
		node.Right = Insert(node.Right, i)
	}

	return node
}

func main() {
	t1 := &Tree{}

	t1.Push(1)
	t1.Push(2)

	t1.Item()
}

func (t *Tree) Item() {
	for cur := t.Root; cur != nil; cur = cur.Left {
		fmt.Println(cur.Left, cur.Right)

	}

	for cur := t.Root; cur != nil; cur = cur.Right {
		fmt.Println(cur.Left, cur.Right)
	}
}
