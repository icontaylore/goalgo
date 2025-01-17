package прикладной_бинарный_поиск

import "fmt"

func main() {
	// Вызов основной функции поиска отрицательных чисел
	findnegative()
}

func findnegative() {
	// Инициализация переменных для считывания входных данных
	a, b := 0, 0
	fmt.Scan(&a, &b) // Считываем количество строк (a) и длину массива (b)
	res := 0         // Переменная для хранения результата

	// Инициализация массива нужного размера
	array := make([]int, b)

	// Цикл по строкам
	for i := 0; i < a; i++ {
		// Считывание элементов в массив
		for j := 0; j < b; j++ {
			fmt.Scan(&array[j])
		}

		// Если последний элемент массива не отрицательный, пропускаем строку
		if array[len(array)-1] >= 0 {
			res += 0
		} else {
			// В противном случае выполняем бинарный поиск
			bsearch(&array, &res)
		}
	}
	// Выводим итоговый результат
	fmt.Println(res)
}

// Функция бинарного поиска для подсчета количества отрицательных чисел в массиве
func bsearch(array *[]int, res *int) {
	// Инициализация границ поиска
	left := 0
	right := len(*array) - 1

	// Если первый элемент массива положительный или 0, пропускаем строку
	if (*array)[0] >= 0 {
		for left+1 != right {
			// Средняя точка массива
			mid := (left + right) / 2

			// Если средний элемент положительный, сдвигаем левую границу
			if (*array)[mid] >= 0 {
				left = mid
			} else {
				// Если средний элемент отрицательный, сдвигаем правую границу
				right = mid
			}
		}

		// Подсчитываем количество отрицательных чисел в массиве
		*res += len(*array) - right
		return
	} else {
		// Если первый элемент отрицательный, все числа в массиве отрицательные
		*res += len(*array)
	}
}
