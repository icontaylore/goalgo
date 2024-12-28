package main

import "fmt"

func main() {
	// Инициализация переменной c и массива arr
	c := 4
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	// Вызов функции поиска
	search(arr, c)
}

// Функция для поиска индекса в массиве с учетом условия
func search(arr []int, i int) {
	// Инициализация границ для бинарного поиска
	r := len(arr) - 1
	lowest := 0

	// Расчет максимальной разницы между элементами массива
	outer := arr[r] - arr[0]

	// Бинарный поиск
	for lowest+1 != r {
		// Средняя точка диапазона
		mid := (lowest + r) / 2

		// Проверка, соответствует ли текущее значение условию
		if check(arr[mid], i, outer, arr[0]) == false {
			// Если условие не выполняется, сдвигаем левую границу
			lowest = mid
		} else {
			// Если условие выполняется, сдвигаем правую границу
			r = mid
		}

		// Выводим текущую правую границу для понимания промежуточных результатов
		fmt.Println(r, "ответ *можно и mid указать*")
	}
}

// Функция для проверки условия
func check(i int, j int, max int, a0 int) bool {
	// Логика проверки условия
	res := true
	if (i-a0)*(j-1) < max {
		res = false
	}
	return res
}
