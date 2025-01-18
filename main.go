package main

import (
	"fmt"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) Push(i int) {
	t.Root = PushItem(t.Root, i)
}

// PUSH AND ITEM
func PushItem(node *Node, i int) *Node {
	n := &Node{Data: i}

	if node == nil {
		return n
	}

	if i < node.Data {
		node.Left = PushItem(node.Left, i)
	} else if i > node.Data {
		node.Right = PushItem(node.Right, i)
	}
	return node
}

func (t *Tree) Item() []int {
	stack := []*Node{}
	cur := t.Root
	arr := []int{}

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		arr = append(arr, cur.Data)
		stack = stack[:len(stack)-1]

		cur = cur.Right
	}

	return arr
}

// PUSH AND ITEM

// POP AND ELEM
func (t *Tree) Pop(i int) {
	t.Root = PopOut(t.Root, i)
}

func PopOut(node *Node, i int) *Node {
	if node.Data == i {
		small := FindSmall(node.Right)
		node.Data = small.Data
		node.Right = DeleteElem(node.Right, small.Data)
	}

	if i < node.Data {
		node.Left = DeleteElem(node.Left, i)
	} else if i > node.Data {
		node.Right = DeleteElem(node.Right, i)
	}

	return node
}

func FindSmall(node *Node) *Node {
	for node.Left == nil {
		return node
	}

	return FindSmall(node.Left)
}

func DeleteElem(node *Node, i int) *Node {
	if node == nil {
		return nil
	}

	if i > node.Data {
		DeleteElem(node.Right, i)
	} else if i < node.Data {
		DeleteElem(node.Left, i)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		smal := FindSmall(node.Right)
		node.Data = smal.Data
		node.Right = DeleteElem(node.Right, smal.Data)
	}

	return node
}

func (t *Tree) ShowWide() *Node {
	if t.Root == nil {
		return nil
	}

	queue := []*Node{t.Root}

	for len(queue) > 0 {
		lvl := len(queue)

		for i := 0; i < lvl; i++ {
			node := queue[0]
			queue = queue[1:]
			fmt.Print(node.Data, " ")

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		fmt.Println()
	}

	return nil
}

func CheckPlenty(arr1 []int, arr2 []int) []int {
	arr := []int{}
	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr2[j] < arr1[i] {
			j++
		} else if arr2[j] > arr1[i] {
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
			i++
		}
	}
	return arr
}

func main() {
	t1 := &Tree{}

	t1.Push(8)
	t1.Push(5)
	t1.Push(6)
	t1.Push(4)
	t1.Push(10)
	t1.Push(9)
	t1.Push(11)

	arr1 := t1.Item()
	fmt.Println(arr1)

	t2 := &Tree{}

	t2.Push(7)
	t2.Push(4)
	t2.Push(9)
	t2.Push(10)

	arr2 := t2.Item()
	fmt.Println(arr2)

	t := CheckPlenty(arr1, arr2)
	fmt.Println(t)

}
