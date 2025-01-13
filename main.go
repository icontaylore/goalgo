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

// Items - метод для получения всех элементов списка в виде среза
func (m *MaxLenList) Items() []interface{} {
	var items []interface{} // Срез для хранения элементов

	// Проходим по списку, начиная с первого элемента, и добавляем значения в срез
	for e := m.List.Front(); e != nil; e = e.Next() {
		items = append(items, e.Value) // Добавляем значение текущего элемента
	}
	return items
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
	dq.PushBack(4)

	// Выводим элементы списка
	fmt.Print(dq.Items()) // Результат: [2 3 4 5 4]
}
