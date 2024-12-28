package дипломы_задача_про_квадрат_равный

import (
	"fmt"
	"math"
)

func main() {
	// Количество дипломов
	n := 12 // diploms
	// Ширина и высота
	w := 13 // width
	h := 12 // high

	// Расчет начального значения для правой границы поиска
	inf := math.Max(float64(w), float64(h)) * float64(n)

	// Левая и правая границы для бинарного поиска
	l := 0
	r := int(inf)

	// Бинарный поиск для нахождения минимального размера квадрата
	for l+1 != r {
		// Среднее значение между левой и правой границей
		mid := (l + r) / 2

		// Если квадрат больше нужного количества дипломов, сужаем правую границу
		if (mid/w)*(mid/h) > n {
			r = mid
		} else {
			// Если квадрат меньше или равен, сужаем левую границу
			l = mid
		}
	}

	// Выводим минимальный размер квадрата
	fmt.Println(r)
}
