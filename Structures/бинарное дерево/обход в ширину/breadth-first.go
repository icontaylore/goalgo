package обход_в_ширину

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
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

	return t.Root
}

func main() {
	t1 := &Tree{}

	//t1.Push(8)
	//t1.Push(5)
	//t1.Push(6)
	//t1.Push(4)
	//t1.Push(10)
	//t1.Push(9)
	//t1.Push(11)

	t1.ShowWide()

}
