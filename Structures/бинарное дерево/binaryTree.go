package бинарное_дерево

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func CreateTree() *Tree {
	return &Tree{}
}

func (t *Tree) Push(i int) {
	t.Root = AddItem(t.Root, i)
}

func AddItem(node *Node, item int) *Node {
	n := &Node{Data: item}

	if node == nil {
		return n
	}

	if item < node.Data {
		node.Left = AddItem(node.Left, n.Data)
	} else if item > node.Data {
		node.Right = AddItem(node.Right, n.Data)
	}

	return node
}

func (t *Tree) Item() {
	t.ItemCheck()
}

func (t *Tree) ItemCheck() {
	stack := []*Node{}
	cur := t.Root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		fmt.Println(cur)
		stack = stack[:len(stack)-1]

		cur = cur.Right
	}
}

func main() {
	t1 := CreateTree()

	t1.Push(10)
	t1.Push(12)
	t1.Push(3)

	t1.Item()

}
