package поиск_отрицательного_числа

import "fmt"

func main() {
	// Инициализация массива
	array := []int{1, 1, 1, -2}

	// Переменная для хранения результата
	res := 0

	// Если первый элемент массива не отрицательный, ищем количество положительных чисел
	if array[0] >= 0 {
		// Количество положительных чисел = длина массива минус индекс первого отрицательного
		res = len(array) - bs(array)
	} else {
		// Если первый элемент отрицателен, то весь массив состоит из отрицательных чисел
		res += len(array)
	}

	// Выводим результат
	fmt.Println(res)
}

// Функция для поиска индекса первого отрицательного числа в отсортированном массиве
func bs(arr []int) int {
	left := 0             // Левая граница для поиска
	right := len(arr) - 1 // Правая граница для поиска

	// Бинарный поиск для нахождения первого отрицательного числа
	for left+1 != right {
		// Средняя точка для раздела массива
		mid := (left + right) / 2

		// Если средний элемент положительный, сдвигаем левую границу
		if arr[mid] >= 0 {
			left = mid
		} else {
			// Если средний элемент отрицательный, сдвигаем правую границу
			right = mid
		}
	}

	// Возвращаем индекс первого отрицательного числа
	return right
}
