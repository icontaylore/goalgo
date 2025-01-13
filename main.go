package main

import (
	"container/list"
	"fmt"
)

// MaxLenList - структура, оборачивающая двусвязный список с ограничением на максимальную длину
type MaxLenList struct {
	List   *list.List // Основной список
	MaxLen int        // Максимальная длина списка
}

type Node struct {
}

// MaxLenCreate - функция-конструктор для создания списка с ограничением длины
func MaxLenCreate(max int) *MaxLenList {
	return &MaxLenList{
		List:   list.New(), // Создаём новый двусвязный список
		MaxLen: max,        // Задаём максимальную длину
	}
}

// PushBack - метод для добавления элемента в конец списка
func (l *MaxLenList) PushBack(i int) {
	// Если текущая длина списка больше или равна максимальной, удаляем первый элемент
	if l.List.Len() >= l.MaxLen {
		l.List.Remove(l.List.Front()) // Удаляем самый старый элемент (FIFO)
	}

	// Добавляем новый элемент в конец списка
	l.List.PushBack(i)
}

func (l *MaxLenList) PushFront(i int) {
	if l.List.Len() >= l.MaxLen {
		l.List.Remove(l.List.Back())
	}

	l.List.PushFront(i)
}

func (l *MaxLenList) Insert(i int, j int) {
	n := &list.Element{Value: j}

	if l.List == nil {
		return
	}

	for cur := l.List.Front(); cur != nil; cur = cur.Next() {
		if cur.Value == i {
			l.List.InsertAfter(n, cur)
			return
		}
	}
}

// Items - метод для получения всех элементов списка в виде среза
func (m *MaxLenList) Items() []interface{} {
	var items []interface{} // Срез для хранения элементов

	// Проходим по списку, начиная с первого элемента, и добавляем значения в срез
	for e := m.List.Front(); e != nil; e = e.Next() {
		items = append(items, e.Value) // Добавляем значение текущего элемента
	}
	return items
}

func (l *MaxLenList) PopLeft(i int) interface{} {
	if l.List == nil {
		return nil
	}

	right := l.List.Front()
	l.List.Remove(right)
	return right.Value
}

func (l *MaxLenList) PopRight(i int) interface{} {
	if l.List == nil {
		return nil
	}

	left := l.List.Back()
	l.List.Remove(left)
	return left.Value
}

func (l *MaxLenList) Pop(i int) interface{} {
	if l.List == nil {
		return nil
	}

	for e := l.List.Front(); e != nil; e = e.Next() {
		if e.Value == i {
			return l.List.Remove(e)
		}
	}

	return nil
}

func (l *MaxLenList) Extend(i []int) interface{} {
	if l.List == nil {
		return nil
	}

	for _, v := range i {
		l.PushBack(v)
	}

	return l.Items()
}

func (l *MaxLenList) ExtendLeft(i []int) interface{} {
	if l.List == nil {
		return nil
	}

	for _, v := range i {
		l.PushFront(v)
	}

	return l.Items()
}

func main() {
	// Исходный массив данных
	arr := []int{1, 2, 3, 4, 5}

	// Создаём список с максимальной длиной 5
	dq := MaxLenCreate(5)

	// Заполняем список элементами массива
	for _, v := range arr {
		dq.PushBack(v)
	}

	// Добавляем ещё один элемент, что вызывает удаление самого старого элемента
	fmt.Println(dq.ExtendLeft([]int{-3, -2, -1}))
	fmt.Println(dq.Extend([]int{6, 7, 8}))

	//Выводим элементы списка
	fmt.Println(dq.Items())

}
