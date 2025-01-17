package рекурсивный_bts

import "fmt"

// Узел в бинарном дереве
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Бинарное дерево
type Tree struct {
	Root *Node
}

// Создание пустого дерева
func createBinary() *Tree {
	return &Tree{}
}

// Добавление значения в дерево
func (t *Tree) Push(i int) {
	t.Root = AddItem(t.Root, i)
}

// Добавление элемента в дерево (рекурсивно)
func AddItem(node *Node, i int) *Node {
	if node == nil {
		return &Node{Data: i}
	}

	if i < node.Data {
		node.Left = AddItem(node.Left, i)
	} else if i > node.Data {
		node.Right = AddItem(node.Right, i)
	}

	return node
}

// Обход и вывод дерева
func (t *Tree) Item() {
	t.CheckItem()
}

// Метод для обхода (инфиксный обход)
func (t *Tree) CheckItem() {
	stack := []*Node{}
	cur := t.Root

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(cur.Data) // Вывод данных узла

		cur = cur.Right
	}
}

// Удаление значения из дерева
func (t *Tree) Pop(i int) {
	t.Root = Delete(t.Root, i)
}

// Удаление элемента из дерева
func Delete(node *Node, i int) *Node {
	if node == nil {
		return nil
	}

	if i < node.Data {
		node.Left = Delete(node.Left, i)
	} else if i > node.Data {
		node.Right = Delete(node.Right, i)
	} else {
		// Если найден узел для удаления
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Поиск минимального элемента в правом поддереве
		small := SmallFind(node.Right)
		node.Data = small.Data
		node.Right = Delete(node.Right, small.Data)
	}

	return node
}

// Поиск минимального узла
func SmallFind(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return SmallFind(node.Left)
}

// Главная функция
func main() {
	t1 := createBinary()

	// Добавление элементов
	t1.Push(8)
	t1.Push(5)
	t1.Push(6)
	t1.Push(9)

	// Удаление элемента
	t1.Pop(8)

	// Вывод дерева
	t1.Item()
}
