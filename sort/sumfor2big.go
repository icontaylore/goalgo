package main

func main() {
	array := []int{3, 7, -10, 65, 33, 14, 84, 35, 60, -100}
	sort(array)
}

func sort(arr []int) {
	l, m := 0, 0

	for _, v := range arr {
		if m < v { // если максимальное меньше значения
			l = m // лоу = макс
			m = v // макс = значение
		} else if l < v { // но если и лоу < значения
			l = v // тогда лоу = значению
		}
	}
}
