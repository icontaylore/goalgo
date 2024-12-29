package ксерокс

import (
	"fmt"
	"math"
)

func main() {
	n, x, y := 5, 1, 2 // 3

	// минимальное время (принтер)
	m := math.Min(float64(x), float64(y)) //

	// границы
	r := n * int(m) // граница справа
	l := 0          // граница слева

	for l+1 != r {
		mid := (l + r) / 2

		if mid/x+mid/y < r-1 {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Println(math.Min(float64(x), float64(y)) + float64(r))
}
